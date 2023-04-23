package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Message struct {
	Text string `json:"text"`
}

type Client struct {
	conn *websocket.Conn
}

type Server struct {
	clients []*Client
	mu      sync.Mutex
}

func (s *Server) addClient(c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.clients = append(s.clients, c)
}

func (s *Server) removeClient(c *Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, client := range s.clients {
		if client == c {
			s.clients = append(s.clients[:i], s.clients[i+1:]...)
			return
		}
	}
}

func (s *Server) broadcastMessage(m *Message) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, client := range s.clients {
		err := client.conn.WriteJSON(m)
		if err != nil {
			log.Printf("error writing message to client: %v", err)
			continue
		}
	}
}

func (s *Server) handleConnection(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("error upgrading connection: %v", err)
		return
	}

	client := &Client{conn}
	s.addClient(client)
	defer s.removeClient(client)

	for {
		var m Message
		err := conn.ReadJSON(&m)
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}
		s.broadcastMessage(&m)
	}
}

func main() {
	server := &Server{}

	http.HandleFunc("/ws", server.handleConnection)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
