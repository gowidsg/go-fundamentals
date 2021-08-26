/*
Go language provides a special feature known as an anonymous function.
An anonymous function is a function which doesn’t contain any name.
It is useful when you want to create an inline function.
In Go language, an anonymous function can form a closure. An anonymous function is also known as function literal.
*/

package main

import (
	"fmt"
	"math"
)

func avg(v int) {
	fmt.Println(v / 2)
}

func Sine() func(x float64) float64 {
	myanonymous := func(x float64) float64 {
		return math.Sin(x)
	}
	return myanonymous
}

//anonymous inline function
func main() {
	func() {
		fmt.Println("Hi! Golang")
	}()

	/*
	   In Go language, you are allowed to assign an anonymous function to a variable.
	   When you assign a function to a variable, then the type of the variable is of function type
	   and you can call that variable like a function call as shown in the below example.
	*/
	val := func(a, b int) int {
		return a + b
	}

	fmt.Println(val(1, 2))

	//You can also pass an anonymous function as an argument into other function
	avg(val(2, 2))

	//You can also return an anonymous function from another function.

	sin_val := Sine()

	fmt.Println(sin_val(180))
}
