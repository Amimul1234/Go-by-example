package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Structs in go lang")
	//no inheritance in go lang, No super or parent there

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)

}
