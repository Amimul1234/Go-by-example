package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	//Creates a new person
	fmt.Println(person{"Bob", 20})

	//Can also give name while initializing
	fmt.Println(person{name: "Alice", age: 30})

	//Omitted field will hold value zero
	fmt.Println(person{name: "Fred"})

	//& prefix will yield a pointer to the struct
	fmt.Println(&person{"Ann", 40})

	//New struct creation in constructor function
	fmt.Println(newPerson("Jon"))

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	//Structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
