package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("this is a class on switch case")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Vallue of dice is: ", diceNumber)
	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")
	case 2:
		fmt.Println("You can move up by 2spot")
	case 3:
		fmt.Println("You can move up by 3 spot")
	case 4:
		fmt.Println("You can move up by 4 spot")
		fallthrough
	case 5:
		fmt.Println("You can move up by 5 spot")
	case 6:
		fmt.Println("Dice Value is 6 and you can open")
	default:
		fmt.Println("Roll not accepted")
	}

}
