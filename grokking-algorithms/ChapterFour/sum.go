package main

/*Tip from the book:
When you’re writing a recursive function involving an array, the base case is
often an empty array or an array with one element. Try it first.
*/

func sum(a []int) int {
	if len(a) == 0 {
		return 0
	}
	return a[0] + sum(a[1:])
}

// func main() {
// 	fmt.Println("sum([]) =", sum([]int{}))                     // = 0
// 	fmt.Println("sum([5]) =", sum([]int{5}))                   // = 5
// 	fmt.Println("sum([5, 3, 9, 1]) =", sum([]int{5, 3, 9, 1})) // = 18
// }
