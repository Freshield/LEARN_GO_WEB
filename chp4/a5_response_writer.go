/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a5_response_writer.py
@Time: 2021-11-02 17:42
@Last_update: 2021-11-02 17:42
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

import "net/http"

func writeExample(w http.ResponseWriter, r *http.Request) {
	str := `<html>
<head><title>Go Web Programming</title></head>
<body><h1>Hello World</h1></body>
</html>`
	w.Write([]byte(str))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/", writeExample)
	server.ListenAndServe()
}