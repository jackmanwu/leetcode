package no27

import "testing"

type Param struct {
	nums []int
	val  int
}

func TestRemoveElement(t *testing.T) {
	params := []Param{
		{nums: []int{4, 4, 1, 2, 3, 4, 5}, val: 4},
	}

	for _, param := range params {
		println(removeElement(param.nums, param.val))
	}
}
