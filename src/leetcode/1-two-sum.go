package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for currIdx, currNum := range nums {
		if targetIdx, isPresent := m[target-currNum]; isPresent {
			return []int{targetIdx, currIdx}
		}
		m[currNum] = currIdx
	}
	return []int{}
}
