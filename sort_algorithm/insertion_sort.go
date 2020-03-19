package sortalgorithm

// InsertionSort is implementation of insertion sort algorithm
func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i

		for j > 0 && arr[j-1] > value {
			arr[j] = arr[j-1]
			arr[j-1] = value
			j--
		}
	}

	return arr
}
