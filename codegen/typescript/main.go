package main

import (
	"fmt"
	"github.com/tada-team/tdproto/codegen"
	"os"
	"text/template"
)

const TypeScriptIndent = "  "
const TypeScriptClassHeaderTemplateStr = (`export class {{.Name}} implements TDProtoClass<{{.Name}}> {`)
const TypeScriptClassFieldTemplateStr = (`{{.Name}}{{if .IsOmitEmpty}} ? {{endif}}: {{ $data.TypeStr }}`)
const TypeScriptClassConstructorField = "constructor ("
const TypeScriptConstructorClose = ") {}"

const TypeScriptClosingBracket = "{"
const TypeScriptInterfaceHeaderTemplateStr = `export interface {{.Name}}JSON {`

type TypeScriptTemplate struct {
	ClassHeader           *template.Template
	InterfaceOrClassField *template.Template
	InterfaceHeader       *template.Template
}

func generateTypeScript(tdprotoInfo *codegen.TadaInfo) []string {
	var typeScriptLines []string

	return typeScriptLines
}

func main() {
	tdprotoInfo, err := codegen.ParseTdproto()

	if err != nil {
		panic(fmt.Sprintf("%s", err))
	}

	for _, line := range generateTypeScript(tdprotoInfo) {
		fmt.Fprintf(os.Stdout, "%s\n", line)
	}

}
