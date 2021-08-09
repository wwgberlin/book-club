package main

import (
	"math"
)

func max(a []int) int {
	if len(a) == 0 {
		// using NaN (Not a number) just for simplicity
		// so I don't have to return an error
		return int(math.NaN())
	}

	maxInt := max(a[1:])
	if a[0] > maxInt {
		return a[0]
	}
	return maxInt
}

// func main() {
// 	fmt.Println("max([]) =", max([]int{}))                     // = NaN
// 	fmt.Println("max([5]) =", max([]int{5}))                   // = 5
// 	fmt.Println("max([5, 3, 9, 1]) =", max([]int{5, 3, 9, 1})) // = 9
// }
