package main

import (
	"htmlTemplate/pkg/handler"
	"log"
	"net/http"
)

func main() {
	log.Println("App is running")
	http.HandleFunc("/", handler.Home)
	http.ListenAndServe(":8000", nil)

}
