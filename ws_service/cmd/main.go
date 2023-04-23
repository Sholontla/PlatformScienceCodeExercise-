package main

import (
	"log"
	"net/http"
	"ws_service/internal/infrastructure/websocket"
)

func main() {
	http.HandleFunc("/api", websocket.HandleWebSocket)
	log.Println("WS Server started on port 1006")
	log.Fatal(http.ListenAndServe(":1006", nil))
}
