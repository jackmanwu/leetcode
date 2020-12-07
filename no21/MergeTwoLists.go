package no21

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeToLists(l1 *ListNode, l2 *ListNode) *ListNode {
	start := &ListNode{Val: -1}
	prev := start
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}
	if l1 == nil {
		prev.Next = l2
	} else {
		prev.Next = l1
	}
	return start.Next
}
