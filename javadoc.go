package main

import (
	"regexp"
	"strings"
)

type JavaDoc struct {
	description                                           string
	modeledBy, initEnsures, invariant                     string
	requires                                              []string
	corresponence                                         string
	version                                               string
	author, param, precondition, postcondition, preserves []string
	returns                                               string
	throws                                                []string
	example                                               string
}

func NewDoc(s string) (j JavaDoc) {
	j.description = regexp.MustCompile("^<div class=\"block\">(.*)</div>").FindStringSubmatch(s)[1]
	s = strings.SplitN(s, "</div>", 2)[1]
	s = regexp.MustCompile("</?span( class=\"strong\")?>").ReplaceAllString(s, "")
	return
}

func (j JavaDoc) String() (s string) {
	s += j.description
	return
}
