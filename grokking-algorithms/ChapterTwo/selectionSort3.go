package main

import (
	"fmt"
)

func selectionSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	return arr
}

func main() {
	arrInt := []int{10, 6, 2, 7, 8}

	sorted := selectionSort(arrInt)
	fmt.Println("Sorted: ", sorted)
}