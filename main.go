package main

import (
	"fmt"
	lc "github.com/NAT1804/go-algo/leetcode"
)

func main() {
	// Problem #1
	//arr := []int{2, 7, 5, 8}
	//target := 9
	//a := lc.TwoSum(arr, target)
	//fmt.Println(a)

	// Problem #7
	//b := lc.Reverse(-123)
	//c := lc.Reverse(123)
	//fmt.Println(b, c)

	// Problem #9
	//fmt.Println(lc.IsPalindrome(-121))

	// Problem #15
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(lc.ThreeSum(nums))
}
