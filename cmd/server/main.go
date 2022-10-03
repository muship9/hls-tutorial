package main

import (
	"hls-tutorial/pkg"
	"log"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", pkg.HandleRequest)
	err := server.ListenAndServe()
	if err != nil {
		log.Print("Server Error", err)
	}
}
