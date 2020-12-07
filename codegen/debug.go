package codegen

import (
	"bytes"
	"encoding/json"
	"log"

	"github.com/pkg/errors"
)

func DebugJSON(v interface{}) string {
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
