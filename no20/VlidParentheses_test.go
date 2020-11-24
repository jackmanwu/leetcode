package no20

import "testing"

func TestValidParentheses(t *testing.T) {
	params := []string{
		"{}",
		"",
		"(",
		")(",
	}

	for _, param := range params {
		println(isValid(param))
	}
}
