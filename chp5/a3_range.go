/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a3_range.py
@Time: 2021-11-04 17:08
@Last_update: 2021-11-04 17:08
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

func process2(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl2.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sum"}
	t.Execute(w, daysOfWeek)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process2)
	server.ListenAndServe()
}