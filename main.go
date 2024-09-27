package main

import (
	"context"
	"go-microservice/handlers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	logger := log.New(os.Stdout, "product-api", log.LstdFlags)
	helloHandler := handlers.NewHello(logger)
	goodByeHandler := handlers.NewGoodbye(logger)

	serveMux := http.NewServeMux()
	serveMux.Handle("/", helloHandler)
	serveMux.Handle("/bye", goodByeHandler)

	// handle go server options
	s := &http.Server{
		Addr: ":9090",
		Handler: serveMux,
		IdleTimeout: 120 * time.Second,
		ReadTimeout: 1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func ()  {
		err := s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	logger.Println("Received terminate, gracefully shutdown", sig)

	// Create a context for shutdown with a timeout of 30 seconds
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	// Shut down the server gracefully
	err := s.Shutdown(tc)
	if err != nil {
		logger.Printf("Error during shutdown: %v", err)
	}
}