package no13

import "testing"

func TestRomanToInt(t *testing.T) {
	params := []string{
		"II",
		"VI",
		"IV",
		"XII",
	}
	for _, param := range params {
		println(romanToInt(param))
	}
}
