/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a3_decoder_xml.py
@Time: 2021-11-13 20:34
@Last_update: 2021-11-13 20:34
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
	"io"
	"os"
)

type Post3 struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author3 `xml:"author"`
	Xml string `xml:",innerxml"`
	Comments []Comment3 `xml:"comments>comment"`
}

type Author3 struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment3 struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author3 `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("post2.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	for {
		t, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			return
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment3
				err = decoder.DecodeElement(&comment, &se)
				if err != nil {
					fmt.Println("Decode element error:", err)
					return
				}
				fmt.Println(comment)
			}
		}
	}
}