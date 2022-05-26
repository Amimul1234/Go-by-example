package main

import "fmt"

//Regular things -> functions
//In structs -> methods

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

// NewMail is creating a copy of the user.
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

func main() {
	hitesh := User{"Hitesh", "hitest@go.dev", true, 15}

	fmt.Println(hitesh)
	fmt.Printf("hitesh details are : %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)

	hitesh.GetStatus()
	hitesh.NewMail()

	fmt.Println(hitesh)
}
