package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

var statusMap = map[int]string{
	0: "active_connections",
	1: "accepts",
	2: "handled",
	3: "requests",
	4: "reading",
	5: "writing",
	6: "waiting",
}

// Type metric describes a metric
type metric struct {
	name  string
	value int64
}

// retrieves the status page, converts the metrics
// and returns an array of type metric
func getMetrics(c conf) []metric {
	attempts := 0
	for {
		if attempts == 500 {
			fmt.Fprintf(os.Stderr, "failed 500 attempts to check status at %s, giving up\n", c.url)
			os.Exit(1)
		}

		attempts += 1

		body, err := retrieveMetrics(c)

		if err != nil {
			log.Printf("Encountered error while requesting stats: %s. Sleeping and trying again (attempt %v/500)", err.Error(), attempts)
		} else {
			return convertMetrics(body)
		}

		// Wait 1 second in between
		time.Sleep(1 * time.Second)
	}
}

// convertMetrics converts a byte array representing the nginx
// status page into an array of type metric, suitable for our use.
func convertMetrics(body []byte) []metric {
	metrics := make([]metric, 0)

	re := regexp.MustCompile(`\d+`)
	matches := re.FindAll(body, -1)
	for i, m := range matches {
		val, _ := strconv.ParseInt(string(m), 0, 64)
		metrics = append(metrics, metric{statusMap[i], val})
	}

	return metrics
}

// Retrieves the status page via http
func retrieveMetrics(c conf) ([]byte, error) {
	url := fmt.Sprintf("http://%s", c.url)

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
