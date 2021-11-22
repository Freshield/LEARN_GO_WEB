/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a4_channel_goroutine.py
@Time: 2021-11-18 15:24
@Last_update: 2021-11-18 15:24
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
	"sync"
	"time"
)

var wg sync.WaitGroup

func printNumbersChannel2(w chan bool) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%d", i)
	}
	fmt.Println("begin send channel numbers")
	w <- true
	fmt.Println("done send channel numbers")
	time.Sleep(1 * time.Second)
	wg.Done()
}

func printLettersChannel2(w chan bool) {
	for i := 'A'; i < 'A' + 10; i++ {
		time.Sleep(1 * time.Microsecond)
		fmt.Printf("%c", i)
	}
	fmt.Println("begin send channel letters")
	w <- true
	fmt.Println("done send channel letters")
	time.Sleep(1 * time.Second)
	wg.Done()
}

func main() {
	w1, w2 := make(chan bool), make(chan bool)
	wg.Add(2)
	go printNumbersChannel2(w1)
	go printLettersChannel2(w2)
	fmt.Println("Done create goroutine")
	time.Sleep(1 * time.Second)
	<-w1
	time.Sleep(1 * time.Second)
	<-w2
	time.Sleep(1 * time.Second)
	fmt.Println("Done all")
	wg.Wait()
}