package main

import "testing"

func TestNewMethod(t *testing.T) {
	a := "<a name=\"find(java.lang.Object)\">\n<!--   -->\n</a><a name=\"find(T)\">\n<!--   -->\n</a>\n<ul class=\"blockList\">\n<li class=\"blockList\">\n<h4>find</h4>\n<pre>private&nbsp;<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedLinkedSet.ListNode.html\" title=\"class in edu.wcu.cs.cs363.project02\">BoundedLinkedSet.ListNode</a>&nbsp;find(<a href=\"../../../../../edu/wcu/cs/cs363/project02/BoundedLinkedSet.html\" title=\"type parameter in BoundedLinkedSet\">T</a>&nbsp;element)</pre>\n<div class=\"block\">Finds the node containing the specified element in the linked list,\nor <code>null</code> if the element is not found.</div>\n<dl><dt><span class=\"strong\">Parameters:</span></dt><dd><code>element</code> - the element to search for.</dd>\n<dt><span class=\"strong\">Returns:</span></dt><dd>a reference to the node in the linked list containing element\n	 or <code>null</code> of no such node exists.</dd></dl>\n</li>\n</ul>"
	//b := ""
	a_expected := "/**\n * Finds the node containing the specified element in the linked list,\n * or <code>null</code> if the element is not found.\n *\n * @param element the element to search for.\n *\n * @return a reference to the node in the linked list containing element\n * or <code>null</code> of no such node exists.\n */\nprivate ListNode find(T element) {\n\n}"
	//b_expected := ""
	a_result := NewMethod(a).String()
	//b_result := NewMethod(b).String()

	if a_result != a_expected {
		t.Errorf("Test failed on input a.\nGot:\n%v\nExpected:\n%v\n", a_result, a_expected)
	}
	/*
		if b_result != b_expected {
			t.Errorf("Test failed on input b.\nGot:\n%v\nExpected:\n%v\n", b_result, b_expected)
		}*/
}
