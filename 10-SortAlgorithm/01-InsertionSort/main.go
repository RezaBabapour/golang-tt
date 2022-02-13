package main

import (
	"fmt"
	"sortAlgorithm/helper"
)

func main() {
	unOrdered := helper.GetRandomSlice(10)
	fmt.Println(unOrdered)
	ordered := insertionSort(unOrdered)
	fmt.Println(ordered)
}

func insertionSort(inputSlice []int) []int {
	for i := 0; i < len(inputSlice); i++ {
		key := inputSlice[i]
		for j := i; j > 0 && key < inputSlice[j-1]; j-- {
			inputSlice[j] = inputSlice[j-1]
			inputSlice[j-1] = key
		}
	}
	return inputSlice
}
