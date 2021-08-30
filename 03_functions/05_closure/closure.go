/*
 A closure is a special type of anonymous function that references variables declared outside of the function itself.
 It is similar to accessing global variables which are available before the declaration of the function.
*/

package main

import "fmt"

func seqGenerator() func() int {
	count := 0
	val := func() int {
		count++
		return count
	}
	return val
}

func main() {

	seq := seqGenerator()

	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())
	fmt.Println(seq())

}
