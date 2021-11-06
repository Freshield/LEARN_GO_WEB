/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a2_condition.py
@Time: 2021-11-04 16:57
@Last_update: 2021-11-04 16:57
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

func process1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl1.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process1)
	server.ListenAndServe()
}