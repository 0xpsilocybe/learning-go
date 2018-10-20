package main

import "testing"

func TestHello(t *testing.T) {

  assertCorrectMessage := func(t *testing.T, got string, want string) {
    t.Helper()
    if got != want {
      t.Errorf("got '%s' want '%s'", got, want)
    } 
  }

  t.Run("Saying hello to people", func(t *testing.T) {
    got := Hello("Bartek", "")
    want := "Hello, Bartek"
    assertCorrectMessage(t, got, want)
  })

  t.Run("Say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
    got := Hello("", "")
    want := "Hello, World"
    assertCorrectMessage(t, got, want)
  })

  t.Run("Say hello in Spanish", func(t *testing.T) {
	  got := Hello("Elodie", "Spanish")
	  want := "Hola, Elodie"
	  assertCorrectMessage(t, got, want)
  })

  t.Run("Say hello in French", func(t *testing.T) {
	  got := Hello("Jean", "French")
	  want := "Bonjour, Jean"
	  assertCorrectMessage(t, got, want)
  })
}

