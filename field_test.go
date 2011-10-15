package main

import "testing"

func TestNewField(t *testing.T) {
	a := NewField("<a name=\"DEFAULT_CAPACITY_1\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>DEFAULT_CAPACITY_1</h4>\n<pre>protected static final&nbsp;int DEFAULT_CAPACITY_1</pre>\n<div class=\"block\">The default capacity for <code>set1</code>.</div>\n<dl><dt><span class=\"strong\">See Also:</span></dt><dd><a href=\"../../../../../constant-values.html#edu.wcu.cs.cs363.project02.AbstractBoundedSetTest.DEFAULT_CAPACITY_2\">Constant Field Values</a></dd></dl>\n</li>\n</ul>")
	b := NewField("<a name=\"data\">\n<!--   -->\n</a>\n<ul class=\"blockListLast\">\n<li class=\"blockList\">\n<h4>data</h4>\n<pre>private&nbsp;<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedArraySet.html\" title=\"type parameter in BoundedArraySet\">T</a>[] data</pre>\n<div class=\"block\">The storage for the elements in this <code>BoundedArraySet</code>. Not\nall elements must be used.</div>\n</li>\n </ul>\n")
	//	c := NewField("<a name=\"set1\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>set1</h4>\n<pre>private&nbsp;<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedSet.html\" title=\"interface in edu.wcu.cs.cs363.project02\">BoundedSet</a>&lt;<a href=\"http://download.oracle.com/javase/6/docs/api/java/lang/String.html?is-external=true\" title=\"class or interface in java.lang\">String</a>&gt; set1</pre>\n<div class=\"block\">A <code>BoundedSet</code> with which to test.</div>\n</li>\n</ul>")
	//	d := NewField("<a name=\"head\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>head</h4>\n<pre>private&nbsp;<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedLinkedSet.ListNode.html\" title=\"class in edu.wcu.cs.cs363.project02\">BoundedLinkedSet.ListNode</a> head</pre>\n<div class=\"block\">The start of the linked list.</div>\n</li>\n</ul>")
	if a.String() != "/**\n * The default capacity for <code>set1</code>.\n */\nprotected static final DEFAULT_CAPACITY_1;" {
		t.Error("Improper output.")
	}

	if b.String() != "/**\n * The storage for the elements in this <code>BoundedArraySet</code>. Not\n * all elements must be used.\n */\nprivate T[] data;" {
		t.Error("Improper output.")
	}
}
