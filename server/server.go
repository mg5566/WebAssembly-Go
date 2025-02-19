package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Server is starting...")

	go func() {
		err := http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
		if err != nil {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	log.Println("Server is running at http://localhost:8080")

	// Keep the main function running
	select {}
}
