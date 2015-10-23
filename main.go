package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ksubedi/go-web-seed/config"
)

func main() {
	// Configuration
	PORT := config.Get("WEBSERVER_PORT")

	// Get the router from router.go
	router := getRouter()

	//Run HTTP Server
	fmt.Println("Running WebServer on Port " + PORT)

	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
