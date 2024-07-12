package leetcode

func RomanToInt(s string) int {
	var r, pr, cur int

	for i := len(s) - 1; i >= 0; i-- {
		cur = convertCharToInt(s[i])
		if cur < pr {
			r -= cur
		} else {
			r += cur
		}
		pr = cur
	}
	return r
}

func convertCharToInt(c byte) int {
	switch c {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}
