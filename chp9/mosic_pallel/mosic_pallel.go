/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: mosic_pallel.py
@Time: 2021-11-22 16:09
@Last_update: 2021-11-22 16:09
@Desc: None
@==============================================@
@      _____             _   _     _   _       @
@     |   __|___ ___ ___| |_|_|___| |_| |      @
@     |   __|  _| -_|_ -|   | | -_| | . |      @
@     |__|  |_| |___|___|_|_|_|___|_|___|      @
@                                    Freshield @
@==============================================@
*/
package mosic_pallel

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io/fs"
	"io/ioutil"
	"math"
	"os"
	"sync"
)

type DB struct {
	mutex *sync.Mutex
	store map[string][3]float64
}

var TILESDB *DB

func CloneTilesDB() *DB {
	tilesDb := &DB{}
	db := make(map[string][3]float64)
	for k, v := range TILESDB.store {
		db[k] = v
	}
	tilesDb.store = db
	tilesDb.mutex = &sync.Mutex{}
	return tilesDb
}

func averageColor(img image.Image) [3]float64 {
	bounds := img.Bounds()
	r, g, b := 0.0, 0.0, 0.0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r1, g1, b1, _ := img.At(x, y).RGBA()
			r, g, b = r+float64(r1), g+float64(g1), b+float64(b1)
		}
	}
	totalPixels := float64(bounds.Max.X * bounds.Max.Y)
	return [3]float64{r / totalPixels, g / totalPixels, b / totalPixels}
}

func Resize(in image.Image, newWidth int) image.NRGBA {
	bounds := in.Bounds()
	ratio := bounds.Dx() / newWidth
	out := image.NewNRGBA(image.Rect(
		bounds.Min.X / ratio, bounds.Min.Y / ratio, bounds.Max.X / ratio, bounds.Max.Y / ratio))
	for y, j := bounds.Min.Y, bounds.Min.Y; y < bounds.Max.Y; y, j = y+ratio, j+1 {
		for x, i := bounds.Min.X, bounds.Min.X; x < bounds.Max.X; x, i = x+ratio, i+1 {
			r, g, b, a := in.At(x, y).RGBA()
			out.SetNRGBA(i, j, color.NRGBA{
				uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8),
			})
		}
	}
	return *out
}

func TilesDB() *DB {
	tilesDb := &DB{}
	fmt.Println("Start populating tiles db ...")
	db := make(map[string][3]float64)
	files, _ := ioutil.ReadDir("tiles")
	mux := sync.Mutex{}
	fileLength := len(files)
	fmt.Println(fileLength)
	//os.Exit(0)
	for i := 0; i < 100; i++ {
		go func(fileList []fs.FileInfo) {
			for _, f := range fileList {
				name := "tiles/" + f.Name()
				file, err := os.Open(name)
				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						mux.Lock()
						db[name] = averageColor(img)
						mux.Unlock()
					} else {
						fmt.Println("error in populating TILEDB", err, name)
					}
				} else {
					fmt.Println("cannot open file", name, err)
				}
				file.Close()
			}
		}(files[i*100: (i+1)*100])
	}
	//for _, f := range files {
	//	name := "tiles/" + f.Name()
	//	file, err := os.Open(name)
	//	if err == nil {
	//		img, _, err := image.Decode(file)
	//		if err == nil {
	//			db[name] = averageColor(img)
	//		} else {
	//			fmt.Println("error in populating TILEDB", err, name)
	//		}
	//	} else {
	//		fmt.Println("cannot open file", name, err)
	//	}
	//	file.Close()
	//}


	fmt.Println("Finished populating files db.")
	tilesDb.store = db
	tilesDb.mutex = &sync.Mutex{}
	return tilesDb
}

func (db *DB) nearest(target [3]float64) string {
	var filename string
	db.mutex.Lock()
	smallest := 10000000.0
	for k, v := range db.store {
		dist := distance(target, v)
		if dist < smallest {
			filename, smallest = k, dist
		}
	}
	delete(db.store, filename)
	db.mutex.Unlock()
	return filename
}

func distance(p1 [3]float64, p2 [3]float64) float64 {
	return math.Sqrt(sq(p2[0]-p1[0]) + sq(p2[1]-p1[1]) + sq(p2[2]-p1[2]))
}

func sq(n float64) float64 {
	return n * n
}

func Cut(original image.Image, db *DB, tileSize, x1, y1, x2, y2 int) <-chan image.Image {
	c := make(chan image.Image)

	sp := image.Point{0, 0}
	go func() {
		newimage := image.NewNRGBA(image.Rect(x1, y1, x2, y2))
		for y := y1; y < y2; y = y + tileSize {
			for x := x1; x < x2; x = x + tileSize {
				r, g, b, _ := original.At(x, y).RGBA()
				color := [3]float64{float64(r), float64(g), float64(b)}

				nearest := db.nearest(color)
				file, err := os.Open(nearest)
				if err == nil {
					img, _, err := image.Decode(file)
					if err == nil {
						t := Resize(img, tileSize)
						tile := t.SubImage(t.Bounds())
						tileBounds := image.Rect(x, y, x+tileSize, y+tileSize)
						draw.Draw(newimage, tileBounds, tile, sp, draw.Src)
					} else {
						fmt.Println("error:", err, nearest)
					}
				} else {
					fmt.Println("error:", nearest)
				}
				file.Close()
			}
		}
		c <- newimage.SubImage(newimage.Rect)
	}()
	return c
}

func Combine(r image.Rectangle, c1, c2, c3, c4 <-chan image.Image) <-chan string {
	c := make(chan string)

	go func() {
		var wg sync.WaitGroup
		img := image.NewNRGBA(r)
		copy := func(dst draw.Image, r image.Rectangle, src image.Image, sp image.Point) {
			draw.Draw(dst, r, src, sp, draw.Src)
			wg.Done()
		}
		wg.Add(4)
		var s1, s2, s3, s4 image.Image
		var ok1, ok2, ok3, ok4 bool
		for {
			select {
			case s1, ok1 = <-c1:
				go copy(img, s1.Bounds(), s1, image.Point{r.Min.X, r.Min.Y})
			case s2, ok2 = <-c2:
				go copy(img, s2.Bounds(), s2, image.Point{r.Max.X/2, r.Min.Y})
			case s3, ok3 = <-c3:
				go copy(img, s3.Bounds(), s3, image.Point{r.Min.X, r.Max.Y/2})
			case s4, ok4 = <-c4:
				go copy(img, s4.Bounds(), s4, image.Point{r.Max.X/2, r.Max.Y/2})

			}
			if (ok1 && ok2 && ok3 && ok4) {
				break
			}
		}
		wg.Wait()
		buf2 := new(bytes.Buffer)
		jpeg.Encode(buf2, img, nil)
		c <- base64.StdEncoding.EncodeToString(buf2.Bytes())
	}()
	return c
}