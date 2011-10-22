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

func SplitClass(html string) []string {
	/* preamble
	 * nested_class
	 * nested_interface
	 * field
	 * constructor
	 * method
	 */
	menu := strings.SplitN(html, "<ul class=\"subNavList\">\n<li>Detail:&nbsp;</li>\n", 2)[1]
	list := strings.SplitN(menu, "</ul>", 2)[0]
	temp := regexp.MustCompile("<li><a href=[^>]+>([^<]+)</a>").FindAllStringSubmatch(list, -1)
	sections := []string{}
	for _, val := range temp {
		sections = append(sections, val[1])
	}
	return sections
}
