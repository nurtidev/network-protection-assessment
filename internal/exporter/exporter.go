package exporter

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	Requests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"path"},
	)

	PacketCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "packet_count_total",
			Help: "Total number of packets",
		},
		[]string{"protocol"},
	)

	DataVolume = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "data_volume_total",
			Help: "Total volume of data",
		},
		[]string{"protocol"},
	)

	PacketDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "packet_duration_seconds",
			Help:    "Histogram of packet processing durations",
			Buckets: prometheus.LinearBuckets(0.001, 0.001, 10), // Buckets from 0.001s to 0.01s
		},
		[]string{"protocol"},
	)

	SuspiciousActivity = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "suspicious_activity_total",
			Help: "Total number of suspicious activities detected",
		},
		[]string{"src_ip"},
	)
)

func init() {
	prometheus.MustRegister(Requests)
	prometheus.MustRegister(PacketCount)
	prometheus.MustRegister(DataVolume)
	prometheus.MustRegister(PacketDuration)
	prometheus.MustRegister(SuspiciousActivity)
}

func RecordMetrics() {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		http.ListenAndServe(":8080", nil)
	}()
}
