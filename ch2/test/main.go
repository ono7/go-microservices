package main

import (
	"log"
	"net/http"
	"os"

	"main/test/handlers"
)

func main() {

	l := log.New(os.Stdout, "api: ", log.LstdFlags)

	hh := handlers.NewHello(l)
	mux := http.NewServeMux()
	mux.Handle("/hello", hh)

	http.ListenAndServe(":9090", mux)
}
