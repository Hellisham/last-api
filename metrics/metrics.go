package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
)

var (
	REQUUEST_COUNT = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "api_request_count",
		Help: "Total API HTTP request count."})
)

func InitMetrics() {
	err := prometheus.Register(REQUUEST_COUNT)
	if err != nil {
		log.Printf("Failed to register API request count: %s", err)
	}
}
