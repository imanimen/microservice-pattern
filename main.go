package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// handler for our web server
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Hello World")
		data, err := io.ReadAll(request.Body)
		if err != nil {
			http.Error(writer, "Oops", http.StatusBadRequest)
			return
		}
		// write back to user
		fmt.Fprintf(writer, "Hello %s\n", data) 
	})
	http.HandleFunc("/goodbye", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("Goodbye World")
	})
	http.ListenAndServe(":9090", nil)
}
