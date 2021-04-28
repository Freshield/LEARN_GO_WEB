/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a10_get_cookie.py
@Time: 2021-11-03 12:10
@Last_update: 2021-11-03 12:10
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

func setCookie1(w http.ResponseWriter, r *http.Request) {
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

func getCookie1(w http.ResponseWriter, r *http.Request) {
	//h := r.Header["Cookie"]
	//fmt.Fprintln(w, h)
	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	cs := r.Cookies()
	fmt.Fprintln(w, c1)
	fmt.Fprintln(w, cs)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/set_cookie", setCookie1)
	http.HandleFunc("/get_cookie", getCookie1)
	server.ListenAndServe()
}