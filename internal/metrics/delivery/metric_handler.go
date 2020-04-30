package delivery

import (
	"github.com/labstack/echo"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PromMetrics struct {
	Tracker *prometheus.CounterVec
	Timings *prometheus.HistogramVec
}

func RegisterMetrics(server *echo.Echo) PromMetrics {
	var metrics PromMetrics

	metrics.Tracker = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "url_tracker",
	}, []string{"status", "path", "method"})

	metrics.Timings = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "timings",
		},
		[]string{"status", "path", "method"},
	)

	prometheus.MustRegister(metrics.Tracker, metrics.Timings)

	server.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	return metrics
}
