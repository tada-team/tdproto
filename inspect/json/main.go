package main

import (
	"fmt"

	"github.com/tada-team/tdproto/inspect"
)

func main() {
	structs, err := inspect.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(inspect.DebugJSON(structs))
}
