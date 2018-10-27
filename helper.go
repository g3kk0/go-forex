package goforex

import (
	"fmt"
	"strings"
)

func Ftoi(f float64) (i int64, dp int) {
	fStr := fmt.Sprintf("%v", f)
	s := strings.Split(fStr, ".")

	// get decimal precision and calculate multiple.
	multiple := 1
	for i := 0; i < len(s[1]); i++ {
		dp++
		multiple = multiple * 10
	}

	i = int64(f * float64(multiple))

	return i, dp
}

func Itof(i int64, dp int) (f float64) {
	divide := 1
	for i := 0; i < dp; i++ {
		divide = divide * 10
	}

	return float64(i) / float64(divide)
}
