package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"platform_science_code_exercise/config"
	"platform_science_code_exercise/internal/domain/entity"
	"platform_science_code_exercise/internal/domain/service"
	"platform_science_code_exercise/internal/infrastructure"

	"github.com/gorilla/websocket"
)

type WSConfigServer struct {
	c config.Config
}

func (c WSConfigServer) WSFileServer() {
	_, _, p, _ := c.c.ServerConfig()
	port := fmt.Sprintf(":%s", p)
	// Create a new WebSocket upgrader
	upgrader := websocket.Upgrader{}

	// Set up a WebSocket endpoint
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("connection_error: ", err)
			return
		}

		// Handle incoming messages
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("error: %v", err)
				}
				return
			}

			// Decode the message as JSON
			var message entity.Message
			err = json.Unmarshal(p, &message)
			if err != nil {
				log.Println("2", err)
				return
			}

			// Print the message to the console
			result := service.PlatformScienceCodeExercise(message)
			infrastructure.ClientService(result, c.c.JSONAnalyticsConfig())

			// Echo the message back to the client
			err = conn.WriteMessage(messageType, p)
			if err != nil {
				log.Println("3", err)
				return
			}
		}
	})

	// Start the server
	err := http.ListenAndServe(port, nil)
	log.Println("fleet_server starts ....")
	if err != nil {
		log.Fatal(err)
	}
}
