package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This is a class on if-else statement")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "WatchOut"
	} else {
		result = "Exactly 10login count"
	}
	fmt.Println(result)

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println(input)
	number := 9

	if number%2 == 0 {
		fmt.Print("%v is an even number\n", number)
	} else {
		fmt.Printf("%v is an odd number\n", number)
	}

	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is more than 10")
	}
}
