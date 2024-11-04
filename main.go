package main

import (
	"fmt"
	"log"
	"net/http"
	"webserver/server"
)

func main() {
	http.HandleFunc("/", server.HomeHandler)
	http.HandleFunc("/about", server.AboutHandler)
	http.HandleFunc("/users", server.UsersHandler)
	http.HandleFunc("/api/status", server.ApiStatusHandler)

	fmt.Println("Server starting on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server failed to start:", err)
	}
}
