package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ipter *int) {
	*ipter = 0
}

func main() {
	i := 1
	fmt.Println("Initial: ", 1)

	zeroval(i)
	fmt.Println("ZeroVal: ", i)

	zeroptr(&i)
	fmt.Println("zero pointer: ", i)

	fmt.Println("pointer: ", &i)
}
