package main

import (
	"fmt"
)

func findMin(a []int) int {
	min := a[0]
	minIndex := 0
	for i, v := range a {
		if v < min {
			minIndex = i
			min = v
		}
	}
	return minIndex
}

// func selectionSort(a []int)[]int {
// 	newArr := make([]int, 0, len(a))
// 	for _= range a {
// 		minIndex := findMin(a)
// 		newArr = append(newArr, a[minIndex])
// 		a = append(a[0:minIndex], a[minIndex+1:]...)
// 		//fmt.Println(a)
// 		//fmt.Println(newArr)
// 	}
// 	return newArr
// }

func selectionSort(a []int)[]int {
	newArr := make([]int, len(a))
	for i, _ := range a {
		minIndex := findMin(a)
		newArr[i] = a[minIndex]
		a = append(a[0:minIndex], a[minIndex+1:]...)
	}
	return newArr
}

func main() {

        arrInt := []int{10, 6, 2, 7, 8}

	min := findMin(arrInt)
	fmt.Println("MinIndex: ", min)
	
	sorted := selectionSort(arrInt)
	fmt.Println("Sorted: ", sorted)
	
}
