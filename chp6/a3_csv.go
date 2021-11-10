/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a3_csv.py
@Time: 2021-11-10 16:51
@Last_update: 2021-11-10 16:51
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
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post1 struct {
	Id int
	Content string
	Author string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []*Post1{
		{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		fmt.Println(post.Id)
		fmt.Println(strconv.Itoa(post.Id))
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post1
	for _, item := range record {
		fmt.Printf("%T", item[0])
		fmt.Printf(string(item[0]))
		id, _ := strconv.Atoi(item[0])
		post := Post1{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0])
}