package main

import (
	"fmt"
)

func binarySearch(list []interface{}, element interface{}) int {
	// what if list has different types than the element
	// what if the list contains different types
	// switch on type (of element?) and call appropiate function
	// if if element has type of int then call binarySearchInteger

	// Or use generics?

}

func binarySearchInteger(input []int, element int) int {
	high := len(input) - 1
	low := 0
	for low <= high {
		pivot := (high + low) / 2
		switch {
		case input[pivot] == element:
			return pivot
		case input[pivot] > element:
			high = pivot - 1
		case input[pivot] < element:
			low = pivot + 1
		}
	}
	return -1
}

func binarySearchString(input []string, element string) int {
	high := len(input) - 1
	low := 0
	for low <= high {
		pivot := (high + low) / 2
		switch {
		case input[pivot] == element:
			return pivot
		case input[pivot] > element:
			high = pivot - 1
		case input[pivot] < element:
			low = pivot + 1
		}
	}
	return -1
}

func main() {
	fmt.Println("Hello, Binary Search")
}