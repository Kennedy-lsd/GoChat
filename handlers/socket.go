package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Kennedy-lsd/GoChat/data"
	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan string)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Server error")
		return
	}
	defer ws.Close()

	clients[ws] = true
	fmt.Println("New client connected!")

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			delete(clients, ws)
			break
		}

		var message data.Message
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println("Error parsing message:", err)
			continue
		}

		message.Time = time.Now().Format(time.RFC3339)

		messageJSON, err := json.Marshal(message)
		if err != nil {
			log.Println("Error marshaling message:", err)
			continue
		}

		log.Printf("Received message from %s: %s at %s", message.Username, message.Content, message.Time)
		broadcast <- (string(messageJSON))
	}
}

func HandleMessages() {
	for msg := range broadcast {
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Println("Error sending message:", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
