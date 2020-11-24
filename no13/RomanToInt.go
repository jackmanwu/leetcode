package no13

var roman = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	if s == "" {
		return 0
	}
	last, total := 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		var num = roman[s[i:i+1]]
		if num < last {
			total -= num
		} else {
			total += num
		}
		last = num
	}
	return total
}
