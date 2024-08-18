package sorting

/**
The idea for this sorting algorithm is that during each cycle,
we find the smallest item from the unsorted pile
and add it to the sorted pile
*/

func selectionSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	return arr
}
