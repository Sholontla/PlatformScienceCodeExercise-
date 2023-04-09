package websocket

import (
	"encoding/json"
	"file_server/config"
	"file_server/internal/domain/entity"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type ClientProducer struct {
	c config.Config
}

func (pr ClientProducer) WSFileClient(message entity.Message) {

	ws, host, port, path := pr.c.ClientConfig()
	route := fmt.Sprintf("%s://%s:%s/%s", ws, host, port, path)

	// Create Dummy Data Message

	// Connect to the server
	// ws://host:port/ws example webSocket route
	conn, _, err := websocket.DefaultDialer.Dial(route, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Encode the message as JSON
	messageJSON, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	// Send the message to fleet_service server
	err = conn.WriteMessage(websocket.TextMessage, messageJSON)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a response from fleet_service server
	_, p, err := conn.ReadMessage()
	if err != nil {
		log.Fatal(err)
	}

	// Decode the response as JSON
	var response entity.Message
	err = json.Unmarshal(p, &response)
	if err != nil {
		log.Fatal(err)
	}

	// Log the response to the console
	log.Printf("Received message: %+v\n", response)

	// Close the connection
	conn.Close()
}
