package main

import "publisher_service/internal/infrastructure/webserver"

func main() {

	s := webserver.HTTPPublisherService{}
	s.HTTPPublisherServer()
}
