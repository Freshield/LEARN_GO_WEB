/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a8_create_json.py
@Time: 2021-11-14 21:55
@Last_update: 2021-11-14 21:55
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Post8 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author8 `json:"author"`
	Comments []Comment8 `json:"comments"`
}

type Author8 struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment8 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	post := Post8{
		Id: 1,
		Content: "Hello World!",
		Author: Author8{
			Id: 2,
			Name: "Sau Sheong",
		},
		Comments: []Comment8{
			{
				Id: 3,
				Content: "Have a great day!",
				Author: "Adam",
			},
			{
				Id: 4,
				Content: "How are you today?",
				Author: "Betty",
			},
		},
	}

	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = ioutil.WriteFile("post8.json", output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
		return
	}
}