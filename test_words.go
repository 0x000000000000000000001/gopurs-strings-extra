package main

import (
	"fmt"
	"regexp"
)

func words(s string) []string {
	// A simple regex to extract words, matching unicode letters and digits
	// But it also needs to split camel case.
	// We can add spaces before uppercase letters, then split by non-alphanumeric.
	reSplitCamel := regexp.MustCompile(`([a-z])([A-Z])`)
	s = reSplitCamel.ReplaceAllString(s, `${1} ${2}`)
	
	reSplitWord := regexp.MustCompile(`[a-zA-Z0-9]+`)
	return reSplitWord.FindAllString(s, -1)
}

func main() {
	fmt.Printf("%q\n", words("Hello_world --from TheAliens"))
}
