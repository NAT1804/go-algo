package sorting

/*
We use a pointer to point at the first item of the list
For each cycle, we compare it to the next element in the list and swap them if
the current item is greater, then move the pointer by one unit it reaches the end of list
We repeat this process
*/

func bubbleSort(list []int) []int {
	n := len(list)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}

	return list
}
