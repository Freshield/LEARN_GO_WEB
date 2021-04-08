/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a9_encoder_json.py
@Time: 2021-11-14 22:02
@Last_update: 2021-11-14 22:02
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
	"os"
)

type Post9 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author9 `json:"author"`
	Comments []Comment9 `json:"comments"`
}

type Author9 struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment9 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	post := Post9{
		Id: 1,
		Content: "Hello World!",
		Author: Author9{
			Id: 2,
			Name: "Sau Sheong",
		},
		Comments: []Comment9{
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
	jsonFile, err := os.Create("post9.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer jsonFile.Close()
	encoder := json.NewEncoder(jsonFile)
	encoder.SetIndent("", "\t\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
}