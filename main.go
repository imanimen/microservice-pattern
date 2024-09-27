package main

import (
	"log"
	"net/http"
)

func main() {
	// handler for our web server
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello World")
	})
	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":9090", nil)
}
