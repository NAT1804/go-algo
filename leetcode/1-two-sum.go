package leetcode

/**
Description: https://leetcode.com/problems/two-sum/

Solution:
Using map with
key: num in nums
value: index of num in nums
*/

func TwoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for currIdx, currNum := range nums {
		if targetIdx, isPresent := m[target-currNum]; isPresent {
			return []int{targetIdx, currIdx}
		}
		m[currNum] = currIdx
	}
	return []int{}
}
