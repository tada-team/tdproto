package codegen

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/pkg/errors"
)

var errorLogger = log.New(os.Stderr, "", 0)

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
