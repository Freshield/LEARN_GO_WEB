/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: data.py
@Time: 2021-11-16 19:35
@Last_update: 2021-11-16 19:35
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

import "database/sql"

type Text interface {
	fetch(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

type Post struct {
	Db *sql.DB
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

func (post *Post) fetch(id int) (err error) {
	post.Id = id
	return
}

func (post *Post) create() (err error) {
	return
}

func (post *Post) update() (err error) {
	return
}

func (post *Post) delete() (err error) {
	return
}