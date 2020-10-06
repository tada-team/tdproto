package inspect

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"reflect"
	"strings"

	"github.com/tada-team/tdproto"
)

type Field struct {
	Name      string `json:"name"`
	Help      string `json:"help"`
	Json      string `json:"json"`
	Type      string `json:"type"`
	Null      bool   `json:"null,omitempty"`
	Omitempty bool   `json:"omitempty,omitempty"`
}

type Struct struct {
	Name   string  `json:"name"`
	Help   string  `json:"help"`
	Fields []Field `json:"fields"`
}

func Parse() ([]Struct, error) {
	path := tdproto.SourceDir()

	fset := token.NewFileSet()
	d, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	structs := make([]Struct, 0)
	for k, f := range d {
		fmt.Println("package", k)
		for _, f := range f.Files {
			for _, decl := range f.Decls {
				gen, ok := decl.(*ast.GenDecl)
				if !ok || gen.Tok != token.TYPE {
					continue
				}

				s := Struct{
					Help:   cleanText(gen.Doc.Text()),
					Fields: make([]Field, 0),
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

					fieldList := v.Type.(*ast.StructType).Fields
					for _, field := range fieldList.List {
						tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
						jsontag := strings.Split(tag.Get("json"), ",")

						var typeNameBuf bytes.Buffer
						if err := printer.Fprint(&typeNameBuf, fset, field.Type); err != nil {
							return nil, err
						}

						nullable := false
						typeDescr := typeNameBuf.String()
						if strings.HasPrefix(typeDescr, "*") {
							nullable = true
							typeDescr = typeDescr[1:]
						}

						s.Fields = append(s.Fields, Field{
							Help:      cleanText(field.Doc.Text()),
							Name:      field.Names[0].Name,
							Json:      jsontag[0],
							Type:      typeDescr,
							Null:      nullable,
							Omitempty: len(jsontag) == 2 && jsontag[1] == "omitempty",
						})
					}

					break
				}

				structs = append(structs, s)
			}
		}

	}
	return structs, nil
}

func cleanText(s string) string {
	return strings.TrimSpace(strings.Join(strings.Fields(s), " "))
}
