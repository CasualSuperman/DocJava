package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

// Used for Stripping out hyperlinks
var UrlReg *regexp.Regexp = regexp.MustCompile("</?a[^>]*>")

var tabs bool

func init() {
	flag.BoolVar(&tabs, "t", false, "Use tabs instead of spaces.")
}

func RemoveUrl(s string) string {
	return UrlReg.ReplaceAllString(s, "")
}

// Used for properly tabbing nested things
func tab(s string, i int) (result string) {
	if tabs {
		result = strings.Replace(s, "\n", "\n\t", -1)
	} else {
		result = strings.Replace(s, "\n", "\n    ", -1)
	}
	return
}

// Used for adding *'s to included newlines
func javaDoc(j JavaDoc) (s string) {
	s = j.String()
	s = strings.Replace(s, "\n", "\n * ", -1)
	s = regexp.MustCompile(" * $").ReplaceAllString(s, " *")
	return
}

// This is used for testing only.
// No calls to it should exist in a final product.
func debugPrint(data ...string) {
	fmt.Println("{")
	for _, info := range data {
		fmt.Println("\t", info)
	}
	fmt.Println("}")
}
