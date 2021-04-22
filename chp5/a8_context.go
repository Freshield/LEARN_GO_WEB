/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a8_context.py
@Time: 2021-11-07 14:34
@Last_update: 2021-11-07 14:34
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

func process8(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t8.html")
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process8)
	server.ListenAndServe()
}