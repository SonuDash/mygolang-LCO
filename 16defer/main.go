package main

import "fmt"

func main() {
	defer fmt.Println("Ok bubye")
	defer fmt.Println("this is line before bubye")
	defer fmt.Println("this is 2 lines before bubye")
	fmt.Println("Welcome to a tutorial on defer")
	endOfLine()
}

//defer is kinda......like LIFO (Last In First Out)

func endOfLine() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
