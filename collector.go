package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strconv"
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

func getMetrics(c *conf) []metric {
	body := retrieveMetrics(c)
	return convertMetrics(c, body)
}

// Retrieves the metrics and converts them into []metrics
func convertMetrics(c *conf, body []byte) []metric {
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
func retrieveMetrics(c *conf) []byte {
	url := fmt.Sprintf("http://%s", c.url)

	resp, err := http.Get(url)
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	return body
}
