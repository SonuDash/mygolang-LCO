package main

import "fmt"

func main() {
	fmt.Println("This is a class on arrays in golang")

	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "cheeku"
	fruitList[2] = "tangerine"
	fruitList[3] = "peaches"

	fmt.Println(fruitList[1])
	fmt.Println(fruitList)
	fmt.Println("the length of the list is:", len(fruitList))

	var vegList = [3]string{"carrot", "raddish", "potato"}
	fmt.Println("The vegie list is:", vegList)
	fmt.Println("The length of the list:", len(vegList))
}
