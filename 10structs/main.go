package main

import "fmt"

func main() {
	fmt.Println("This is a class on structures")
	//no inheritance in golang; No Super/Parent

	ayushman := User{"Ayushman", "ayushmand560@gmail.com", true, 22}
	fmt.Println(ayushman)
	fmt.Printf("Ayushman details are: %+v\n", ayushman)
	fmt.Printf("Name is %v and email is %v", ayushman.Name, ayushman.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
