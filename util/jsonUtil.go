package util

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// JSON出力整形
func JSONFormatOut(jsonString []byte) {

	out := new(bytes.Buffer)
	// prefixなし、スペース4つでインデント
	json.Indent(out, jsonString, "", "    ")
	fmt.Println("jsonString=", out.String())
}
