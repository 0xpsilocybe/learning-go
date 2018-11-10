package main

import (
	"net/http"
	"time"
)

func measureDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer(a, b string) (winner string) {
	select{
	case <- ping(a):
		return a
	case <- ping(b):
		return b
	}
}

func ping(url string) chan bool {
	channel := make(chan bool)
	go func() {
		http.Get(url)
		channel <- true
	}()
	return channel
}

