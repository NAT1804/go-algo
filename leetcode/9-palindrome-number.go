package leetcode

func IsPalindrome(x int) bool {
	reverseNum := 0
	num := x
	if x < 0 {
		return false
	}
	for x > 0 {
		reverseNum = reverseNum*10 + x%10
		x = x / 10
	}
	return reverseNum == num
}
