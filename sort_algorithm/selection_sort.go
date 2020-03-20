package sortalgorithm

// SelectionSort is implementation of selection sort algorithm
func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min, minPos := arr[i], i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] < min {
				min = arr[j]
				minPos = j
			}
		}

		arr[minPos] = arr[i]
		arr[i] = min
	}

	return arr
}
