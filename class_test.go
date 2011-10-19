package main

import "testing"

func TestNewClass(t *testing.T) {
	// NewClass is formatted like so:
	// 0) preamble
	// 1) nested_class
	// 2) nested_interface
	// 3) field
	// 4) constructor
	// 5) method
	a_result := NewClass(
		"<div class=\"header\">\n<div class=\"subTitle\">edu.wcu.cs.cs363.project02</div>\n<h2 title=\"Class BoundedLinkedSet\" class=\"title\">Class BoundedLinkedSet&lt;T&gt;</h2>\n</div>\n<div class=\"contentContainer\">\n<ul class=\"inheritance\">\n<li><a href=\"http://download.oracle.com/javase/6/docs/api/java/lang/Object.html?is-external=true\" title=\"class or interface in java.lang\">java.lang.Object</a></li>\n<li>\n<ul class=\"inheritance\">\n<li><a href=\"../../../../../edu/wcu/cs/cs363/project02/AbstractBoundedSet.html\" title=\"class in edu.wcu.cs.cs363.project02\">edu.wcu.cs.cs363.project02.AbstractBoundedSet</a>&lt;T&gt;</li>\n<li>\n<ul class=\"inheritance\">\n<li>edu.wcu.cs.cs363.project02.BoundedLinkedSet&lt;T&gt;</li>\n</ul>\n</li>\n</ul>\n</li>\n</ul>\n<div class=\"description\">\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<dl>\n<dt>All Implemented Interfaces:</dt>\n<dd><a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedSet.html\" title=\"interface in edu.wcu.cs.cs363.project02\">BoundedSet</a>&lt;T&gt;, <a href=\"http://download.oracle.com/javase/6/docs/api/java/lang/Cloneable.html?is-external=true\" title=\"class or interface in java.lang\">Cloneable</a></dd>\n</dl>\n<hr>\n<br>\n<pre>public class <span class=\"strong\">BoundedLinkedSet&lt;T&gt;</span>\nextends <a href=\"../../../../../edu/wcu/cs/cs363/project02/AbstractBoundedSet.html\" title=\"class in edu.wcu.cs.cs363.project02\">AbstractBoundedSet</a>&lt;T&gt;</pre>\n<div class=\"block\">This class provides a linked-based bounded set implementation.</div>\n<dl><dt><span class=\"strong\">Correspondence:</span></dt>\n  <dd>self.contents =&gt; items in linked list pointed to by\n                   head\n                 and self.capacity =&gt; this.capacity</dd>\n<dt><span class=\"strong\">Version:</span></dt>\n  <dd>October 7, 2011</dd>\n<dt><span class=\"strong\">Author:</span></dt>\n  <dd>Dr. Dalton</dd></dl>\n</li>\n</ul>\n</div>\n<div class=\"summary\">\n<ul class=\"blockList\">\n<li class=\"blockList\">",
		"",
		"",
		"",
		"",
		"").String()
	a_expected := "/**\n * This class provides a linked-based bounded set implementation.\n *\n * @correspondence: self.contents =&gt; items in linked list pointed to by\n *	                     head\n *                  and self.capacity =&gt; this.capacity\n *\n * @version Octover 7, 2011\n *\n * @author Dr. Dalton\n */\npublic class BoundedLinkedSet<T> extends AbstracBoundedSet<T> {\n}"

	if a_result != a_expected {
		t.Errorf("Test failed on input a.\nGot:\n%v\nExpected:\n%v\n", a_result, a_expected)
	}

}
