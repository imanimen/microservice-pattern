package handlers

import (
	"log"
	"net/http"
)

type Products struct {
	logger *log.Logger
}


func NewProducts(l *log.Logger) *Products  {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, h *http.Request) {
	
}