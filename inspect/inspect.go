package inspect

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/tada-team/tdproto"
)

type Field struct {
	Name         string `json:"name"`
	JSName       string `json:"js_name"`
	Help         string `json:"help"`
	Json         string `json:"json"`
	Type         string `json:"type"`
	Null         bool   `json:"null,omitempty"`
	List         bool   `json:"list,omitempty"`
	InternalType bool   `json:"internal_type,omitempty"`
	Omitempty    bool   `json:"omitempty,omitempty"`
	TSType       string `json:"ts_type"`
}

type Struct struct {
	Name   string   `json:"name"`
	Help   string   `json:"help"`
	Fields []*Field `json:"fields"`
}

var tsTypeMap = map[string]string{
	"str":         "string",
	"int":         "number",
	"int64":       "number",
	"uint16":      "number",
	"uint":        "number",
	"bool":        "boolean",
	"interface{}": "any",
}

func Render(wr io.Writer, s string) error {
	type tplContext struct {
		Structs []*Struct
	}

	tpl, err := template.New("").Parse(s)
	if err != nil {
		return err
	}

	structs, err := Parse()
	if err != nil {
		return err
	}

	return tpl.Execute(wr, tplContext{
		Structs: structs,
	})
}

func Parse() ([]*Struct, error) {
	path := tdproto.SourceDir()

	fset := token.NewFileSet()
	d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	structs := make([]*Struct, 0)
	types := make(map[string]bool)

	for _, f := range d {
		files := make([]ast.File, 0, len(f.Files))
		for _, file := range f.Files {
			files = append(files, *file)
		}

		sort.Slice(files, func(i, j int) bool {
			return files[i].Name.Name < files[j].Name.Name
		})

		for _, f := range files {
			for _, decl := range f.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.TYPE {
					continue
				}

				s := Struct{
					Help:   cleanHelp(gen.Doc.Text()),
					Fields: make([]*Field, 0),
				}

				if s.Help == "" || strings.HasPrefix(s.Help, "deprecated") {
					continue
				}

				for _, spec := range gen.Specs {
					v, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					s.Name = v.Name.Name
					types[s.Name] = true

					fieldList := v.Type.(*ast.StructType).Fields
					for _, field := range fieldList.List {
						if field.Tag == nil {
							continue
						}

						tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
						jsontag := strings.Split(tag.Get("json"), ",")

						var typeNameBuf bytes.Buffer
						if err := printer.Fprint(&typeNameBuf, fset, field.Type); err != nil {
							return nil, err
						}

						nullable := false
						list := false

						typeDescr := typeNameBuf.String()
						if strings.HasPrefix(typeDescr, "*") {
							nullable = true
							typeDescr = typeDescr[1:]
						}

						if strings.HasPrefix(typeDescr, "[]") {
							list = true
							typeDescr = typeDescr[2:]
						}

						tsType := tsTypeMap[typeDescr]
						if tsType == "" {
							tsType = typeDescr
						}

						omitempty := len(jsontag) == 2 && jsontag[1] == "omitempty"

						if nullable && !omitempty && typeDescr == "JID" {
							nullable = false
						}

						name := field.Names[0].Name
						help := cleanHelp(field.Doc.Text())

						if help == "" {
							help = name
						}

						s.Fields = append(s.Fields, &Field{
							Name:      name,
							Help:      help,
							JSName:    strings.ToLower(name[:1]) + name[1:],
							TSType:    tsType,
							Json:      jsontag[0],
							Type:      typeDescr,
							Null:      nullable,
							List:      list,
							Omitempty: omitempty,
						})
					}

					break
				}

				structs = append(structs, &s)
			}
		}
	}

	for _, s := range structs {
		for _, f := range s.Fields {
			f.InternalType = types[f.Type]
		}
	}

	return structs, nil
}

func cleanHelp(s string) string {
	return strings.TrimSuffix(strings.TrimSpace(strings.Join(strings.Fields(s), " ")), ".")
}
