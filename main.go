package main

import (
	"fmt"
	lc "github.com/NAT1804/go-algo/leetcode"
)

func main() {
	arr := []int{2, 7, 5, 8}
	target := 9
	a := lc.TwoSum(arr, target)
	fmt.Println(a)
}
