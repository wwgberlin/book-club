package main

import (
	"fmt"
	"os"
	"strconv"
)

func findMin(a []int) (min int) {
	min = a[0]
	for _, value := range a {
		if value < min {
			min = value
		}
	}
	return min
}

//func selectionSort(a []int)(sorted []int) {

//}

func main() {

	arrStr := os.Args[1:]
    arrInt := make([]int, len(arrStr))
    
    for idx, i := range arrStr {
        j, err := strconv.Atoi(i)
        if err != nil {
            panic(err)
        }
        arrInt[idx] = j
    }
	
	min := findMin(arrInt)
	fmt.Println("Min: ", min)
	//sorted := selectionSort(a)
	//fmt.Println("Sorted: ", sorted)
}
