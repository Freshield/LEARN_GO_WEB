/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_file_upload.py
@Time: 2021-11-02 14:55
@Last_update: 2021-11-02 14:55
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
	"net/http"
)

func process4(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(1024)
	//fmt.Fprintln(w, r.MultipartForm.File["uploaded"])
	//fileHeader := r.MultipartForm.File["uploaded"][0]
	//file, err := fileHeader.Open()
	file, header, err := r.FormFile("uploaded")
	fmt.Fprintln(w, header)
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/process", process4)
	server.ListenAndServe()
}