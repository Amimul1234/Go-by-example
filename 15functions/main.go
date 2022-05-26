package main

import "fmt"

//Main acts the entry point for the go lang
func main() {

	fmt.Println("Welcome to functions in go lang")
	greeter()

	//function should not contain another function definition
	//func greeter2(){
	//	fmt.Println("Another method")
	//}

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	result = proAdder(1, 3, 4, 6)
	fmt.Println(result)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}

func greeter() {
	fmt.Println("Greeting from go lang")
}
