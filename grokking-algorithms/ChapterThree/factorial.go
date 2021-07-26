package main

import "fmt"

func fact(x int) int {
	if x == 0 {
		return 1
	}

	return x * fact(x-1)
}

func fact2(x int) int {
	return factTail(x, 1)
}

// factTail is the tail-recusrion implementation for fact
// the resulat accumulator `acc` should be = 1
func factTail(x, acc int) int {
	if x == 0 {
		return acc
	}
	return factTail(x-1, acc*x)
}

func main() {
	fmt.Println("fact ", fact(5))
	fmt.Println("fact2", fact2(5))
}
