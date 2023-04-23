package main

import (
	"finance_server/internal/infrastructure/webserver"
)

func main() {

	s := webserver.HTTPServer{}
	s.HTTPFileServer()

}

// func loadDatabase() {
// 	postgresql.Connect()
// 	postgresql.Database.AutoMigrate(&entity.Sale{})
// }
