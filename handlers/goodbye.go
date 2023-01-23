package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

// Creates a new hello handler with given logger
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

// ServeHTTP implements the go http.Handler interface
func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye request sent")

	//Response
	rw.Write([]byte("Bye"))
}
