package no21

import "testing"

type Param struct {
	l1 *ListNode
	l2 *ListNode
}

func TestMergeTwoLists(t *testing.T) {
	params := []Param{
		{&ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 7}}},
			&ListNode{Val: 10, Next: &ListNode{Val: 21, Next: &ListNode{Val: 3}}},
		},
	}

	for _, param := range params {
		result := mergeToLists(param.l1, param.l2)
		for result != nil {
			print(result.Val, "->")
			result = result.Next
		}
		println()
	}
}
