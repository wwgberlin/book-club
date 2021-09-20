package main

import (
	"fmt"
)

func partition(a []int, pivotIdx int) ([]int, []int) {
	less := make([]int, 0, len(a))
	greater := make([]int, 0, len(a))
	for i, e := range a {
		if i == pivotIdx {
			continue
		}
		if e <= a[pivotIdx] {
			less = append(less, e)
			continue
		}
		greater = append(greater, e)
	}
	return less, greater
}

// quickSortW implemented with pivot at index 0 (worst-case scenario)
func quickSortW(a []int) []int {
	if len(a) < 2 {
		return a
	}
	less, greater := partition(a, 0)
	r := append(quickSortW(less), a[0])
	r = append(r, quickSortW(greater)...)
	return r
}

// quickSortA implemented with random pivot (average-case scenario)
func quickSortA(a []int) []int {
	if len(a) < 2 {
		return a
	}
	// pivot := rand.Intn(len(a) - 1)
	pivot := len(a) / 2
	less, greater := partition(a, pivot)
	r := append(quickSortA(less), a[pivot])
	r = append(r, quickSortA(greater)...)
	return r
}

func main() {
	// fmt.Println(quickSortA([]int{}))                       // []
	// fmt.Println(quickSortA([]int{5}))                      // [5]
	fmt.Println(quickSortA([]int{1, 2, 3, 4, 5, 6, 7, 8})) // [2 3 4 5 6 7 9 9]
	fmt.Println("=======================")
	fmt.Println(quickSortW([]int{1, 2, 3, 4, 5, 6, 7, 8})) // [2 3 4 5 6 7 9 9]
}
