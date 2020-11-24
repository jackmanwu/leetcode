package no14

import "testing"

type Param struct {
	str []string
}

func TestLongestCommonPrefix(t *testing.T) {
	params := []Param{
		{str: []string{"acccccaa", "xxxx"}},
	}

	for _, param := range params {
		println(longestCommonPrefix(param.str))
	}
}
