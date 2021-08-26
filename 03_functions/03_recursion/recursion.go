/*
The Three Laws of Recursion
A recursive algorithm must have a base case .
A recursive algorithm must change its state and move toward the base case .
A recursive algorithm must call itself, recursively.
*/

package main

import "fmt"

func sum_recursion(n int) int {
	if n < 0 {
		fmt.Println("invalid number")
		return -1
	}
	if n == 1 {
		return 1
	}
	return n + sum_recursion(n-1)
}

func main() {
	fmt.Println(sum_recursion(-10))
}
