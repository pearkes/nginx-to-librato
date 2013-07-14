package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	cpath := flag.String("config", "/etc/nginx-to-librato.conf", "Path to a configuration file")
	debug := flag.Bool("debug", false, "Turn on debugging")
	version := flag.Bool("version", false, "Prints the version and exits")

	flag.Parse()

	// Discard logging if debug is turned off.
	if *debug == false {
		log.SetOutput(ioutil.Discard)
	} else {
		log.Printf("Debugging enabled for nginx-to-librato %s", versionString())
	}

	// Print the version and exit
	if *version == true {
		fmt.Println(versionString())
		os.Exit(0)
	}

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
