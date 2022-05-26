package main

import "fmt"

func main() {
	var username string = "Amimul Ehsan"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallFloat float32 = 255.455554545345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default vales and some aliases
	var anotherVatiable int
	fmt.Printf("Variable is of type : %T \n", anotherVatiable)
	fmt.Println(anotherVatiable)

	var another string
	fmt.Printf("Variable is of type : %T \n", another)
	fmt.Println(another)
}
