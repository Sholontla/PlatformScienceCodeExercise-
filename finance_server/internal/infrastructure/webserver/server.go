package webserver

import (
	"finance_server/config"
	"finance_server/internal/infrastructure/grpc"
	"finance_server/internal/infrastructure/kafka"
	"finance_server/internal/infrastructure/persistance/redis"

	"finance_server/internal/interfaces"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HTTPServer struct {
	c config.Config
}

func (c HTTPServer) HTTPFileServer() {

	port := fmt.Sprintf(":%s", c.c.HTTPConfig())

	app := fiber.New()

	redis.SetUpRedis()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	interfaces.RouterFinance(app)

	go func() {
		kafka.ConsumerSupplierRegistration()
	}()

	go func() {
		grpc.GrpcProductServer()
	}()

	// Start the HTTP server in the main goroutine
	if err := app.Listen(port); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

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
