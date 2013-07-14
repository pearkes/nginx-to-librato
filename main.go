package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	cpath := flag.String("config", "", "Path to a configuration file")
	flag.Parse()

	con, errs := NewConf(*cpath)
	// Print the errors from the config and exit if there are any
	if len(errs) > 0 {
		fmt.Fprintf(os.Stderr, "Configuration errors:\n")
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, "* %s\n", e.Error())
		}
		os.Exit(1)
	}

	flusher := NewMetricFlusher(con)
	// Start publishing metrics
	flusher.publishLoop()
}
