package main

import (
	"Microservices/week_2/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// create the log var importing from our handlers folder
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	// create the handlers
	ph := handlers.NewProducts(l)
	// create new servemux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// lets manually create a server
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()
	// use the os.Signal package to register for the notification of certain signals
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	// create a timeout context
	tc, err := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		l.Fatal(err)
	}
	s.Shutdown(tc)
}
