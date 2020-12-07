package no110

import "testing"

type Param struct {
	root *TreeNode
}

func TestIsBalanced(t *testing.T) {
	params := []Param{
		{&TreeNode{Val: 12, Left: &TreeNode{Val: 12, Left: &TreeNode{Val: 20, Left: &TreeNode{Val: 22}}}, Right: &TreeNode{Val: 23}}},
	}

	for _, param := range params {
		print(isBalanced(param.root))
	}
}
