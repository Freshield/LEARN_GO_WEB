/*
@Author: Freshield
@Contact: yangyufresh@163.com
@File: a2_test_goroutine.py
@Time: 2021-11-18 10:43
@Last_update: 2021-11-18 10:43
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
	"testing"
	"time"
)

func TestPrint1(t *testing.T) {
	print1()
}

func TestGoPrint1(t *testing.T) {
	goPrint1()
	time.Sleep(1 * time.Millisecond)
}

func TestGoPrint2(t *testing.T) {
	goPrint2()
	time.Sleep(1 * time.Millisecond)
}

func BenchmarkPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		print1()
	}
}

func BenchmarkGoPrint1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint1()
	}
}

func BenchmarkGoPrint2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goPrint2()
	}
}