package monitoring

import (
	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func PrometheusRoute(a *fiber.App) {

	a.Get("metrics", adaptor.HTTPHandler(promhttp.Handler()))
}
