/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a5_encoder_xml.py
@Time: 2021-11-13 20:50
@Last_update: 2021-11-13 20:50
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
	"os"
)

type Post5 struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author5 `xml:"author"`
}

type Author5 struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post5{
		Id: "1",
		Content: "Hello World!",
		Author: Author5{
			Id: "2",
			Name: "Sau Sheong",
		},
	}

	xmlFile, err := os.Create("post5.xml")
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
		return
	}
}