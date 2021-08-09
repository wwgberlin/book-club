package main

func count(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return 1 + count(a[1:])
}

// func main() {
// 	fmt.Println("count([]) =", count([]int{}))                     // = 0
// 	fmt.Println("count([5]) =", count([]int{5}))                   // = 1
// 	fmt.Println("count([5, 3, 9, 1]) =", count([]int{5, 3, 9, 1})) // = 4
// }
