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
	aDuration := measureDuration(a)
	bDuration := measureDuration(b)
	if aDuration < bDuration {
		return a
	}
	return b
}

