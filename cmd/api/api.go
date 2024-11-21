package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kennedy-lsd/GoChat/handlers"
)

func Api() {
	const PORT int = 8080

	http.HandleFunc("/ws", handlers.HandleConnections)

	log.Printf("Server starts at %d port", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), nil); err != nil {
		log.Fatalf("Error starting server on %d", PORT)
	}
}
