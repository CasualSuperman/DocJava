// DocJava is a program for generating Java code from a JavaDoc API webpage
// that will give the same result after JavaDoc as the webpage it was generated
// from.
package main
/* 
   Author: Bobby Wertman
   Version: 0.0.0
   Date: October 20, 2011
   This application and its source are provided 'as-is' and are available under
   the MIT License.
*/
import (
	"io/ioutil"
	"regexp"
	"strings"
)

func main() {
	data, _ := ioutil.ReadFile("_test_data.html")
	debugPrint(SplitClass(string(data))...)
}

func SplitClass(html string) []string {
	menu := strings.SplitN(html, "<ul class=\"subNavList\">\n<li>Detail:&nbsp;</li>\n", 2)[1]
	list := strings.SplitN(menu, "</ul>", 2)[0]
	temp := regexp.MustCompile("<li><a href=[^>]+>([^<]+)</a>").FindAllStringSubmatch(list, -1)
	sections := []string{}
	for _, val := range temp {
		sections = append(sections, val[1])
	}
	return sections
}
