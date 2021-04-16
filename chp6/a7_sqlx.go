/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_sqlx.py
@Time: 2021-11-11 19:55
@Last_update: 2021-11-11 19:55
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
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Post7 struct {
	Id int
	Content string
	AuthorName string `db:"author"`
}

var Db7 *sqlx.DB

func init() {
	var err error
	Db7, err = sqlx.Open("postgres", "user=gwp dbname=gwp password= sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetPost7(id int) (post Post7, err error) {
	post = Post7{}
	err = Db7.QueryRowx("select id, content, author from posts where id = $1", id).StructScan(&post)
	if err != nil {
		return
	}
	return
}

func (post *Post7) Create() (err error) {
	err = Db7.QueryRow("insert into posts (content, author) values ($1, $2) returning id",
		post.Content, post.AuthorName).Scan(&post.Id)
	return
}

func main() {
	//post := Post7{Content: "Hello World!", AuthorName: "Sau Sheong"}
	//post.Create()
	//fmt.Println(post)

	post, err := GetPost7(3)
	fmt.Println(post)
	fmt.Println(err)
}