package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Int to Float

	var var1 = 10
	fmt.Printf("Type:- %T , Value:- %v\n", var1, var1)
	fmt.Printf("Type:- %T , Value:- %v\n", float64(var1), float64(var1))

	// Float to Int

	var var2 = 10.3433
	fmt.Printf("Type:- %T , Value:- %v\n", var2, var2)
	fmt.Printf("Type:- %T , Value:- %v\n", int(var2), int(var2))

	var var3 = "101.1"
	f, _ := strconv.ParseFloat(var3, 64)
	fmt.Printf("Type:- %T , Value:- %v\n", f, f)

	// i, _ := strconv.ParseInt(var3, 2, 32)//can't convert float string into integer using parseInt
	// fmt.Printf("Type:- %T , Value:- %v\n", i, i)

	var var4 = "501"
	i, _ := strconv.Atoi(var4)
	fmt.Printf("Type:- %T , Value:- %v\n", i, i)

	str := strconv.Itoa(var1)
	fmt.Printf("Conversion of integer var1 (%v) into (%T): %v\n", var1, str, str)

	str1 := strconv.FormatFloat(var2, 'f', 2, 64)
	fmt.Printf("Conversion of float var2 (%v) into (%T): %v\n", var2, str1, str1)

	fmt.Printf("%T, %v", fmt.Sprintf("%.2f", var2), fmt.Sprintf("%.3f", var2))

	// b := fmt.Sprintf("%.2f", a)
	// fmt.Printf("%T, %v", b, b)
}
