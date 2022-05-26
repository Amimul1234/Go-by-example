package main

import "fmt"

func main() {
	fmt.Println("If else in go lang")

	var result string
	logInCount := 23

	if logInCount < 10 {
		result = "Regualr user"
	} else if logInCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is evenb")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is not less than 10")
	}

}
