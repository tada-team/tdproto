package main

import (
	"fmt"
	"os"

	"github.com/tada-team/tdproto/codegen"
)

func main() {
	if err := codegen.Render(os.Stdout, `## Structures
Autogenerated from source code. Do not edit.
{{range $s := .Structs}}
### <a name="{{$s.Name}}"></a>{{$s.Name}}
{{$s.Help}}.
{{range $f := $s.Fields }}
 * **{{$f.Json}}** ({{ if $f.InternalType -}}
		[{{$f.Type}}](#{{$f.Type}})
	{{- else -}}
		{{$f.Type}}
	{{- end -}}
	{{- if $f.Readonly }}, readonly for clients{{ end -}}
	{{- if $f.Null }}, nullable{{ end -}}
	{{- if $f.List }}, list{{end -}}
	{{- if $f.Omitempty }}, omitempty{{ end }}) — {{$f.Help}}.
{{end}}{{range $v := $s.EnumValues }}
 * **{{ $v.Value }}** {{ $v.Help }}{{end}}
{{end}}
`); err != nil {
		fmt.Println(err)
	}
}
