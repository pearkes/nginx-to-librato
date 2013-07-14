package main

import (
	"reflect"
	"testing"
)

const examplePage = `Active connections: 291
server accepts handled requests
  16630948 16630948 31070465
Reading: 6 Writing: 179 Waiting: 106`

func TestCollector_convertMetrics(t *testing.T) {
	body := []byte(examplePage)
	metrics := convertMetrics(body)

	expected := []metric{
		metric{name: "active_connections", value: 291},
		metric{name: "accepts", value: 16630948},
		metric{name: "handled", value: 16630948},
		metric{name: "requests", value: 31070465},
		metric{name: "reading", value: 6},
		metric{name: "writing", value: 179},
		metric{name: "waiting", value: 106},
	}

	if reflect.DeepEqual(expected, metrics) != true {
		t.Fatalf("metrics not in expected format:\nexpected:\n%v\ngot:\n%v", expected, metrics)
	}
}

func TestCollector_statusMap(t *testing.T) {
	if len(statusMap) != 7 {
		t.Fatalf("incorrect number in status map: %v", len(statusMap))
	}
}
