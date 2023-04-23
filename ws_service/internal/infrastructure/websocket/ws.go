package websocket

import (
	"log"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// var upgrader = websocket.Upgrader{
// 	CheckOrigin: func(r *http.Request) bool { return true },
// }

// var clients = make(map[*websocket.Conn]bool)
// var broadcast = make(chan []byte)
// var mutex = &sync.Mutex{}

// func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
// 	conn, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println("WebSocket upgrade error:", err)
// 		return
// 	}

// 	// Add new client to the list
// 	mutex.Lock()
// 	clients[conn] = true
// 	mutex.Unlock()

// 	// Set read message deadline
// 	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
// 	if err != nil {
// 		log.Println("WebSocket set read deadline error:", err)
// 		conn.Close()
// 		return
// 	}

// 	// Handle incoming messages from client
// 	for {
// 		_, message, err := conn.ReadMessage()
// 		if err != nil {
// 			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
// 				log.Println("WebSocket closed:", err)
// 			} else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
// 				log.Println("WebSocket unexpected close error:", err)
// 			} else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
// 				log.Println("WebSocket read timeout:", err)
// 			} else {
// 				log.Println("WebSocket read error:", err)
// 			}
// 			break
// 		}
// 		log.Printf("WS Received message: %s", message)

// 		// Check for a close message
// 		if string(message) == "close" {
// 			conn.Close()
// 			break
// 		}

// 		// Broadcast message to all clients
// 		broadcast <- message

// 		// Set read message deadline
// 		err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
// 		if err != nil {
// 			log.Println("WebSocket set read deadline error:", err)
// 			conn.Close()
// 			break
// 		}
// 	}

// 	// Remove client from the list
// 	mutex.Lock()
// 	delete(clients, conn)
// 	mutex.Unlock()

// 	if len(clients) == 0 {
// 		// Close the broadcast channel when all clients are disconnected
// 		close(broadcast)
// 	}
// }

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var mutex = &sync.Mutex{}

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		return
	}

	// Add new client to the list
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Set read message deadline
	err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Println("WebSocket set read deadline error:", err)
		conn.Close()
		return
	}

	// Handle incoming messages from client
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseGoingAway) {
				log.Println("WebSocket closed:", err)
			} else if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("WebSocket unexpected close error:", err)
			} else if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("WebSocket read timeout:", err)
			} else {
				log.Println("WebSocket read error:", err)
			}
			break
		}
		log.Printf("WS Received message: %s", message)

		// Check for a close message
		if string(message) == "close" {
			conn.Close()
			break
		}

		// Broadcast message to all clients
		mutex.Lock()
		for client := range clients {
			err := client.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("server: WebSocket write error:", err)
				client.Close()
				delete(clients, client)
			}
		}
		mutex.Unlock()

		// Set read message deadline
		err = conn.SetReadDeadline(time.Now().Add(10 * time.Second))
		if err != nil {
			log.Println("WebSocket set read deadline error:", err)
			conn.Close()
			break
		}
	}

	// Remove client from the list
	mutex.Lock()
	delete(clients, conn)
	mutex.Unlock()

	if len(clients) == 0 {
		// Close the broadcast channel when all clients are disconnected
	}
}

// func BroadcastMessages() {
// 	for {
// 		message := <-broadcast
// 		log.Printf("Broadcasting message: %s", message)

// 		// Send message to all clients
// 		mutex.Lock()
// 		for conn := range clients {
// 			err := conn.WriteMessage(websocket.TextMessage, message)
// 			if err != nil {
// 				log.Println("WebSocket write error:", err)
// 				conn.Close()
// 				delete(clients, conn)
// 			}
// 		}
// 		mutex.Unlock()
// 	}
// }
