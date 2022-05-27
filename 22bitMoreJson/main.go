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

func DecodeJson() {
	jsonFromWeb := []byte(`
		{
                "coursename": "ReactJs BootCamp",
                "price": 299,
                "wesites": "LearnCode.in",
                "tags": ["web-dev", "js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonFromWeb)
	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json was not valid")
	}

	//Don't want to crate model classes each time. Rather give me key value pair
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
}

func main() {
	fmt.Println("Welcome to JSON video")
	//EncodeJson()
	DecodeJson()
}
