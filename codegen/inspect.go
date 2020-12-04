package codegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"reflect"
	"sort"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"
	"github.com/tada-team/tdproto"
)

type EnumValue struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Help  string `json:"help"`
}

func (e EnumValue) DartValue() string { return strings.ReplaceAll(e.Value, `"`, `'`) }
func (e EnumValue) DartName() string  { return strcase.ToLowerCamel(e.Name) }

type Field struct {
	Name         string `json:"name"`
	JSName       string `json:"js_name"`
	Help         string `json:"help"`
	Json         string `json:"json"`
	Type         string `json:"type"`
	Readonly     bool   `json:"readonly"`
	Null         bool   `json:"null,omitempty"`
	List         bool   `json:"list,omitempty"`
	InternalType bool   `json:"internal_type,omitempty"`
	Omitempty    bool   `json:"omitempty,omitempty"`
	TSType       string `json:"ts_type"`
	DartType     string `json:"dart_type"`
	TSDefault    string `json:"ts_default"`
}

func (f Field) DartRequired() bool {
	return !(f.Null || f.Omitempty)
}

type Struct struct {
	Name       string      `json:"name"`
	Help       string      `json:"help"`
	Fields     []*Field    `json:"fields"`
	Readonly   bool        `json:"readonly,omitempty"`
	EnumValues []EnumValue `json:"enum_values,omitempty"`
}

func (s Struct) SnakeName() string { return strcase.ToSnake(s.Name) }
func (s Struct) IsEnum() bool      { return len(s.EnumValues) > 0 }

var tsTypeMap = map[string]string{
	"string":            "string",
	"int":               "number",
	"int64":             "number",
	"uint16":            "number",
	"uint":              "number",
	"bool":              "boolean",
	"interface{}":       "any",
	"ISODateTimeString": "string",
}

var dartTypeMap = map[string]string{
	"string":            "String",
	"int":               "int",
	"int64":             "int",
	"uint16":            "int",
	"uint":              "int",
	"bool":              "bool",
	"interface{}":       "dynamic",
	"ISODateTimeString": "DateTime",
}

var tsDefaultMap = map[string]string{
	"string":       `''`,
	"number":       `0`,
	"boolean":      `false`,
	"MessageLinks": `[]`,
	"Mediasubtype": `''`,
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

func Parse() (structs []*Struct, err error) {
	fset := token.NewFileSet()
	enumsMap := make(map[string][]EnumValue)
	if err := doParse(fset, token.CONST, func(gen *ast.GenDecl) error {
		for _, spec := range gen.Specs {
			valueSpec, ok := spec.(*ast.ValueSpec)
			if !ok || valueSpec.Type == nil {
				continue
			}

			typeName := fmt.Sprintf("%s", valueSpec.Type) // XXX:
			if typeName == "" {
				continue
			}

			name := valueSpec.Names[0]
			val, ok := valueSpec.Values[0].(*ast.BasicLit)
			if !ok {
				continue
			}

			if enumsMap[typeName] == nil {
				enumsMap[typeName] = []EnumValue{}
			}
			enumsMap[typeName] = append(enumsMap[typeName], EnumValue{
				Name:  name.Name,
				Value: val.Value,
				Help:  cleanHelp(valueSpec.Doc.Text()),
			})
		}
		return nil
	}); err != nil {
		return nil, err
	}

	types := make(map[string]struct{})
	if err := doParse(fset, token.TYPE, func(gen *ast.GenDecl) error {
		s := Struct{Help: cleanHelp(gen.Doc.Text())}
		if s.Help == "" || strings.HasPrefix(strings.ToLower(s.Help), "deprecated") {
			return nil
		}

		s.Readonly = strings.Contains(s.Help, "Readonly") // XXX
		for _, spec := range gen.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			s.Name = typeSpec.Name.Name
			types[s.Name] = struct{}{}

			s.EnumValues = enumsMap[s.Name]

			switch typeSpec.Type.(type) {
			case *ast.StructType:
				fieldList := typeSpec.Type.(*ast.StructType).Fields
				for _, field := range fieldList.List {
					name := field.Names[0].Name

					omitempty := false
					jsonName := name
					readonly := false
					if field.Tag != nil {
						tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
						jsontag := strings.Split(tag.Get("json"), ",")
						omitempty = len(jsontag) == 2 && jsontag[1] == "omitempty"
						jsonName = jsontag[0]

						if jsonName == "-" {
							continue
						}

						for _, tag := range strings.Split(tag.Get("tdproto"), ",") {
							if tag == "readonly" {
								readonly = true
							}
						}
					}

					var typeNameBuf bytes.Buffer
					if err := printer.Fprint(&typeNameBuf, fset, field.Type); err != nil {
						return err
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

					dartType := dartTypeMap[typeDescr]
					if dartType == "" {
						dartType = typeDescr
					}

					if nullable && !omitempty && typeDescr == "JID" { // XXX:
						nullable = false
					}

					help := cleanHelp(field.Doc.Text())

					if help == "" {
						help = name
					}

					tsDefault := ""
					if omitempty {
						if nullable {
							tsDefault = "null"
						} else {
							tsDefault = tsDefaultMap[tsType]
							if tsDefault == "" {
								tsDefault = "unknown: " + tsType
							}
						}
					}

					s.Fields = append(s.Fields, &Field{
						Name:      name,
						Help:      help,
						JSName:    strings.ToLower(name[:1]) + name[1:],
						TSType:    tsType,
						TSDefault: tsDefault,
						DartType:  dartType,
						Json:      jsonName,
						Type:      typeDescr,
						Null:      nullable,
						List:      list,
						Omitempty: omitempty,
						Readonly:  readonly,
					})
				}
			}

			structs = append(structs, &s)
		}
		return nil
	}); err != nil {
		return nil, err
	}

	sort.Slice(structs, func(i, j int) bool {
		return structs[i].Name < structs[j].Name
	})

	for _, s := range structs {
		for _, f := range s.Fields {
			_, f.InternalType = types[f.Type]
		}
	}

	return structs, nil
}

func doParse(fset *token.FileSet, tok token.Token, fn func(gen *ast.GenDecl) error) error {
	path := tdproto.SourceDir()
	d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return err
	}
	for _, f := range d {
		files := make([]ast.File, 0, len(f.Files))
		for _, file := range f.Files {
			files = append(files, *file)
		}
		for _, f := range files {
			for _, decl := range f.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok {
					continue
				}
				if gen.Tok == tok {
					if err := fn(gen); err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func cleanHelp(s string) string {
	return strings.TrimSuffix(strings.TrimSpace(strings.Join(strings.Fields(s), " ")), ".")
}
