package leetcode

import (
	"math"
)

/**
Description: https://leetcode.com/problems/reverse-integer/
*/

func Reverse(x int) int {
	result := 0
	isNegative := x < 0
	if isNegative {
		x = -x
	}
	for x > 0 {
		result = result*10 + x%10
		x /= 10
		if result > math.MaxInt32 || result < math.MinInt32 {
			return 0
		}
	}
	if isNegative {
		result = -result
	}
	return result
}
