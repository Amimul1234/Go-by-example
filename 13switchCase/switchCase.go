package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Switch and case in go lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move two spots")
	case 3:
		fmt.Println("You can move to 3 spots")
	case 4:
		fmt.Println("You can move to 4 spots")
	case 5:
		fmt.Println("You can move to 5 spots")
	case 6:
		fmt.Println("You can  move to 6 spot and roll the dice again")
	default:
		fmt.Println("What was that!")
	}

}
