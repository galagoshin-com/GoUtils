package strings

import (
	"math"
	"unicode/utf8"
)

func Distance(a, b string) int {
	if len(a) == 0 {
		return utf8.RuneCountInString(b)
	}
	if len(b) == 0 {
		return utf8.RuneCountInString(a)
	}
	if a == b {
		return 0
	}
	s1 := []rune(a)
	s2 := []rune(b)
	if len(s1) > len(s2) {
		s1, s2 = s2, s1
	}
	lenS1 := len(s1)
	lenS2 := len(s2)
	x := make([]uint16, lenS1+1)
	for i := 1; i < len(x); i++ {
		x[i] = uint16(i)
	}
	_ = x[lenS1]
	for i := 1; i <= lenS2; i++ {
		prev := uint16(i)
		for j := 1; j <= lenS1; j++ {
			current := x[j-1]
			if s2[i-1] != s1[j-1] {
				current = uint16(math.Min(math.Min(float64(x[j-1]+1), float64(prev+1)), float64(x[j]+1)))

			}
			x[j-1] = prev
			prev = current
		}
		x[lenS1] = prev
	}
	return int(x[lenS1])
}
