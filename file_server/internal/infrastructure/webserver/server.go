package webserver

import (
	"file_server/config"
	"file_server/internal/interfaces"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

type HTTPServer struct {
	c config.Config
}

func (c HTTPServer) HTTPFileServer() {

	port := fmt.Sprintf(":%s", c.c.HTTPConfig())

	app := fiber.New()

	interfaces.RouterDriverNames(app)

	go app.Listen(port)

	// SIGINT is the signal sent when we press Ctrl+C
	// SIGTERM gracefully kills the process
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down Service Demo server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Service Demo Server: %v\n", err)
	}
}
