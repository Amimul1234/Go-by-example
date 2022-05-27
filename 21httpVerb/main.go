package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func PerformGetRequest() {
	const myUrl = "http://localhost:8000/get"

	response, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.Status)
	fmt.Println("Content length is: ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println(byteCount)
	fmt.Println(responseString.String())
}

func performPostJsonRequest() {
	const myUrl string = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"courseName" : "Let's go with the go lang",
			"price" : 0,
			"platform" : "abcd.com"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))
}

func performPostFromRequest() {
	const myUrl string = "http://localhost:8000/postform"

	data := url.Values{}
	data.Add("firstName", "Amimul")
	data.Add("lastName", "Ehsan")
	data.Add("email", "1805056@ugrad.cse.buet.ac.bd")

	response, _ := http.PostForm(myUrl, data)
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

	defer response.Body.Close()
}

func main() {
	//PerformGetRequest()
	//performPostJsonRequest()
	performPostFromRequest()
}
