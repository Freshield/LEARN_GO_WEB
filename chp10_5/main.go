/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-11-27 19:42
@Last_update: 2021-11-27 19:42
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
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "0.0.0.0:9666",
	}
	http.HandleFunc("/", handleIndex)
	fmt.Println("Begin the server")
	server.ListenAndServe()
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get request", r.URL.Path)
	w.Write([]byte("Hello World!"))
}