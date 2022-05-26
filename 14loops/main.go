package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in go lang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	fmt.Println()

	for i := range days {
		fmt.Println(days[i])
	}

	fmt.Println()

	for index, days := range days {
		fmt.Printf("Index is %v and value is %v\n", index, days)
	}

	rougeValue := 1

	for rougeValue < 10 {

		if rougeValue == 2 {
			goto lco
		}

		if rougeValue == 5 {
			rougeValue++
			continue
		}

		fmt.Println("Value is : ", rougeValue)
		rougeValue++
	}

lco:
	fmt.Println("Jumping at Learn Code Online")
}
