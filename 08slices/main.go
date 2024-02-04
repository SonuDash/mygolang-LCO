package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("This is a class on Slices")

	//Slice Syntax 1
	var fruitlist = []string{"Apple", "Peach", "Tangerine"}
	fmt.Println("The fruitlist is: \n", fruitlist)

	//Append function
	fruitlist = append(fruitlist, "Mango", "Banana")
	fmt.Println("The fruitlist is: \n", fruitlist)

	//Slicing the slice
	//1. slicing from atop
	fmt.Println("\n1st Updated fruitlist is: ", fruitlist[1:])

	//2. slicing from below
	fmt.Println("\n2nd Updated fruitlist is: ", fruitlist[:4])

	//3. slicing from the middle with upper and lower limit
	fmt.Println("\n3rd Update fruitlist is: ", fruitlist[1:4])

	//Length of the slice
	fmt.Println("\nThe length of the fruitlist is: ", len(fruitlist))
	fmt.Printf("Type of fruitlist is %T\n", fruitlist)

	//Slice Syntax 2
	highScores := make([]int, 4)
	highScores[0] = 344
	highScores[1] = 367
	highScores[2] = 677
	highScores[3] = 856
	// highScores[4] = 898

	fmt.Println("\nHigh scores are: ", highScores)
	highScores = append(highScores, 789, 900, 458)
	fmt.Println("The updated hishgscores are: ", highScores)
	sort.Ints(highScores)
	fmt.Println("The sorted highscores are: ", highScores)

	//how to remove a value from a slice based on index
	var courses = []string{"python", "ruby", "C++", "swift"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
