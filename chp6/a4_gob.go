/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_gob.py
@Time: 2021-11-10 17:40
@Last_update: 2021-11-10 17:40
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
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post4 struct {
	Id int
	Content string
	Author string
}

func store4(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load4(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	post := Post4{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	store4(post, "post1")
	var postRead Post4
	load4(&postRead, "post1")
	fmt.Println(postRead)
}