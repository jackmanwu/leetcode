package no1

import (
	"fmt"
	"testing"
)

type Param struct {
	nums   []int
	target int
}

func TestTwoSum(t *testing.T) {
	params := []Param{
		{[]int{1, 2}, 3},
		{[]int{0, 4, 6}, 10},
		{[]int{8, 1, 22, 45, 78}, 79},
	}

	for _, param := range params {
		var arr = TwoSum(param.nums, param.target)
		for i := 0; i < len(arr); i++ {
			fmt.Print(arr[i], " ")
		}
		fmt.Println()
	}
}
