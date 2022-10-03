package main

import (
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Print("Server Error", err)
	}
}
