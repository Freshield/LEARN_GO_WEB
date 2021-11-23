/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a3_wait_goroutine.py
@Time: 2021-11-18 11:06
@Last_update: 2021-11-18 11:06
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
)

func printNumbersWait2(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
	}
	wg.Done()
}

func printLettersWait2(wg *sync.WaitGroup) {
	for i := 'A'; i < 'A' + 10; i++ {
		fmt.Printf("%c", i)
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbersWait2(&wg)
	go printLettersWait2(&wg)
	wg.Wait()
}