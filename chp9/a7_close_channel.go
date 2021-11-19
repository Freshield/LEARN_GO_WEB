/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a7_close_channel.py
@Time: 2021-11-18 16:00
@Last_update: 2021-11-18 16:00
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

import "fmt"

func callerCloseA(c chan string) {
	c <- "Hello World!"
	close(c)
}

func callerCloseB(c chan string) {
	c <- "Hola Mundo!"
	close(c)
}

func main() {
	a, b := make(chan string), make(chan string)
	go callerCloseA(a)
	go callerCloseB(b)
	var msg string
	ok1, ok2 := true, true
	for ok1 || ok2 {
		select {
		case msg, ok1 = <-a:
			if ok1 {
				fmt.Println(msg, " from A")
			}
		case msg, ok2 = <-b:
			if ok2 {
				fmt.Println(msg, " from B")
			}		
		}
	}
}