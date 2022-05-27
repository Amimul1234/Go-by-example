package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"wesites"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs BootCamp", 299, "LearnCode.in", "abc123", []string{"web-dev", "js"}},
		{"Mern BootCamp", 199, "LearnCode.in", "bcd123", []string{"web-dev", "js"}},
		{"Angular BootCamp", 299, "LearnCode.in", "asdfsdaf", []string{"web-dev", "js"}},
	}

	finalJson, _ := json.MarshalIndent(lcoCourses, "", "\t")

	fmt.Printf("%s\n", finalJson)
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
}
