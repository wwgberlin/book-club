package main

import "fmt"

func minIndex(o []int) int {
	min := o[0]
	index := 0
	for i := range o {
		if o[i] < min {
			min = o[i]
			index = i
		}
	}
	return index
}

// removeIndex creates new slice from slice `s` after removing the element at index `i`
func removeIndex(s []int, i int) []int {
	r := make([]int, len(s[:i]))
	copy(r, r[:i])
	return append(r, r[i+1:]...)
}

// selectionSort is a recursive implementation of the selection sort algorithm.
// The implementation is called tail-recursive because there is no more instruction
// need to be executed after calling itself.
// Note: Yet worth noting Go does not optimize tail-recursion.
func selectionSort(original, result []int) []int {
	if len(original) == 0 {
		return result
	}
	i := minIndex(original)              // get the index originalf the smallest element
	result = append(result, original[i]) // append the new min to the accumulator (result array)
	s := removeIndex(original, i)        // removing the the min from the old array
	return selectionSort(s, result)
}

func main() {
	original := []int{2, 4, 1, 3, 7, 5, 9}
	acc := make([]int, 0, len(original))
	sorted := selectionSort(original, acc)
	fmt.Println("sorted  ", sorted)
	fmt.Println("original", original)
}
