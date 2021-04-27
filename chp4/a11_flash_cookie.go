/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a11_flash_cookie.py
@Time: 2021-11-03 12:21
@Last_update: 2021-11-03 12:21
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
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello World!")
	c := http.Cookie{
		Name: "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No message found")
		}
	} else {
		rc := http.Cookie{
			Name: "flash",
			MaxAge: -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)
	server.ListenAndServe()
}