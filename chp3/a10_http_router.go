/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a10_http_router.py
@Time: 2021-11-01 14:53
@Last_update: 2021-11-01 14:53
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
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func hello5(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello5)

	server := http.Server{
		Addr: "127.0.0.1:9666",
		Handler: mux,
	}
	server.ListenAndServe()
}