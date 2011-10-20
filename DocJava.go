// DocJava is a program for generating Java code from a JavaDoc API webpage
// that will give the same result after JavaDoc as the webpage it was generated
// from.
package main
/* 
   Author: Bobby Wertman
   Version: 1.0.0
   Date: March 20, 2011
   This application and its source are provided 'as-is' and are available under
   the MIT License.
*/
import (
	"flag"
	"fmt"
	//	"http"
	"io/ioutil"
	//	"os"
	"regexp"
	"strconv"
	"strings"
	//	"sync"
)

var usage = "DocJava [-user USERNAME -pass PASSWORD] (-labID || -assignmentID) ID -semester SEMESTER -course COURSE"

var user *string = flag.String("user", "", "HTTP username required for the page.")
var pass *string = flag.String("pass", "", "HTTP password required for the page.")
var url *string = flag.String("url", "", "Address to the JavaDoc root.")
var labnum *int = flag.Int("labID", -1, "Lab number.")
var semester *string = flag.String("semester", "spring11", "The semester")
var course *string = flag.String("course", "cs151", "The course")
var assignment *int = flag.Int("assignmentID", -1, "Assignment number")

func init() {
	UrlReg = regexp.MustCompile("</?a[^>]*>")
}

func SplitClass(html string) []string {
	result := []string{}
	temp := regexp.MustCompile("<li><a href=[^>]+>([^<]+)</a>").FindAllStringSubmatch(strings.SplitN(
				strings.SplitN(
					html,
					"<ul class=\"subNavList\">\n<li>Detail:&nbsp;</li>\n",
					2)[1],
				"</ul>",
				2)[0], -1)
	sections := []string{}
	for _, val := range temp {
		sections = append(sections, val[1])
	}
	return result
}

func debugPrint(data ...string) {
	fmt.Println("{")
	for _, info := range data {
		fmt.Println("\t", info)
	}
	fmt.Println("}")
}

func javaDoc(j JavaDoc) (s string) {
	s = j.String()
	s = strings.Replace(s, "\n", "\n * ", -1)
	s = regexp.MustCompile(" * $").ReplaceAllString(s, " *")
	return
}

func main() {
	data, _ := ioutil.ReadFile("data.html")
	SplitClass(string(data))
}

func twoDigit(num int) string {
	result := ""
	if num < 10 {
		result = "0" + strconv.Itoa(num)
	} else if num >= 100 {
		result = twoDigit(num % 100)
	}
	return result
}

var UrlReg *regexp.Regexp

func RemoveUrl(s string) string {
	return UrlReg.ReplaceAllString(s, "")
}

func tab(s string, i int) string {
	return strings.Replace(s, "\n", "\n\t", -1)
}
