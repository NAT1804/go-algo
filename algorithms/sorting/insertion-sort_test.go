package sorting

import "testing"

func TestInsertionSort(t *testing.T) {
	actual := insertionSort([]int{5, 3, 1, 2, 4})

	if actual[0] != 1 || actual[1] != 2 || actual[2] != 3 || actual[3] != 4 || actual[4] != 5 {
		t.Error("Expected [1, 2, 3, 4, 5], got ", actual)
	}
}
