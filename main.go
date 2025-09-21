package main

import (
	"flag"
	"fmt"
	"go_url_shortener/handlers"
	"go_url_shortener/models"
	"net/http"
)



func main() {
	models.ConnectDatabase()

	port := flag.String("port", "8080", "Port to run the load balancer on")
	flag.Parse()

	fmt.Println("Starting server on :" + *port + " -> http://localhost:" + *port)
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)
	http.ListenAndServe(":"+*port, nil)
}