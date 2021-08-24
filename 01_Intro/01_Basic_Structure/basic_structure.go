// Package Declaration
// The main package is the starting point to run the program.

package main

// Import Packages
//For importing multiple package at once so put each package name inside paranthese using new line delimiter
// for aliasing give the alais name before the package name

import (
	f "fmt"
	"os"
	"reflect"
	"unsafe"
)

//for importing single package use `import <package-name>`

//Global Variables
//Declaring and Assigning multiple variables in a single line
var (
	A, B int    = 1, 2
	C, D string = "Go", "Lang"
)

// Functions
//The func main() is the main function where the program execution begins.
//function opening braces should be in the same line as function declaration
func main() {
	//statement//
	f.Println(os.Getpid())
	// Variables AND Datatypes
	//Various ways to declare and assigned variables
	var var1 int8
	var1 = 100
	f.Printf("Type := %T, Value := %d\n", var1, var1)
	//typeof to find datatype and sizeof to find the size of variable in byte
	f.Println(reflect.TypeOf(var1), unsafe.Sizeof(var1))

	var var2 int16 = 200
	f.Printf("Type := %T, Value := %d\n", var2, var2)
	//default datatype for integer is int
	var var3 = 123
	f.Printf("Type := %T, Value := %d\n", var3, var3)

	//default datatype for float is float64
	var var4 = 123.4
	f.Printf("Type := %T, Value := %f\n", var4, var4)

	var var5 float32 = 123.4
	f.Printf("Type := %T, Value := %.2f\n", var5, var5) //Precision control

	//	Short hand variable declaration and assignment
	// 	Short hand can't be used in global declaration
	var6 := 125
	f.Printf("Type := %T, Value := %v\n", var6, var6) //Automatic identify datatype Not Recommended!!!
	f.Printf("Type := %T, Value := %s\n", var6, var6) //Type := int, Value := %!s(int=125)

	//string declaration
	var var7 string = "Go Lang"
	f.Printf("Type := %T, Value := %s\n", var7, var7)

	f.Printf("Accessing global variables A:= %d  and B:= %d\n", A, B)
	f.Printf("Accessing global variables C:= %s  and D:= %s", C, D)

}

/*
#	Notes:-
The Go compiler automatically inserts semicolons at the end of statements.
However, if multiple statements are written on one line (which is not encouraged for readability reasons),
they must be separated by ;

#	Saving file -
	Filenames consist of lowercase-letters (that may separated by underscores _)


#	For running  a go file use syntax-
	go run <file-name.go>

# 	For creating exec and binary file-
	go install <file-name.go> //To make exe in workspace/bin folder
	go build <file-name.go> //To make exe where your go file exist
*/
