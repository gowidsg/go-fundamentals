//if condition checks for boolean only
package main

import "fmt"

func main() {
	if true {
		println("True")
	} else {
		println("False")
	}

	var a, b, c = 15, 16, 7
	if c > a {
		if c > b {
			fmt.Printf("%v is greater", c)
		} else {
			fmt.Printf("%v is greater", b)
		}
	} else if a > b {
		fmt.Printf("%v is greater", a)
	} else {
		fmt.Printf("%v is greater", b)
	}
}
