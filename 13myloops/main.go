package main

import (
	"fmt"
)

func main() {
	fmt.Println("This is a class about functions")
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, days := range days {
	// 	fmt.Printf("Index is %v and the value is %v\n", index, days)
	// }

	// for _, days := range days {
	// 	fmt.Printf("Index is nahh and the value is %v\n", days)
	// }

	rogueValue := 1
	for rogueValue <= 10 {
		// if rogueValue == 5 {
		// 	break
		// }

		if rogueValue == 4 {
			goto letsplay
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}
		fmt.Println(rogueValue)
		rogueValue++
	}

letsplay:
	fmt.Println("Let's play it's 4 already")

}
