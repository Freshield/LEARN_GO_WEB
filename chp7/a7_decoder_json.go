/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_decoder_json.py
@Time: 2021-11-14 21:37
@Last_update: 2021-11-14 21:37
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
	"io"
	"os"
)

type Post7 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author7 `json:"author"`
	Comments []Comment7 `json:"comments"`
}

type Author7 struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment7 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post6.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	decoder := json.NewDecoder(jsonFile)
	for {
		var post Post7
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}