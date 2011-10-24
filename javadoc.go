package main

import (
	"fmt"
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
	if s == "" {
		return
	}
	j.description = strings.Replace(regexp.MustCompile("^<div class=\"block\">(.*)</div>").FindStringSubmatch(s)[1], "\n", "\n * ", -1)
	s = strings.SplitN(s, "</div>", 2)[1]
	s = regexp.MustCompile("</?span( class=\"strong\")?>").ReplaceAllString(s, "")
	sets := [][]string{}
	temp := strings.Split(s, "</dd>")
	info := regexp.MustCompile("<dt>([^<]+):</dt>\n? *<dd>(.+)</dd>")

	for _, val := range temp {
		sets = append(sets, info.FindStringSubmatch(val+"</dd>"))
	}
	strip := regexp.MustCompile("</?a[^>]*>")
	for i := 0; i < len(sets); i++ {
		if len(sets[i]) > 0 {
			sets[i][2] = strip.ReplaceAllString(sets[i][2], "")
			j.AddInfo(sets[i][1:])
		}
	}
	return
}

func (j *JavaDoc) AddInfo(info []string) {
	info[1] = strings.Trim(regexp.MustCompile("\n\t*").ReplaceAllString(info[1], "\n"), " ")
	info[1] = strings.Replace(info[1], "\n", "\n * ", -1)
	switch info[0] {
	case "Modeled By":
		j.modeledBy = info[1]
	case "Initialization Ensures":
		j.initEnsures = info[1]
	case "Representation Invariant":
		j.invariant = info[1]
	case "Requires":
		j.requires = append(j.requires, info[1])
	case "Correspondence":
		j.correspondence = info[1]
	case "Version":
		j.version = info[1]
	case "Author":
		j.author = append(j.author, info[1])
	case "Parameters":
		info[1] = strings.Replace(strings.Replace(info[1], "<code>", "", 1), "</code> - ", " ", 1)
		j.param = append(j.param, info[1])
	case "Preconditions":
		j.precondition = append(j.precondition, info[1])
	case "Postconditions":
		j.postcondition = append(j.postcondition, info[1])
	case "Preserves":
		j.preserves = append(j.preserves, info[1])
	case "Returns":
		j.returns = info[1]
	case "Throws":
		info[1] = strings.Replace(strings.Replace(info[1], "<code>", "", 1), "</code> - ", " ", 1)
		j.throws = append(j.throws, info[1])
	case "Example":
		j.example = info[1]
	case "See Also":
		// Go doesn't fall through by default, so this actually does nothing.
	default:
		fmt.Println("Woah, what's this?")
		debugPrint(append([]string{""}, info...)...)
	}
}

func (j JavaDoc) String() (s string) {
	s += "\n/**"
	s += "\n * " + j.description
	if j.modeledBy != "" {
		s += "\n *"
		s += "\n *@modeledby " + j.modeledBy
	}
	if j.initEnsures != "" {
		s += "\n *"
		s += "\n * @initEnsures " + j.initEnsures
	}
	if j.invariant != "" {
		s += "\n *"
		s += "\n * @invariant " + j.invariant
	}
	if len(j.requires) > 0 {
		s += "\n *"
		for _, val := range j.requires {
			s += "\n * @requires " + val
		}
	}
	if j.correspondence != "" {
		s += "\n *"
		s += "\n * @correspondence " + j.correspondence
	}
	if j.version != "" {
		s += "\n *"
		s += "\n * @version " + j.version
	}
	if len(j.author) > 0 {
		s += "\n *"
		for _, val := range j.author {
			s += "\n * @author " + val
		}
	}
	if len(j.param) > 0 {
		s += "\n *"
		for _, val := range j.param {
			s += "\n * @param " + val
		}
	}
	if len(j.precondition) > 0 {
		s += "\n *"
		for _, val := range j.precondition {
			s += "\n * @precondition " + val
		}
	}
	if len(j.postcondition) > 0 {
		s += "\n *"
		for _, val := range j.postcondition {
			s += "\n * @postcondition " + val
		}
	}
	if len(j.preserves) > 0 {
		s += "\n *"
		for _, val := range j.preserves {
			s += "\n * @preserves " + val
		}
	}
	if j.returns != "" {
		s += "\n *"
		s += "\n * @return " + j.returns
	}
	if len(j.throws) > 0 {
		s += "\n *"
		for _, val := range j.throws {
			s += "\n * @throws " + val
		}
	}
	if j.example != "" {
		s += "\n *"
		s += "\n * @example " + j.example
	}
	s += "\n */"
	if s == "\n/**\n * "+j.description+"\n */" && !strings.Contains(j.description, "\n") {
		s = "\n/** " + j.description + " */"
	}
	if s == "\n/**  */" {
		s = ""
	}
	return
}
