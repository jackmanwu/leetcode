package no141

import "testing"

type Param struct {
	head *ListNode
}

func TestHasCycle(T *testing.T) {
	node1 := ListNode{Val: 0}
	node2 := ListNode{Val: 2, Next: &node1}
	node1.Next = &node2
	params := []Param{
		{&ListNode{Val: 10, Next: &ListNode{Val: 2, Next: &node2}}},
		{&ListNode{Val: 3, Next: &ListNode{Val: 5}}},
	}

	for _, param := range params {
		println(hasCycle(param.head))
	}
}
