package main

import (
	"file_server/internal/infrastructure/webserver"
)

func main() {

	s := webserver.HTTPServer{}
	s.HTTPFileServer()
}
