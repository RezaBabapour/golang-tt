package main

import "fmt"

var test = 10

func main() {

	fmt.Println(test)

	nested()

	fmt.Println(test)
}

func nested() {
	test = 5
}
