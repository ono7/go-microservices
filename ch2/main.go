package main

import (
	"log"
	"net/http"
	"os"

	"main/handlers"
)

func main() {

	l := log.New(os.Stdout, "Service: ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)
	mux := http.NewServeMux()
	mux.Handle("/hello", hh)
	mux.Handle("/goodbye", gb)

	http.ListenAndServe(":9090", mux)
}
