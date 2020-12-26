// Utility to provide indentation to compressed json fragment
// uses clipboard as both: source and destination
// to avoid console creation on windows compile:
// go build -ldflags -H=windowsgui
//

package main

import (
	// stdlib
	"encoding/json"
	"fmt"

	// third party packages
	"github.com/atotto/clipboard"
)

func main() {
	run()
}

func run() {
	// read from clipboard
	s, _ := clipboard.ReadAll()
	// un-marshall json
	var v interface{}
	err := json.Unmarshal([]byte(s), &v)
	if err != nil {
		serr := fmt.Sprintf("%s", err)
		clipboard.WriteAll(serr)
		return
	}
	// marshal json with indentation
	jss, _ := json.MarshalIndent(v, "", "\t")
	// write to clipboard
	clipboard.WriteAll(string(jss))

}
