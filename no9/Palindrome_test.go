package no9

import "testing"

func TestPalindrome(t *testing.T) {
	params := []int{
		12,
		-121,
		0,
		100,
		22,
		13431,
		1313,
		1331,
		1<<31 - 1,
	}
	for _, param := range params {
		println(isPalindrome(param))
	}
}
