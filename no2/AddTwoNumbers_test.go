package no2

import (
	"fmt"
	"testing"
)

type Param struct {
	l1 *ListNode
	l2 *ListNode
}

func TestAddTwoNumbers(t *testing.T) {
	params := []Param{
		/**
		852+901=1753
		*/
		{&ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 8}}}, &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 9}}}},
	}
	for _, param := range params {
		var result = addRTwoNumbers(param.l1, param.l2)
		var current = result
		for current != nil {
			fmt.Print(current.Val, "->")
			current = current.Next
		}
	}
}
