package main

import "fmt"

func main() {
	greeter()
	fmt.Println("This is a class on functions")
	result := adder(3, 5)
	fmt.Println("The result is: ", result)

	sum := proadder(2, 3, 4, 5, 4, 7, 10, 34)
	fmt.Println("The sum of multiple numbers is: ", sum)
}

func adder(valOne int, valTwo int) int {
	result := valOne + valTwo
	return result
}

func proadder(values ...int) int {
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

func greeter() {
	fmt.Println("Hello from golang")
}
