package main

import (
	"fmt"
	"log"
	"sortAlgorithm/helper"
)

func main() {
	unOrdered := helper.GetRandomSlice(10)
	fmt.Println(unOrdered)
	ordered := insertionSort(unOrdered)
	fmt.Println(ordered)
}

func insertionSort(inputSlice []int) []int {
	num:=0
	for i := 0; i < len(inputSlice); i++ {
		key := inputSlice[i]
		num++
		for j := i; j > 0 && key < inputSlice[j-1]; j-- {
			inputSlice[j] = inputSlice[j-1]
			inputSlice[j-1] = key
			num=num+2
		}
	}
	log.Println("num: ", num)
	return inputSlice
}
