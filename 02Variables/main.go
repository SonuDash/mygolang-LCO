package main

import "fmt"

var LoginToken = "vab10112371" //public

func main() {

	//string
	var username string = "Ayushman"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	//boolean
	var isloggedin bool = true
	fmt.Println(isloggedin)
	fmt.Printf("Variable is of type: %T \n", isloggedin)

	//small
	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	//float
	var smallfloat float32 = 14.5983024743
	fmt.Println(smallfloat)
	fmt.Printf("Variable is of type: %T \n", smallfloat)

	//default values and aliases
	var randomNum int
	fmt.Println(randomNum)
	fmt.Printf("Variable is of type: %T \n", randomNum)

	var randomStr string
	fmt.Println(randomStr)
	fmt.Printf("Variable is of type: %T \n", randomStr)

	/*var randomChar char
	fmt.Println(randomChar)
	fmt.Printf("Variable is of type: %T \n", randomChar)
	wont work coz there's no character dataa type in golang hehe*/

	//other ways of declaring a variable
	//implicit
	var website = "ayushman.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	var activeUsers = 89
	fmt.Println(activeUsers)
	fmt.Printf("Variable is of type: %T \n", activeUsers)

	//no var required hurray!!
	//Note: this syntax is not allowed outside the method
	howManyLetters := 26
	fmt.Println(howManyLetters)

	name := "Ayushman"
	fmt.Println("What is your name?: ", name)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
