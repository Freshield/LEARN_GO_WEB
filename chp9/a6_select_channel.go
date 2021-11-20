/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a6_select_channel.py
@Time: 2021-11-18 15:46
@Last_update: 2021-11-18 15:46
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
	"time"
)

func callerA(c chan string) {
	c <- "Hello World!"
}

func callerB(c chan string) {
	c <- "Hola Mundo!"
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerA(a)
	go callerB(b)
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Microsecond)
		select {
		case msg := <-a:
			fmt.Println(msg, " from A")
		case msg := <-b:
			fmt.Println(msg, " from B")
		default:
			fmt.Println("Default")
		}
	}
}