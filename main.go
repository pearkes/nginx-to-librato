package main

import (
	"flag"
)

func main() {
	cpath := flag.String("config", "", "Path to a configuration file")
	flag.Parse()
	con := NewConf(*cpath)
	flusher := NewMetricFlusher(con)
	// Start publishing metrics
	flusher.publishLoop()
}
