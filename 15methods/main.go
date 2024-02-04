package main

import "fmt"

func main() {
	fmt.Println("This is a class on structures")
	//no inheritance in golang; No Super/Parent

	ayushman := User{"Ayushman", "ayushmand560@gmail.com", true, 22}
	fmt.Println(ayushman)
	fmt.Printf("Ayushman details are: %+v\n", ayushman)
	fmt.Printf("Name is %v and email is %v", ayushman.Name, ayushman.Email)
	ayushman.GetStatus()
	ayushman.NewMail()
	fmt.Printf("Name is %v and email is %v", ayushman.Name, ayushman.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	//oneAge int
}

func (u User) GetStatus() {
	fmt.Println("\nIs user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "ayushman.dash2020@vitstudent.ac.in"
	fmt.Println("The school mail id of the user is: ", u.Email)
}
