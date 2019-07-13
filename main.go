package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/ request")
		fmt.Fprintf(w, "Hello World\n")
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("/ping request")
		fmt.Fprintf(w, "pong\n")
	})

	server := &http.Server{
		Addr: ":3000",
	}

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
