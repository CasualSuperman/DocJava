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
	a := NewField("<A NAME=\"head\"><!-- --></A><H3>\nhead</H3>\n<PRE>\nprivate <A HREF=\"../../../../../edu/wcu/cs/cs363/project02/BoundedLinkedSet.ListNode.html\" title=\"class in edu.wcu.cs.cs363.project02\">BoundedLinkedSet.ListNode</A> <B>head</B></PRE>\n<DL>\n<DD>The start of the linked list.\n<P>\n<DL>\n</DL>\n</DL>\n<HR>")
	b := NewField("<A NAME=\"capacity\"><!-- --></A><H3>\ncapacity</H3>\n<PRE>\nprivate int <B>capacity</B></PRE>\n<DL>\n<DD>The maximum number of elements to store in this\n <code>BoundedLinkedSet</code>.\n<P>\n<DL>\n</DL>\n</DL>")
	c := NewField("<A NAME=\"DEFAULT_CAPACITY_1\"><!-- --></A><H3>\nDEFAULT_CAPACITY_1</H3>\n<PRE>\nprotected static final int <B>DEFAULT_CAPACITY_1</B></PRE>\n<DL>\n<DD>The default capacity for <code>set1</code>.\n<P>\n<DL>\n<DT><B>See Also:</B><DD><A HREF=\"../../../../../constant-values.html#edu.wcu.cs.cs363.project02.AbstractBoundedSetTest.DEFAULT_CAPACITY_1\">Constant Field Values</A></DL>\n</DL>")
	d := NewField("<A NAME=\"set1\"><!-- --></A><H3>\nset1</H3>\n<PRE>\nprivate <A HREF=\"../../../../../edu/wcu/cs/cs363/project02/BoundedSet.html\" title=\"interface in edu.wcu.cs.cs363.project02\">BoundedSet</A>&lt;<A HREF=\"http://download.oracle.com/javase/6/docs/api/java/lang/String.html?is-external=true\" title=\"class or interface in java.lang\">String</A>&gt; <B>set1</B></PRE>\n<DL>\n<DD>A <code>BoundedSet</code> with which to test.\n<P>")
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
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
