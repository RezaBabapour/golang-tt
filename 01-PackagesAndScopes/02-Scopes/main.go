package main

//file scope

import "fmt"

//package scope

const ok = true

func main() { // block scope starts
	var hello = "hello"

	fmt.Println(hello, ok)
} // block scope ends
