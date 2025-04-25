package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"log"
	"os"
	"strings"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func split(s string, size int) string {
	ss := make([]string, 0, len(s)/size+1)
	for len(s) > 0 {
		if len(s) < size {
			size = len(s)
		}
		ss, s = append(ss, s[:size]), s[size:]

	}
	return strings.Join(ss, "\n")
}

func isEmptyOrComment(l []byte) bool {
	if len(l) == 0 {
		return true
	}
	if l[0] == '#' {
		return true
	}
	return false
}

func main() {
	f, err := os.ReadFile("../blns.txt")
	if err != nil {
		log.Fatal(err)
	}
	fB64, err := os.OpenFile("../blns.base64.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer fB64.Close()

	log.Printf("%T", fB64)
	var blnsB64 []string
	var blns []string
	for l := range bytes.Lines(f) {
		l = bytes.TrimSpace(l)
		if isEmptyOrComment(l) {
			c, err := fB64.Write(append(l, 10))
			if err != nil {
				log.Print(c)
				log.Fatal(err)
			}
			continue
		}
		lB64 := base64.StdEncoding.EncodeToString(l)

		c, err := fB64.WriteString(split(lB64, 76) + "\n")
		if err != nil {
			log.Print(c)
			log.Fatal(err)
		}
		blnsB64 = append(blnsB64, lB64)
		blns = append(blns, string(l))
	}
	buf, err := json.MarshalIndent(blnsB64, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("../blns.base64.json", buf, 0644)
	if err != nil {
		log.Fatal(err)
	}
	buf, err = json.MarshalIndent(blns, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("../blns.json", buf, 0644)
}
