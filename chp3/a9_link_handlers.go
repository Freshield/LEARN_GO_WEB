/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a9_link_handlers.py
@Time: 2021-11-01 14:29
@Last_update: 2021-11-01 14:29
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

type HelloHandler3 struct {}

func (h HelloHandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log3(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(writer, request)
	})
}

func protect3(h http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("here is protect")
		h.ServeHTTP(writer, request)
	})
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.Handle("/hello/", protect3(log3(HelloHandler3{})))
	server.ListenAndServe()
}

