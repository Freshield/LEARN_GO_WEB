/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_with.py
@Time: 2021-11-04 17:18
@Last_update: 2021-11-04 17:18
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
	"html/template"
	"net/http"
)

func process3(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl3.html")
	t.Execute(w, "hello")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process3)
	server.ListenAndServe()
}