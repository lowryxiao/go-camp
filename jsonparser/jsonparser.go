package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social struct {
		Facebook string `json:"facebook"`
		Twitter  string `json:"twitter"`
	} `json:"social"`
}

func main() {
	type users struct {
		Users []User `json:"users"`
	}
	var u users
	jsonFile := "./users.json"
	fl, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		fmt.Printf("Cannot read such file %s\n", jsonFile)
		os.Exit(1)
	}
	err = json.Unmarshal(fl, &u)
	if err != nil {
		fmt.Printf("Cannot parse such Json file %s\n", jsonFile)
	}
	fmt.Println("Successfully opened users.json")
	for _, user := range u.Users {
		fmt.Printf("User Type:%s\n", user.Type)
		fmt.Printf("User Age:%v\n", user.Age)
		fmt.Printf("User Name:%s\n", user.Name)
		fmt.Printf("Facebook Url:%s\n", user.Social.Facebook)
	}
	fmt.Printf("%v\n", u.Users)
}
