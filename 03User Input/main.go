package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating of our pizza")

	//coma err || coma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for the rating: ", input)
	fmt.Printf("Variable type of input is: %T", input)
}
