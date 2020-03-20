package sortalgorithm

// MergeSort is implementation of merge sort algorithm
func MergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	size := len(arr)
	half := size / 2
	left := arr[0:half]
	right := arr[half:size]
	left = MergeSort(left)
	right = MergeSort(right)
	return merge(left, right)
}

func merge(left, right []int) []int {
	lastLeft := left[len(left)-1]
	if lastLeft <= right[0] {
		return append(left, right...)
	}

	result := make([]int, 0)
	var min int

	for len(left) > 0 && len(right) > 0 {
		if left[0] <= right[0] {
			min = left[0]
			left = left[1:]
		} else {
			min = right[0]
			right = right[1:]
		}

		result = append(result, min)
	}

	if len(left) > 0 {
		result = append(result, left...)
	}

	if len(right) > 0 {
		result = append(result, right...)
	}

	return result
}
