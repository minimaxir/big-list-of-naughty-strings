// Package naughtystrings is a collection of strings that have a high probability of causing issues when used as user input.
package naughtystrings

import (
	"encoding/json"
	"os"
	"path"
	"runtime"
)

// Base64Encoded returns the strings encoded in base 64.
func Base64Encoded() []string {
	return strings("blns.base64.json")
}

// Unencoded returns the strings.
func Unencoded() []string {
	return strings("blns.json")
}

func strings(file string) []string {
	var _, here, _, _ = runtime.Caller(0)

	there, err := os.Open(path.Join(path.Dir(path.Dir(here)), file))

	if err != nil {
		panic(err)
	}

	var strings []string

	if err := json.NewDecoder(there).Decode(&strings); err != nil {
		panic(err)
	}

	return strings
}
