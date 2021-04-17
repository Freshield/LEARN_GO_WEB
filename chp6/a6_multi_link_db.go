/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a6_multi_link_db.py
@Time: 2021-11-11 15:20
@Last_update: 2021-11-11 15:20
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
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type Post6 struct {
	Id int
	Content string
	Author string
	Comments []Comment
}

type Comment struct {
	Id int
	Content string
	Author string
	Post *Post6
}

var Db6 *sql.DB

func init() {
	var err error
	Db6, err = sql.Open("postgres", "user=gwp dbname=gwp password= sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func (comment *Comment) Create() (err error) {
	if comment.Post == nil {
		err = errors.New("Post not found")
		return
	}
	err = Db6.QueryRow("insert into comments (content, author, post_id)" +
		"values ($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(
			&comment.Id)
	return
}

func GetPost6(id int) (post Post6, err error) {
	post = Post6{}
	post.Comments = []Comment{}
	err = Db6.QueryRow("select id, content, quthor from posts where id = $1", id).Scan(
		&post.Id, &post.Content, &post.Author)

	rows, err := Db6.Query("select id, content, author from comments where post_id = $1", id)
	if err != nil {
		return
	}
	for rows.Next() {
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil {
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	rows.Close()
	return
}

func (post *Post6) Create() (err error) {
	err = Db6.QueryRow("insert into posts (content, author) values ($1, $2)" +
		"returning id", post.Content, post.Author).Scan(&post.Id)
	return
}

func main() {
	post := Post6{Content: "Hello World!", Author: "Sau Sheong"}
	post.Create()

	comment := Comment{Content: "Good post!", Author: "Joe", Post: &post}
	comment.Create()
	readPost, _ := GetPost6(post.Id)

	fmt.Println(readPost)
	fmt.Println(readPost.Comments)
	fmt.Println(readPost.Comments[0].Post)
}