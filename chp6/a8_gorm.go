/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a8_gorm.py
@Time: 2021-11-12 10:08
@Last_update: 2021-11-12 10:08
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
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type Post8 struct {
	Id int
	Content string
	Author string `sql:"not null"`
	Comments []Comment8
	CreateAt time.Time
}

type Comment8 struct {
	Id int
	Content string
	Author string `sql:"not null"`
	Post8Id int `sql:"index"`
	CreateAt time.Time
}

var Db8 *gorm.DB

func init() {
	var err error
	Db8, err = gorm.Open("postgres", "user=gwp dbname=gwp password= sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db8.AutoMigrate(&Post8{}, &Comment8{})
}

func main() {
	post := Post8{Content: "Hellow World!", Author: "Sau Sheong"}
	fmt.Println(post)

	Db8.Create(&post)
	fmt.Println(post)

	comment := Comment8{Content: "Good post!", Author: "Joe"}
	fmt.Println(comment)
	Db8.Model(&post).Association("Comments").Append(comment)

	var readPost Post8
	Db8.Where("author = $1", "Sau Sheong").Last(&readPost)
	fmt.Println(readPost)
	fmt.Println(post)
	var comments []Comment8
	Db8.Model(&readPost).Related(&comments)
	fmt.Println(comments)
}