/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a5_msg_transfer.py
@Time: 2021-11-18 15:38
@Last_update: 2021-11-18 15:38
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

func thrower(c chan int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Println("Threw >>", i)
		time.Sleep(1 * time.Microsecond)
	}
}

func catcher(c chan int) {
	for i := 0; i < 5; i++ {
		num := <- c
		fmt.Println("Caught <<", num)
		time.Sleep(1 * time.Microsecond)
	}
}

func main() {
	c := make(chan int, 3)
	go thrower(c)
	go catcher(c)
	time.Sleep(100 * time.Millisecond)
}