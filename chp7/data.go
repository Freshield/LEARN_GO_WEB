/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: data.py
@Time: 2021-11-15 12:18
@Last_update: 2021-11-15 12:18
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
	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password= sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func retrieve(id int) (post Post10, err error) {
	post = Post10{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(
		&post.Id, &post.Content, &post.Author)
	return
}

func (post *Post10) create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post10) update() (err error) {
	_, err = Db.Exec(
		"update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post Post10) delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}