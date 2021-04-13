/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_create_xml.py
@Time: 2021-11-13 20:41
@Last_update: 2021-11-13 20:41
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
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post4 struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id"`
	Content string `xml:"content"`
	Author Author4 `xml:"author"`
}

type Author4 struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post4{
		Id: "1",
		Content: "Hello World!",
		Author: Author4{
			Id: "2",
			Name: "Sau Sheong",
		},
	}

	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = ioutil.WriteFile("post4.xml", []byte(xml.Header + string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}