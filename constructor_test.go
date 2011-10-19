package main

import "testing"

func TestNewConstructor(t *testing.T) {
	a := "<a name=\"BoundedLinkedSet(int)\">\n<!--   -->\n</a>\n<ul class=\"blockListLast\">\n<li class=\"blockList\">\n<h4>BoundedLinkedSet</h4>\n<pre>public&nbsp;BoundedLinkedSet(int&nbsp;capacity)</pre>\n<div class=\"block\">Creates a new, empty BoundedLinkedSet.</div>\n<dl><dt><span class=\"strong\">Parameters:</span></dt><dd><code>capacity</code> - The capacity of this newly created\n                <code>BoundedLinkedSet</code>.</dd>\n<dt><span class=\"strong\">Throws:</span></dt>\n<dd><code><a href=\"http://download.oracle.com/javase/6/docs/api/java/lang/IllegalArgumentException.html?is-external=true\" title=\"class or interface in java.lang\">IllegalArgumentException</a></code> - if capacity &lt;= 0.</dd></dl>\n</li>\n</ul>"
	b := "<a name=\"BoundedArraySet(int)\">\n<!--   -->\n</a>\n<ul class=\"blockListLast\">\n<li class=\"blockList\">\n<h4>BoundedArraySet</h4>\n<pre>public&nbsp;BoundedArraySet(int&nbsp;capacity)</pre>\n<div class=\"block\">Creates a new, empty BoundedArraySet.</div>\n<dl><dt><span class=\"strong\">Parameters:</span></dt><dd><code>capacity</code> - The capacity of this newly created\n			<code>BoundedArraySet</code>.</dd>\n<dt><span class=\"strong\">Throws:</span></dt>\n<dd><code><a href=\"http://download.oracle.com/javase/6/docs/api/java/lang/IllegalArgumentException.html?is-external=true\" title=\"class or interface in java.lang\">IllegalArgumentException</a></code> - if capacity &lt;= 0.</dd></dl>\n</li>\n</ul>"

	a_expected := "/**\n * Creates a new, empty BoundedLinkedSet.\n *\n * @param capacity The capacity of this newly created\n *                 <code>BoundedLinkedSet</code>.\n *\n * @throws IllegalArgumentException if capacity &lt;= 0.\n */\npublic BoundedLinkedSet(int capacity) {\n\n}"
	b_expected := "/**\n * Creates a new, empty BoundedArraySet.\n *\n * @param capacity The capacity of this newly created\n * <code>BoundedArraySet</code>.\n *\n * @throws IllegalArgumentException if capacity &lt;= 0.\n */\npublic BoundedArraySet(int capacity) {\n\n}"
	a_result := NewConstructor(a).String()
	b_result := NewConstructor(b).String()
	if a_result != a_expected {
		t.Errorf("Test failed on input a.\nGot:\n%v\nExpected:\n%v\n", a_result, a_expected)
	}

	if b_result != b_expected {
		t.Errorf("Test failed on input b.\nGot:\n%v\nExpected:\n%v\n", b_result, b_expected)
	}
}
