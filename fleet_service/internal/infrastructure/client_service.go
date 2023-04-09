package infrastructure

import (
	"log"
	"net/url"
	"platform_science_code_exercise/internal/domain/entity"

	"github.com/gorilla/websocket"
)

func ClientService(message []entity.MessageResponse, analytic_path string) {
	u := url.URL{Scheme: "ws", Host: analytic_path, Path: "/ws"}
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatalf("failed to connect to WebSocket server: %v", err)
	}
	defer c.Close()

	// Start a goroutine to read messages from the WebSocket server
	go func() {
		for {
			var msg entity.MessageResponse
			err := c.ReadJSON(&msg)
			if err != nil {
				log.Printf("error reading message: %v", err)
				break
			}
			// Process the incoming message
		}
	}()

	// Example: send a message to the server
	err = c.WriteJSON(message)
	if err != nil {
		log.Fatalf("error sending message: %v", err)
	}
}
