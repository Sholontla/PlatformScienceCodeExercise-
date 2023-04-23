package webserver

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"publisher_service/config"
	"publisher_service/internal/interfaces"
	"publisher_service/pkg/monitoring"
	"publisher_service/pkg/monitoring/middleware"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HTTPPublisherService struct {
	conf   config.Config
	router interfaces.UsersRouterService
}

func (pub HTTPPublisherService) HTTPPublisherServer() {

	port := fmt.Sprintf(":%s", pub.conf.HTTPConfig())

	app := fiber.New()

	monitoring.PrometheusRoute(app)
	middleware.RegisterPrometheusMetrics()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	pub.router.PublisherRouter(app)

	go app.Listen(port)

	// SIGINT is the signal sent when we press Ctrl+C
	// SIGTERM gracefully kills the process
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit

	log.Println("Shutting down Publisher Demo server.....")
	if err := app.Shutdown(); err != nil {
		log.Fatalf("Shutting Down Publisher Demo Server: %v\n", err)
	}
}
