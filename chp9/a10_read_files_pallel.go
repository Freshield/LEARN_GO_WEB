/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a10_read_files_pallel.py
@Time: 2021-11-22 16:53
@Last_update: 2021-11-22 16:53
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
	"image"
	_ "image/jpeg"
	"io/fs"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println("Start populating tiles db ...")
	db := make([][3]float64, 1000)
	files, _ := ioutil.ReadDir("tiles")
	fmt.Println(len(files))
	var wg sync.WaitGroup
	wg.Add(10)
	//mutex := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		go func(fileList []fs.FileInfo) {
			for _, f := range fileList {
				name := "tiles/" + f.Name()
				file, err := os.Open(name)
				if err == nil {
					_, _, err := image.Decode(file)
					if err == nil {
						db = append(db, [3]float64{0, 0, 0})
					} else {
						fmt.Println("error in populating TILEDB", err, name)
					}
				} else {
					fmt.Println("cannot open file", name, err)
				}
				file.Close()
			}
			wg.Done()
		}(files[i*100:(i+1)*100])
	}
	wg.Wait()

	fmt.Println("Finished populating files db.")
	fmt.Println("Using ", time.Now().Sub(t1))
}