package sortalgorithm

// BubbleSort is implementation of bubble sort algorithm
func BubbleSort(arr []int) []int {
	flag := true

	for flag {
		flag = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				aux := arr[i]
				arr[i] = arr[i-1]
				arr[i-1] = aux
				flag = true
			}
		}
	}

	return arr
}
