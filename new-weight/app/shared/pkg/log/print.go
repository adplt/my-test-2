package log

import (
	"encoding/json"
	"fmt"
)

// PrintIndent godoc.
func PrintIndent(arg ...interface{}) {
	var err error
	var b []byte
	for _, i := range arg {
		b, err = json.MarshalIndent(i, "", " ")
		if err != nil {
			fmt.Print(string(b))
		}
	}
}

// PrintIndentln godoc.
func PrintIndentln(arg ...interface{}) {
	PrintIndent(arg...)
	fmt.Println()
}
