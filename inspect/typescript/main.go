package main

import (
	"fmt"
	"os"

	"github.com/tada-team/tdproto/inspect"
)

func main() {
	if err := inspect.Render(os.Stdout, `// Autogenerated. Do not edit.

type JID = string;  {{- range $s := .Structs}}

/**
 * {{$s.Help}}.
 */
export interface {{$s.Name}} { {{- range $f := $s.Fields }}
   /**
    * {{$f.Help}}.{{ if $f.Omitempty }} Omitempty.{{ end }}
    */
   {{$f.JSName}}: {{$f.TSType}}{{ if $f.List }}[]{{end}}{{ if $f.Null }} | null{{ end }};
{{end}}
}
{{end}}`); err != nil {
		fmt.Println(err)
	}
}
