/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a6_parse_json.py
@Time: 2021-11-14 20:51
@Last_update: 2021-11-14 20:51
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
	"os"
)

type Post6 struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author6 `json:"author"`
	Comments []Comment6 `json:"comments"`
}

type Author6 struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type Comment6 struct {
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
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	var post Post6
	json.Unmarshal(jsonData, &post)
	fmt.Println(post)
}