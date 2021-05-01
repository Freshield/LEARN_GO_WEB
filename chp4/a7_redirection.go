/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_redirection.py
@Time: 2021-11-02 19:22
@Last_update: 2021-11-02 19:22
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

func writeExample7(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func writeHeaderExample7(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample7(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/write", writeExample7)
	http.HandleFunc("/writeheader", writeHeaderExample7)
	http.HandleFunc("/redirect", headerExample7)
	server.ListenAndServe()
}