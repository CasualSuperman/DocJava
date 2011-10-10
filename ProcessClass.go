package main

import (
	"http"
	"io/ioutil"
	"strings"
	"sync"
	"time"
)

const (
	STARTCLASS = "<!-- ======== START OF CLASS DATA ======== -->"
	ENDCLASS   = "<!-- ========= END OF CLASS DATA ========= -->"
)

func (c *Class) Init(docUrl, className string, done *sync.WaitGroup) *Class {
	doc, _ := http.Get(docUrl)
	data, _ := ioutil.ReadAll(doc.Body)
	ClassHtml := string(data)
	if strings.Contains(className, ".") {
		// Subclass
		className = className[strings.LastIndex(className, ".")+1:]
	}
	ClassHtml = strings.Split(
		strings.Split(ClassHtml, STARTCLASS)[1],
		ENDCLASS,
	)[0]
	c.Name = className
	return c
}
