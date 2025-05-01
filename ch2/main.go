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
	mux := http.NewServeMux()
	mux.Handle("/", hh)

	http.ListenAndServe(":9090", mux)
}
