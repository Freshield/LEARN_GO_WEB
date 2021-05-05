/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a3_parse_form.py
@Time: 2021-11-01 17:25
@Last_update: 2021-11-01 17:25
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
	"log"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//fmt.Fprintln(w, r.Form)
	//fmt.Fprintln(w, r.Form["hello"])
	//fmt.Fprintln(w, r.PostForm)
	r.ParseMultipartForm(1024)
	fmt.Fprintln(w, r.MultipartForm)
	fmt.Fprintln(w, r.FormValue("hello"))
	fmt.Fprintln(w, len(r.FormValue("hello")))
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, len(r.Form["hello"]))
	fmt.Fprintln(w, r.Form["hello"][0])
	fmt.Fprintln(w, r.Form["hello"][1])
	fmt.Fprintln(w, len(r.MultipartForm.Value["hello"]))
	fmt.Fprintln(w, r.PostFormValue("hello"))
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/process", process)
	log.Fatal(server.ListenAndServe())
}