package leetcode

import "testing"

func TestTwoSum(t *testing.T) {
	actual := TwoSum([]int{2, 7, 11, 15}, 9)

	if actual[0] != 0 || actual[1] != 1 {
		t.Error("Expect value of actual is [0, 1]")
	}
}
