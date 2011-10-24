// DocJava is a program for generating Java code from a JavaDoc API webpage
// that will give the same result after JavaDoc as the webpage it was generated
// from.
package main
/* 
   Author: Bobby Wertman
   Version: 0.0.0
   Date: October 20, 2011
*/

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var parse *string = flag.String("f", "The file to parse.", "")

func main() {
	flag.Parse()
	data, err := ioutil.ReadFile(*parse)
	if err != nil {
		panic(err)
	}
	class := SplitClass(string(data))
	fmt.Println(NewClass(class).String())
}

func SplitClass(html string) (result []string) {
	result = []string{
		"", // Preamble
		"", // Nested_Class
		"", // Nested_Interface
		"", // Field
		"", // Constructor
		""} // Method
	menu := strings.SplitN(html, "<ul class=\"subNavList\">\n<li>Detail:&nbsp;</li>\n", 2)[1]
	list := strings.SplitN(menu, "</ul>", 2)[0]
	temp := regexp.MustCompile("<li><a href=[^>]+>([^<]+)</a>").FindAllStringSubmatch(list, -1)
	sections := []string{"START"}
	for _, val := range temp {
		sections = append(sections, val[1])
	}
	section := map[string]detail{
		"START":     detail{0, false},
		"Nested":    detail{1, false},
		"Interface": detail{2, false},
		"Field":     detail{3, true},
		"Constr":    detail{4, true},
		"Method":    detail{5, true}}
	sections = append(sections, "END")
	for i, val := range sections {
		defer func() {
			if x := recover(); x != nil {
				fmt.Println(x, val)
			}
		}()
		if val == "END" {
			continue
		}
		index := section[val].Index
		if section[val].Detail {
			result[index] = splitDetailSection(html, val, sections[i+1])
		} else {
			result[index] = splitSummarySection(html, val, sections[i+1])
		}
	}
	return result
}

type detail struct {
	Index  int
	Detail bool
}

func splitSummarySection(data, start, end string) (result string) {
	delimeters := map[string]string{
		"START":  "<!-- ======== START OF CLASS DATA ======== -->",
		"Nested": "<!-- ======== NESTED CLASS SUMMARY ======== -->",
		"Field":  "<!-- =========== FIELD SUMMARY =========== -->",
		"Constr": "<!-- ======== CONSTRUCTOR SUMMARY ======== -->",
		"Method": "<!-- ========== METHOD SUMMARY =========== -->",
		"END":    "<div class=\"details\">\n<ul class=\"blockList\">\n<li class=\"blockList\">"}
	return strings.Split(strings.Split(data, delimeters[start])[1], delimeters[end])[0]
}

func splitDetailSection(data, start, end string) string {
	delimeters := map[string]string{
		"START":  "<div class=\"details\">\n<ul class=\"blockList\">\n<li class=\"blockList\">",
		"Field":  "<!-- ============ FIELD DETAIL =========== -->\n<ul class=\"blockList\">\n<li class=\"blockList\"><a name=\"field_detail\">",
		"Constr": "<!-- ========= CONSTRUCTOR DETAIL ======== -->\n<ul class=\"blockList\">\n<li class=\"blockList\"><a name=\"constructor_detail\">",
		"Method": "<!-- ============ METHOD DETAIL ========== -->\n<ul class=\"blockList\">\n<li class=\"blockList\"><a name=\"method_detail\">",
		"END":    "</div>\n</div>\n<!-- ========= END OF CLASS DATA ========= -->"}
	return strings.Split(strings.Split(data, delimeters[start])[1], delimeters[end])[0]
}
