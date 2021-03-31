/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a9_mosic_pallel.py
@Time: 2021-11-22 16:08
@Last_update: 2021-11-22 16:08
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
	"bytes"
	"chp9/mosic_pallel"
	"encoding/base64"
	"fmt"
	"html/template"
	"image"
	"image/jpeg"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", upload1)
	mux.HandleFunc("/mosaic", mosaic1)
	server := &http.Server{
		Addr: "127.0.0.1:9666",
		Handler: mux,
	}
	t1 := time.Now()
	mosic_pallel.TILESDB = mosic_pallel.TilesDB()
	fmt.Println("Prepare for", time.Now().Sub(t1), " seconds")
	fmt.Println("Mosaic server started.")
	server.ListenAndServe()
}

func upload1(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("upload.html")
	t.Execute(w, nil)
}

func mosaic1(w http.ResponseWriter, r *http.Request) {
	t0 := time.Now()

	r.ParseMultipartForm(10485760)
	file, _, _ := r.FormFile("image")
	defer file.Close()
	tileSize, _ := strconv.Atoi(r.FormValue("tile_size"))

	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	db := mosic_pallel.CloneTilesDB()

	c1 := mosic_pallel.Cut(original, db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := mosic_pallel.Cut(original, db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := mosic_pallel.Cut(original, db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := mosic_pallel.Cut(original, db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)
	c := mosic_pallel.Combine(bounds, c1, c2, c3, c4)

	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string {
		"original": originalStr,
		"mosaic": <-c,
		"duration": fmt.Sprintf("%v", t1.Sub(t0)),
	}
	t, _ := template.ParseFiles("results.html")
	t.Execute(w, images)
}
