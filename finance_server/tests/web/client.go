package web

import (
	"encoding/json"
	"finance_server/internal/domain/entity"
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type ClientProducer struct {
	//c config.Config
}

// Function to send a message over the WebSocket connection
func (pr ClientProducer) sendWSMessage(message entity.Message, wg *sync.WaitGroup) error {
	defer wg.Done()

	// ws, host, port, path := pr.c.ClientConfig()
	route := fmt.Sprintf("%s://%s:%s/%s", "ws", "localhost", "8080", "ws")

	// Connect to the server
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

	// Send the message to the server
	err = conn.WriteMessage(websocket.TextMessage, messageJSON)
	if err != nil {
		log.Fatal(err)
	}

	// Wait for a response from the server
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

	return nil
}

// Function to create a worker pool and send messages over the WebSocket connection
func (pr ClientProducer) SendMessages(numWorkers int, messages []entity.Message) error {
	// Create a wait group to wait for all workers to finish
	var wg sync.WaitGroup
	wg.Add(len(messages))

	// Create a channel to pass messages to the worker goroutines
	messageChan := make(chan entity.Message)

	// Start the worker goroutines
	for i := 0; i < numWorkers; i++ {
		go func() {
			for message := range messageChan {
				pr.sendWSMessage(message, &wg)
			}
		}()
	}

	// Pass messages to the worker goroutines
	for _, message := range messages {
		messageChan <- message

	}

	return nil

}
