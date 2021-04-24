/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a6_pipe.py
@Time: 2021-11-07 13:45
@Last_update: 2021-11-07 13:45
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

func process6(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t6.html")
	t.Execute(w, "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process6)
	server.ListenAndServe()
}