/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a11_diff_layout.py
@Time: 2021-11-07 15:24
@Last_update: 2021-11-07 15:24
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
	"math/rand"
	"net/http"
	"time"
)

func process11(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("t11.html", "red_hello.html")
	} else {
		t, _ = template.ParseFiles("t11.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process11)
	server.ListenAndServe()
}