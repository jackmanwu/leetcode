package no26

import "testing"

type Param struct {
	nums []int
}

func TestRemoveDuplicates(t *testing.T) {
	params := []Param{
		{[]int{1, 1, 1, 2, 2}},
	}

	for _, param := range params {
		println(removeDuplicates(param.nums))
	}
}
