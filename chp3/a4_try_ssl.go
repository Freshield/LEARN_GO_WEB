/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_try_ssl.py
@Time: 2021-10-28 17:38
@Last_update: 2021-10-28 17:38
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

type customerHandler struct {}

func (c customerHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
		Handler: customerHandler{},
	}
	server.ListenAndServeTLS("cert.pem", "key.pem")
}