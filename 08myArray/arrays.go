package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in go lang")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[4] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	//It will calculate the declared memory. Not the occupied one
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string{"Potato", "Beans", "Mushroom"}
	fmt.Println("Vegy list is: ", vegList)

}
