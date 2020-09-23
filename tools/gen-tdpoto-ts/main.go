package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/tada-team/tdproto"
)

type Field struct {
	Doc  string
	Name string
}

type Struct struct {
	Doc    string
	Name   string
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
					Doc:    strings.TrimSpace(gen.Doc.Text()),
					Fields: make([]Field, 0),
				}

				if s.Doc == "" || strings.HasPrefix(s.Doc, "deprecated") {
					continue
				}

				for _, spec := range gen.Specs {
					v, ok := spec.(*ast.TypeSpec)
					if !ok {
						continue
					}
					s.Name = v.Name.Name
					break
				}
				structs = append(structs, s)
			}
		}

	}
	return structs, nil
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
