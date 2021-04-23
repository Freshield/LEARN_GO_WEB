/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_funcmap.py
@Time: 2021-11-07 14:00
@Last_update: 2021-11-07 14:00
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
	"time"
)

func formatDate(t time.Time) string {
	layout := "2006-01-02"
	return t.Format(layout)
}

func process7(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("t7.html").Funcs(funcMap)
	t, _ = t.ParseFiles("t7.html")
	t.Execute(w, time.Now())
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", process7)
	server.ListenAndServe()
}