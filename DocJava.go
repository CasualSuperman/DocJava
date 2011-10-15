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
	//	"io/ioutil"
	//	"os"
	//	"regexp"
	"strconv"
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

/*
func main() {
	flag.Parse()
	var url string
	if *user != "" && *pass != "" {
		url = "http://" + *user + ":" + *pass + "@"
	} else if u, p := (*user != ""), (*pass != ""); !(u && p) && !(!u && !p) {
		fmt.Println(usage)
		fmt.Println("You must either specify both a username and a password or neither.")
		os.Exit(1)
	}
	url += "polaris.cs.wcu.edu/" +
		"~adalton/teaching/" + *semester + "/" + *course
	if *labnum > 0 {
		url += "/labs/lab" + twoDigit(*labnum)
	} else if *assignment > 0 {
		url += "/assignments/assignment" + twoDigit(*assignment)
	} else {
		fmt.Println(usage)
		fmt.Println("You must specify either an assignment or a lab.")
		os.Exit(1)
	}
	url += "docs/"
	classUrl := url + "allclasses-noframe.html"
	classHTTP, err := http.Get(classUrl)
	if err != nil {
		fmt.Println("The requested assignment could not be retreived.")
		fmt.Println(err)
		os.Exit(1)
	}
	classHTML, _ := ioutil.ReadAll(classHTTP.Body)
	reg, _ := regexp.Compile("A HREF=\"([^\"+)\" [^>]+>([^<+])</A>")
	classTemp := reg.FindAllStringSubmatch(string(classHTML), -1)
	classPairs := make([][2]string, len(classTemp))
	for i, class := range classTemp {
		for j, val := range class {
			if j != 0 {
				classPairs[i][j-1] = val
			}
		}
	}
	os.Mkdir("src", 0664)
	os.Chdir("src")
	//	classes := make([]*Class, len(classPairs))
	done := new(sync.WaitGroup)
	for i, val := range classPairs {
		done.Add(1)
		if i == 0 || len(val) == 0 {
		}
		//		classes[i] = new(Class).Init(url+val[0], val[1], done)
	}
	done.Wait()
}
*/

func main() {
	a := NewField("<a name=\"DEFAULT_CAPACITY_1\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>DEFAULT_CAPACITY_1</h4>\n<pre>protected static final&nbsp;int DEFAULT_CAPACITY_1</pre>\n<div class=\"block\">The default capacity for <code>set1</code>.</div>\n<dl><dt><span class=\"strong\">See Also:</span></dt><dd><a href=\"../../../../../constant-values.html#edu.wcu.cs.cs363.project02.AbstractBoundedSetTest.DEFAULT_CAPACITY_2\">Constant Field Values</a></dd></dl>\n</li>\n</ul>")
	fmt.Println(a)
	b := NewField("")
	fmt.Println(b)
	c := NewField("")
	fmt.Println(c)
	d := NewField("")
	fmt.Println(d)
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
