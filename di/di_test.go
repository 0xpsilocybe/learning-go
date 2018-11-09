package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Psilo")
	got := buffer.String()
	want := "Hello, Psilo"
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

