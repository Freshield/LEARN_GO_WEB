/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-10-25 21:31
@Last_update: 2021-10-25 21:31
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
	"chp2/index"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	fmt.Println(files)
	fmt.Println(http.StripPrefix("/static/", files))
	fmt.Println(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	mux.HandleFunc("/", index.Index)

	server := &http.Server{
		Addr:    "0.0.0.0:9666",
		Handler: mux,
	}
	log.Fatal(server.ListenAndServe())
}
