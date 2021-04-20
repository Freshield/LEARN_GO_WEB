/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a10_define.py
@Time: 2021-11-07 15:20
@Last_update: 2021-11-07 15:20
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

func process10(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("t10.html")
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process10)
	server.ListenAndServe()
}