package main

import (
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("Error running http server: %s\n", err)
	}
}

func run() error {
	r := router()
	err := http.ListenAndServe(":4242", r)
	if err != nil {
		return err
	}

	return nil
}
