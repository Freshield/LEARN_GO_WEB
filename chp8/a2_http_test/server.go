/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: server.py
@Time: 2021-11-15 20:29
@Last_update: 2021-11-15 20:29
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
	"path"
	"strconv"
)

func handleRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Println("The method is not get:", r.Method)
		return
	}
	err := handleGet(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"id":"` + strconv.Itoa(id) + `"}`))
	return
}