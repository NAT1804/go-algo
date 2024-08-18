package sorting

/*
Insertion sort
Idea: Initially, only the first item is considered sorted.
Then, for each item in the sequence, we insert that item into the sorted list by swapping it with the item
before it until the item is smaller than the current item.

*/

func insertionSort(data []int) []int {
	for i := 1; i < len(data); i++ {
		current := i

		// Loop from current to 0
		for current > 0 && data[current-1] > data[current] {
			data[current], data[current-1] = data[current-1], data[current]
			current--
		}
	}
	return data
}
