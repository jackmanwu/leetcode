package no3

import "testing"

func TestLongestSubstring(t *testing.T) {
	params := []string{
		"abcabcabc",
	}
	for _, param := range params {
		println(lengthOfLongestSubstring(param))
		println(lengthOfLongestSubstring_(param))
	}
}
