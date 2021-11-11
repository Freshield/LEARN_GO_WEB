/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a2_write_file.py
@Time: 2021-11-10 16:22
@Last_update: 2021-11-10 16:22
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
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello World!\n")
	err := ioutil.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}
	read1, _ := ioutil.ReadFile("data1")
	fmt.Println(string(read1))

	file1, _ := os.Create("data2")
	defer file1.Close()
	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to filen", bytes)
	file2, _ := os.Open("data2")
	defer file2.Close()
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}