/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a2_config_server.py
@Time: 2021-10-28 15:06
@Last_update: 2021-10-28 15:06
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
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "hello world")
	})
	server := http.Server{
		Addr: "127.0.0.1:9666",
		Handler: mux,
	}
	server.ListenAndServe()
}