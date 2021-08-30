package main

import "fmt"

func main() {
	var str1 string = "HI"
	var str2 = []string{"Golang", "Shivam"}
	fmt.Printf("%#v,%#v\n", str1, str2) // value + structure
	fmt.Printf("%s,%s\n", str1, str2)   //string
	fmt.Printf("%v\n", str1[0])         //Ascii
	fmt.Printf("%x\n", str1[0])         //hexadecimal in byte

}
