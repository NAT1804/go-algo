package sorting

func bubbleSort(list []int) []int {
	n := len(list)

	for i := 0; i < n; i++ {
		swapped := false
		for j := i + 1; j < n; j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
				swapped = true
			}
		}
		if !swapped {
			return list
		}
	}
	return list
}
