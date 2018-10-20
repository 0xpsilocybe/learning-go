package main

import "fmt"

const helloPrefixEnglish = "Hello, "
const spanish = "Spanish"
const helloPrefixSpanish = "Hola, "
const french = "French"
const helloPrefixFrench = "Bonjour, "
const worldSuffix = "World"

func getGreetingPrefix(language string) (prefix string){
  switch language {
  case french:
	prefix = helloPrefixFrench
  case spanish:
    prefix = helloPrefixSpanish
  default:
	prefix = helloPrefixEnglish
  }
  return
}

func Hello(name string, language string) string {
  if name == "" {
    name = worldSuffix
  }
  return getGreetingPrefix(language) + name
}

func main() {
  fmt.Println(Hello("", ""))
}

