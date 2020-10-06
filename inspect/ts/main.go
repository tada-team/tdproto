package main

import (
	"fmt"
	"text/template"

	"github.com/tada-team/tdproto/inspect"
)

func main() {
	structs, err := inspect.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	template.Must(template.New("").Parse(`
`))

	for _, s := range structs {
		fmt.Println("//", s.Help)
		fmt.Println("export interface", s.Name, "{")
		//"  ", s.
		for _, f := range s.Fields {
			name := f.Name
			jsType := f.Type
			fmt.Printf("   /**\n")
			fmt.Printf("    * %s\n", f.Help)
			fmt.Printf("    */\n")
			fmt.Printf("   %s: %s;\n", name, jsType)
		}
		fmt.Println("}")
		fmt.Println()
	}
	// export interface Icon {
	//    height: number;
	//    width: number;
	//    url: string;
	//}
	//
	//export interface Icons {
	//    sm: Icon;
	//    lg: Icon;
	//    stub: string;
	//}
}
