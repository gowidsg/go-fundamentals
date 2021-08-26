package main

import "fmt"

func Greet(name string) string {
	return "Hi! I am " + name
}

func Standard_Calc(a, b, c int) (int, int) {
	return a + b + c, (a + b + c) / 3
}
func main() {
	fmt.Println(Greet("Golang"))
	sum, avg := Standard_Calc(1, 2, 3)
	fmt.Println(sum, avg)
}
