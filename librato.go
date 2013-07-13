package main

import (
	"github.com/rcrowley/go-librato"
	"log"
	"time"
)

type metricFlusher struct {
	lib      librato.Metrics // The instance of librato we want
	interval time.Duration
	config   *conf
}

func NewMetricFlusher(c *conf) *metricFlusher {
	met := &metricFlusher{
		librato.NewSimpleMetrics(c.libUser, c.libToken, c.libSource),
		c.flushInterval,
		c,
	}
	return met
}

// Publishes metrics at the configured interval
func (m *metricFlusher) publishLoop() {
	for {
		<-time.After(m.interval)
		metrics := getMetrics(m.config)
		log.Println("Flushing to Librato...")
		for _, met := range metrics {
			sink := m.lib.GetGauge(met.name)
			sink <- met.value
		}
	}
}
