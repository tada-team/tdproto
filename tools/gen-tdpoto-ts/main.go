package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
	"reflect"
	"strings"

	"github.com/pkg/errors"
	"github.com/tada-team/tdproto"
)

type Field struct {
	Name      string
	Help      string
	Tag       string
	Type      string
	Omitempty bool
}

type Struct struct {
	Name   string
	Help   string
	Fields []Field
}

func main() {
	structs, err := parse(tdproto.SourceDir())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(debugJSON(structs))
}

func parse(path string) ([]Struct, error) {
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
						json := strings.Split(tag.Get("json"), ",")
						var typeNameBuf bytes.Buffer
						printer.Fprint(&typeNameBuf, fset, field.Type)
						s.Fields = append(s.Fields, Field{
							Help:      cleanText(field.Doc.Text()),
							Name:      field.Names[0].Name,
							Tag:       json[0],
							Type:      typeNameBuf.String(),
							Omitempty: len(json) == 2 && json[1] == "omitempty",
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

func debugJSON(v interface{}) string {
	b := new(bytes.Buffer)
	debugEncoder := json.NewEncoder(b)
	debugEncoder.SetIndent("", "    ")
	debugEncoder.SetEscapeHTML(false)
	err := debugEncoder.Encode(v)
	if err != nil {
		log.Panicln(errors.Wrap(err, "json marshall fail"))
	}
	return b.String()
}
