package main

import (
	"testing"

	"github.com/atotto/clipboard"
)

func TestOK(t *testing.T) {
	clipboard.WriteAll(`{"header": {"precomputed": true},"body": "Hello Gophers!"}`)
	run()
	indented, _ := clipboard.ReadAll()
	result :=
		"{\n" +
			"	\"body\": \"Hello Gophers!\",\n" +
			"	\"header\": {\n" +
			"		\"precomputed\": true\n" +
			"	}\n" +
			"}"

	if indented != result {
		t.Errorf("Wrong json after format:\n%s\nExpected Length:%d\nActual Length:%d\n",
			indented, len(result), len(indented))
	}
}

func TestErr(t *testing.T) {
	clipboard.WriteAll(`{"header": {"precomputed": true},"body": "Hello Gophers!"`)
	run()
	indented, _ := clipboard.ReadAll()
	result := "unexpected end of JSON input"
	if indented != result {
		t.Errorf("Not properly detected json error:\n%s\nExpected Length:%d\nActual Length:%d\n",
			indented, len(result), len(indented))
	}
}
