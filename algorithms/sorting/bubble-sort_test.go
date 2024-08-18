package sorting

import "testing"

func TestBubbleSort(t *testing.T) {
	actual := bubbleSort([]int{4, 3, 2, 1})

	if actual[0] != 1 || actual[1] != 2 || actual[2] != 3 || actual[3] != 4 {
		t.Error("Expected [1,2,3,4], got ", actual)
	}
}
