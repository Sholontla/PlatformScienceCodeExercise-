package websocket

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan []byte)
var mutex = &sync.Mutex{}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

	// Add new client to the list
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Handle incoming messages from client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("WebSocket read error:", err)
			break
		}
		log.Printf("WS Received message: %s", message)

		// Broadcast message to all clients
		broadcast <- message
	}

	// Remove client from the list
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()
}

func BroadcastMessages() {
	for {
		message := <-broadcast
		log.Printf("Broadcasting message: %s", message)

		// Send message to all clients
		mutex.Lock()
		for conn := range clients {
			err := conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("WebSocket write error:", err)
				conn.Close()
				delete(clients, conn)
			}
		}
		mutex.Unlock()
	}
}
