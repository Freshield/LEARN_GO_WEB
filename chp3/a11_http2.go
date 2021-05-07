/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a11_http2.py
@Time: 2021-11-01 15:03
@Last_update: 2021-11-01 15:03
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
	"golang.org/x/net/http2"
	"net/http"
)

type MyHandler6 struct {}

func (h *MyHandler6) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	handler := MyHandler6{}
	server := http.Server{
		Addr: "127.0.0.1:9666",
		Handler: &handler,
	}
	http2.ConfigureServer(&server, &http2.Server{})
	server.ListenAndServeTLS("cert.pem", "key.pem")
}