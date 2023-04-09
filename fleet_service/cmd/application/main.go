package main

import (
	"platform_science_code_exercise/internal/infrastructure/websocket"
)

func main() {

	ws := websocket.WSConfigServer{}
	ws.WSFileServer()
}
