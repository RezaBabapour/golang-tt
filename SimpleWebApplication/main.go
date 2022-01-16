package main

import (
	"fmt"
	"net/http"
)

func main() {
	const portNumber = ":8000"
	http.HandleFunc("/about", about)
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		st := "Hello World"
		fmt.Println(st)
		fmt.Fprintf(rw, st)
	})
	fmt.Println("Starting application on port", portNumber)
	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func about(rw http.ResponseWriter, r *http.Request) {
	sum := addNum(2, 6)
	fmt.Fprintf(rw, "it's me")
	fmt.Fprintf(rw, fmt.Sprintf("Sum is: %d", sum))
}

func addNum(a int, b int) int {
	return a + b
}
