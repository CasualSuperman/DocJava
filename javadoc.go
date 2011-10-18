package main

import (
	"regexp"
	"strings"
)

type JavaDoc struct {
	description                       string
	modeledBy, initEnsures, invariant string
	requires                          []string
	correspondence, version           string
	author, param, precondition       []string
	postcondition, preserves          []string
	returns                           string
	throws                            []string
	example                           string
}

func NewDoc(s string) (j JavaDoc) {
	j.description = regexp.MustCompile("^<div class=\"block\">(.*)</div>").FindStringSubmatch(s)[1]
	s = strings.SplitN(s, "</div>", 2)[1]
	s = regexp.MustCompile("</?span( class=\"strong\")?>").ReplaceAllString(s, "")
	sets := [][]string{}
	temp := strings.Split(s, "</dd>")
	info := regexp.MustCompile("<dt>([^<]+):</dt>\n?<dd>(.+)</dd>")
	for i := 0; i < len(temp); i++ {
		sets = append(sets, info.FindStringSubmatch(temp[i]+"</dd>"))
	}
	strip := regexp.MustCompile("</?a[^>]*>")
	for i := 0; i < len(sets); i++ {
		if len(sets[i]) > 0 {
			sets[i][2] = strip.ReplaceAllString(sets[i][2], "")
			debugPrint(sets[i]...)
		} else {
			sets[i] = sets[len(sets)-1]
			sets = sets[:len(sets)-1]
			i--
		}
	}
	return
}

func (j JavaDoc) String() (s string) {
	s += j.description
	if j.modeledBy != "" {
		s += "\n"
		s += "\n@modeledby " + j.modeledBy
	}
	if j.initEnsures != "" {
		s += "\n"
		s += "\n@initEnsures " + j.initEnsures
	}
	if j.invariant != "" {
		s += "\n"
		s += "\n@invariant " + j.invariant
	}
	if len(j.requires) > 0 {
		s += "\n"
		for i := 0; i < len(j.requires); i++ {
			s += "\n@requires " + j.requires[i]
		}
	}
	if j.correspondence != "" {
		s += "\n"
		s += "\n@correspondence " + j.correspondence
	}
	if j.version != "" {
		s += "\n"
		s += "\n@version " + j.version
	}
	if len(j.author) > 0 {
		s += "\n"
		for i := 0; i < len(j.author); i++ {
			s += "\n@author " + j.author[i]
		}
	}
	if len(j.author) > 0 {
		s += "\n"
		for i := 0; i < len(j.author); i++ {
			s += "\n@author " + j.author[i]
		}
	}
	if len(j.param) > 0 {
		s += "\n"
		for i := 0; i < len(j.param); i++ {
			s += "\n@param " + j.author[i]
		}
	}
	if len(j.precondition) > 0 {
		s += "\n"
		for i := 0; i < len(j.precondition); i++ {
			s += "\n@precondition " + j.precondition[i]
		}
	}
	if len(j.postcondition) > 0 {
		s += "\n"
		for i := 0; i < len(j.postcondition); i++ {
			s += "\n@postcondition " + j.postcondition[i]
		}
	}
	if len(j.preserves) > 0 {
		s += "\n"
		for i := 0; i < len(j.preserves); i++ {
			s += "\n@postcondition " + j.preserves[i]
		}
	}
	if j.returns != "" {
		s += "\n"
		s += "\n@return " + j.returns
	}
	if len(j.throws) > 0 {
		s += "\n"
		for i := 0; i < len(j.throws); i++ {
			s += "\n@postcondition " + j.throws[i]
		}
	}
	if j.example != "" {
		s += "\n"
		s += "\n@example " + j.example
	}
	return
}
