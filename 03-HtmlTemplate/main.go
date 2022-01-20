package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("App is running")
	http.HandleFunc("/", home)
	http.ListenAndServe(":8000", nil)

}
