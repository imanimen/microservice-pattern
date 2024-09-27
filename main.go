package main

import (
	"go-microservice/handlers"
	"log"
	"net/http"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	goodByeHandler := handlers.NewGoodbye(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/bye", goodByeHandler)

	http.ListenAndServe(":9090", serveMux)
}