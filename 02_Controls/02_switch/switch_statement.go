package main

import "fmt"

func main() {
	var a int
	fmt.Scanf("%d", &a)

	switch a {
	case 1:
		fmt.Println("Checking for next case")
		fallthrough
	case 2:
		fmt.Print("a")
	case 3:
		fmt.Println("Checking for next case")
		fallthrough
	case 4:
		fmt.Print("b")
	case 5, 6:
		fmt.Println("c")
	default:
		fmt.Print("Value is not between 1 - 6")

	}

}
