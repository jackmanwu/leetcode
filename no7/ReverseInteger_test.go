package no7

import "testing"

func TestReverseInteger(t *testing.T) {
	params := []int{
		123,
		423,
		901,
		40,
		9100000009,
	}

	for _, param := range params {
		println(reverse(param))
	}
}
