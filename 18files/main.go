package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	fmt.Println("Welcome to files in go lang")

	content := "This need to go in file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err)
	} else {

		length, err := io.WriteString(file, content)
		checkNilError(err)

		fmt.Println("Length is: ", length)

		//This is the recommended way. So, that it does not create any blocker
		defer func(file *os.File) {
			err := file.Close()
			checkNilError(err)
		}(file)
	}

	readFile("./mylcogofile.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("Text data inside file is:- \n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
