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

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)

	http.ListenAndServe(":9090", serveMux)
}