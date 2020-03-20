package main

import (
	"fmt"

	sortalgorithm "github.com/jorgeAM/algorithm/sort_algorithm"
)

func main() {
	arr := []int{90, -80, 70, -60, 50, -40, 30, -20, 10, 0}

	arr = sortalgorithm.BubbleSort(arr)

	fmt.Println(arr)
}
