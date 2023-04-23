package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/prometheus/client_golang/prometheus"
)

var latency = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Namespace:  "api",
		Name:       "latency_seconds",
		Help:       "Latency distributions.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	},
	[]string{"method", "path"},
)

func RecordRequestLatency(ctx *fiber.Ctx) error {
	var (
		start   = time.Now()
		next    = ctx.Next()
		elapsed = time.Since(start).Seconds()
	)
	latency.WithLabelValues(
		ctx.Route().Method,
		ctx.Route().Path,
	).Observe(elapsed)

	return next

}

func RegisterPrometheusMetrics() {
	prometheus.MustRegister(latency)
}
