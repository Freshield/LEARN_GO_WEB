/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a1_store_as_map.py
@Time: 2021-11-10 15:52
@Last_update: 2021-11-10 15:52
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

import "fmt"

type Post struct {
	Id int
	Content string
	Author string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post *Post) {
	PostById[post.Id] = post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], post)
}

func main() {
	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := &Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := &Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := &Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := &Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	post1.Content = "test"

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
	fmt.Println("----------------")
	for _, post := range PostById {
		fmt.Println(post)
	}
}