package main

import (
	"fmt"
	"net/http"
	"time"
)

var defaultTimeout = 10 * time.Second

func measureDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, defaultTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	select{
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for '%s' and '%s'", a, b)
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

