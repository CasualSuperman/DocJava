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
	"io/ioutil"
	"regexp"
	"strings"
)

var parse *string = flag.String("f", "The file to parse.", "")

func main() {
	flag.Parse()
	data, _ := ioutil.ReadFile(*parse)
	debugPrint(SplitClass(string(data))...)
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
	section := map[string]int{
		"Field":  3,
		"Constr": 4,
		"Method": 5}
	sections = append(sections, "END")
	for i, val := range sections {
		if val == "START" || val == "END" {
			continue
		}
		result[section[val]] = splitSection(html, sections[i-1], val)
	}
	return result
}

func splitSection(data, first, second string) (result string) {
	delimeters := map[string]string{
		"START":  "<div class=\"details\">\n<ul class=\"blockList\">\n<li class=\"blockList\">",
		"Field":  "<!-- ============ FIELD DETAIL =========== -->",
		"Constr": "<!-- ========= CONSTRUCTOR DETAIL ======== -->",
		"Method": "<!-- ============ METHOD DETAIL ========== -->",
		"END":    "<!-- ========= END OF CLASS DATA ========= -->"}
	result = strings.Split(strings.Split(data, delimeters[first])[1], delimeters[second])[0]
	return
}
