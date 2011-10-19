package main

import "testing"

func TestNewField(t *testing.T) {
	a := "<a name=\"DEFAULT_CAPACITY_1\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>DEFAULT_CAPACITY_1</h4>\n<pre>protected static final&nbsp;int DEFAULT_CAPACITY_1</pre>\n<div class=\"block\">The default capacity for <code>set1</code>.</div>\n<dl><dt><span class=\"strong\">See Also:</span></dt><dd><a href=\"../../../../../constant-values.html#edu.wcu.cs.cs363.project02.AbstractBoundedSetTest.DEFAULT_CAPACITY_2\">Constant Field Values</a></dd></dl>\n</li>\n</ul>"
	b := "<a name=\"data\">\n<!--   -->\n</a>\n<ul class=\"blockListLast\">\n<li class=\"blockList\">\n<h4>data</h4>\n<pre>private&nbsp;<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedArraySet.html\" title=\"type parameter in BoundedArraySet\">T</a>[] data</pre>\n<div class=\"block\">The storage for the elements in this <code>BoundedArraySet</code>. Not\nall elements must be used.</div>\n</li>\n</ul>\n"
	c := "<a name=\"set1\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>set1</h4>\n<pre>private&nbsp;<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedSet.html\" title=\"interface in edu.wcu.cs.cs363.project02\">BoundedSet</a>&lt;<a href=\"http://download.oracle.com/javase/6/docs/api/java/lang/String.html?is-external=true\" title=\"class or interface in java.lang\">String</a>&gt; set1</pre>\n<div class=\"block\">A <code>BoundedSet</code> with which to test.</div>\n</li>\n</ul>"
	d := "<a name=\"head\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>head</h4>\n<pre>private&nbsp;<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedLinkedSet.ListNode.html\" title=\"class in edu.wcu.cs.cs363.project02\">BoundedLinkedSet.ListNode</a> head</pre>\n<div class=\"block\">The start of the linked list.</div>\n</li>\n</ul>"
	a_expected := "/**\n * The default capacity for <code>set1</code>.\n */\nprotected static final int DEFAULT_CAPACITY_1;"
	b_expected := "/**\n * The storage for the elements in this <code>BoundedArraySet</code>. Not\n * all elements must be used.\n */\nprivate T[] data;"
	c_expected := "/**\n * A <code>BoundedSet</code> with which to test.\n */\nprivate BoundedSet<String> set1;"
	d_expected := "/**\n * The start of the linked list.\n */\nprivate ListNode head;"

	a_result := NewField(a).String()
	b_result := NewField(b).String()
	c_result := NewField(c).String()
	d_result := NewField(d).String()
	if a_result != a_expected {
		t.Errorf("Test failed on input a.\nGot:\n%v\nExpected:\n%v\n", a_result, a_expected)
	} else {
		t.Log(a_result)
	}

	if b_result != b_expected {
		t.Errorf("Test failed on input b.\nGot:\n%v\nExpected:\n%v\n", b_result, b_expected)
	} else {
		t.Log(b_result)
	}

	if c_result != c_expected {
		t.Errorf("Test failed on input c.\nGot:\n%v\nExpected:\n%v\n", c_result, c_expected)
	} else {
		t.Log(c_result)
	}

	if d_result != d_expected {
		t.Errorf("Test failed on input d.\nGot:\n%v\nExpected:\n%v\n", d_result, d_expected)
	} else {
		t.Log(d_result)
	}
}
