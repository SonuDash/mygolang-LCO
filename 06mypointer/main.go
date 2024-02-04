package main

import "fmt"

func main() {
	fmt.Println("welcome to a class on pointers")

	//pass on by memory value
	// var ptr *int

	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber //& is used for referencing
	fmt.Println("Memory value is:", ptr)
	fmt.Println("Value of actual pointer is:", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value", myNumber)

}
