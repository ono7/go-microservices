package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("goodbye world")
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "oops", http.StatusBadRequest)
	}
	fmt.Fprintf(rw, "(goodbye) Request Body: %s\n", d)
	fmt.Fprintf(rw, "incomming request header:  %s\n", r.Header)
	return
}
