package handlers

import (
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}


func NewProducts(log *log.Logger) *Products  {
	return &Products{log}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	
}