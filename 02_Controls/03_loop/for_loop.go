package main

import "fmt"

func main() {
	//infinite loop halted using if control with goto
	var count = 0
	for {
		if count == 5 {
			goto a
		}
		fmt.Println("Value of count = ", count)
		count++

	}
a:
	fmt.Print("At label")
	//while loop using for control

	count = 0
	for count < 10 {
		fmt.Println("Value of count = ", count)
		count++
	}

	//proper for loop
	for i := 0; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println("Value of i at even pos- ", i)
	}

	/*if we are iteration over any derived dataype then we get key,value pair.
	If we don't want key then we can use blank identifier (_) in place of it
	*/
	var a = []int{10, 20, 30}
	for i, v := range a {
		fmt.Printf("Value of a at pos %d is - %d\n", i, v)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
