package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func HttpGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world!")
}

func main() {
	host := ":5000"
	fmt.Printf("Server listening on %s", host)
	http.ListenAndServe(host, http.HandlerFunc(HttpGreetHandler))
}

