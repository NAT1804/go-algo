package leetcode

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		// Skip all duplicates from left
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target, left, right := -nums[i], i+1, n-1
		for left < right {
			sum := nums[left] + nums[right]
			if sum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				left++
				right--
				// Skip all duplicate from left
				for left < right && nums[left] == nums[left-1] {
					left++
				}
				// Skip all duplicates from right
				for left < right && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}
	return result
}
