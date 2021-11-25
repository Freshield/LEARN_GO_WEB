/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a1_goroutine.py
@Time: 2021-11-18 10:40
@Last_update: 2021-11-18 10:40
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
	"time"
)

func printNumbers1() {
	for i := 0; i < 10; i++ {
		//fmt.Printf("%d", i)
		time.Sleep(1 * time.Microsecond)
	}
}

func printLetters1() {
	for i := 'A'; i < 'A' + 10; i++ {
		//fmt.Printf("%c", i)
		time.Sleep(1 * time.Microsecond)
	}
}

func print1() {
	printNumbers1()
	printLetters1()
}

func goPrint1() {
	go printNumbers1()
	go printNumbers1()
}

func printNumbers2() {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%d", i)
	}
}

func printLetters2() {
	for i := 'A'; i < 'A' + 10; i++ {
		time.Sleep(1 * time.Microsecond)
		//fmt.Printf("%c", i)
	}
}

func goPrint2() {
	go printNumbers2()
	go printLetters2()
}