/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a9_cookie.py
@Time: 2021-11-03 10:56
@Last_update: 2021-11-03 10:56
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

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name: "first_cookie",
		Value: "Go Web Programming",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name: "second_cookie",
		Value: "Manning Publications Co lollol",
		HttpOnly: true,
	}
	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	w.Write([]byte("set the cookie"))
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/set_cookie", setCookie)
	server.ListenAndServe()
}