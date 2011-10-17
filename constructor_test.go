package main

import "testing"

func TestNewConstructor(t *testing.T) {
	a := "<a name=\"BoundedLinkedSet(int)\">\n<!--   -->\n</a>\n<ul class=\"blockListLast\">\n<li class=\"blockList\">\n<h4>BoundedLinkedSet</h4>\n<pre>public&nbsp;BoundedLinkedSet(int&nbsp;capacity)</pre>\n<div class=\"block\">Creates a new, empty BoundedLinkedSet.</div>\n<dl><dt><span class=\"strong\">Parameters:</span></dt><dd><code>capacity</code> - The capacity of this newly created\n                <code>BoundedLinkedSet</code>.</dd>\n<dt><span class=\"strong\">Throws:</span></dt>\n<dd><code><a href=\"http://download.oracle.com/javase/6/docs/api/java/lang/IllegalArgumentException.html?is-external=true\" title=\"class or interface in java.lang\">IllegalArgumentException</a></code> - if capacity &lt;= 0.</dd></dl>\n</li>\n</ul>"
	//	b := ""
	//	c := ""
	//	d := ""

	a_expected := "/**\n * Creates a new, empty BoundedLinkedSet.\n *\n * @param capacity The capacity of this newly created <code>BoundedLinkedSet</code>.\n *\n * @throws IllegalArgumentException if capacity &lt;= 0.\n */\npublic BoundedLinkedSet(int capacity) {\n}"
	//	b_expected := ""
	//	c_expected := ""
	//	d_expected := ""
	a_result := NewConstructor(a).String()
	if a_result != a_expected {
		t.Errorf("Test failed on input a.\nGot:\n%v\nExpected:\n%v\n", a_result, a_expected)
	}
	/*
		if NewConstructor(b).String() != b_expected {
			t.Error("Test failed on input:\n", b, "\nImproper output.")
		}

		if NewConstructor(c).String() != c_expected {
			t.Error("Test failed on input:\n", c, "\nImproper output.")
		}

		if NewConstructor(d).String() != d_expected {
			t.Error("Test failed on input:\n", d, "\nImproper output.")
		}
	*/
}
