package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("hello world")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "Request Body: %s", d)
	fmt.Fprintf(rw, "incomming request header:  %s", r.Header)
	return
}
