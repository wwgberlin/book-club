package main

import (
	"fmt"
)

func findMin(a []int, indices []bool) int {
	min := a[0]
	minIndex := 0
	for i, v := range a {
		if !indices[i] && v < min {
			minIndex = i
			min = v
		}
	}
	return minIndex
}

func selectionSort(a []int)[]int {
	newArr := make([]int, len(a))
	indices := make([]bool, len(a))
	for i, _ := range a {
		minIndex := findMin(a, indices)
		newArr[i] = a[minIndex]
		indices[minIndex] = true
	}
	return newArr
}

func main() {
        arrInt := []int{10, 6, 2, 7, 8}

	sorted := selectionSort(arrInt)
	fmt.Println("Sorted: ", sorted)	
}
