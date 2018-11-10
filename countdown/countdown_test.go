package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
	sleep = "sleep"
	write = "write"
)

type CountdownOperationSpy struct {
	Calls []string
}

func (s *CountdownOperationSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {

	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &CountdownOperationSpy{})
		got := buffer.String()
		want := 
`3
2
1
Go!`
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("sleeps after every print but last", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string {
			write, // Output: 3
			sleep, // >delay<
			write, // Output: 2
			sleep, // >delay<
			write, // Output: 1
			sleep, // >delay<
			write, // Output: Go1
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, got %v", want, spySleepPrinter.Calls	)
		}
	})

}

