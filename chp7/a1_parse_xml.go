/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a1_parse_xml.py
@Time: 2021-11-13 20:15
@Last_update: 2021-11-13 20:15
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
	"os"
)

type Post1 struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author1 `xml:"author"`
	Xml string `xml:",innerxml"`
}

type Author1 struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	xmlFile, err := os.Open("post1.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	var post Post1
	err = xml.Unmarshal(xmlData, &post)
	fmt.Println(err)
	fmt.Println(post)
}