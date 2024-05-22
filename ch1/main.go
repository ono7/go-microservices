package main

import (
	"ch1/handlers"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// create new log with predend of product-api: for each log entry, has standard log flags
	port := ":9090"
	l := log.New(os.Stdout, "product-api: ", log.LstdFlags)
	gl := log.New(os.Stdout, "goodbye-handler: ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(gl)

	// create a new server Mux that can handle incomming requests
	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// create custom server to run with some optimizations
	// TODO(jlima): lookup optimizations for slow/high latency links
	s := &http.Server{
		Addr:         port,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}

	fmt.Printf("Listening on http://localhost%s", port)

	// graceful shutdown pattern
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigchan := make(chan os.Signal)
	signal.Notify(sigchan, os.Interrupt)
	signal.Notify(sigchan, os.Kill)

	sig := <-sigchan

	l.Println("Recieved terminate signal, gracefull shutdown", sig)
	tc, _ := context.WithTimeout(context.Background(), 10*time.Second)
	s.Shutdown(tc)

}
