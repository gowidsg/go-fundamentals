/*
The function that called with the varying number of arguments is known as variadic function.
Or in other words, a user is allowed to pass zero or more arguments in the variadic function.
fmt.Printf is the example of the variadic function,
it required one fixed argument at the starting after that it can accept any number of arguments
*
1. 	In the declaration of the variadic function, the type of the last parameter is preceded by an ellipsis, i.e, (…).
	It indicates that the function can be called at any number of parameters of this type.

2. 	Inside the function …type behaves like a slice. For example, suppose we have a function signature, i.e, add( b…int)int, now the a parameter of type[]int.

3.	You can also pass an existing slice in a variadic function. To do this, we pass a slice of the complete array to the function as shown in Example 2 below.

4.	When you do not pass any argument in the variadic function, then the slice inside the function is nil.

5.	You can also pass multiple slice in the variadic function.

6.	You can not use variadic parameter as a return value, but you can return it as a slice.
*/

package main

import "fmt"

func sum_of_array(s string, elem ...int) (string, int) { //argument
	var sum = 0
	for _, v := range elem {
		sum += v
	}
	return s, sum

}

func main() {

	fmt.Println(sum_of_array("HI", 1, 2, 3)) //parameter

	fmt.Println(sum_of_array("Slice", []int{1, 6, 7}...))

}
