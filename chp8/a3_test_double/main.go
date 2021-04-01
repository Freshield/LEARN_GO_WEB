/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: main.py
@Time: 2021-11-16 19:40
@Last_update: 2021-11-16 19:40
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
	"database/sql"
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	_ "github.com/lib/pq"
)

func handleRequest(t Text) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		var err error
		if request.Method != "GET" {
			http.Error(writer, "request method is not get, "+request.Method, http.StatusInternalServerError)
			return
		}
		err = handleGet(writer, request, t)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.fetch(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func main() {
	var err error
	db, err := sql.Open("postgres", "user=gwp dbname=gwp password= sslmode=disable")
	if err != nil {
		panic(err)
	}

	server := http.Server{
		Addr: "127.0.0.1:9666",
	}
	http.HandleFunc("/post/", handleRequest(&Post{Db: db}))
	server.ListenAndServe()
}