package main

import (
	"context"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	bucket := "example-bucket"
	org := "example-org"
	token := "example-token"
	// Store the URL of your InfluxDB instance
	url := "http://127.0.0.1"

	// Create new client with default option for server url authenticate by token
	client := influxdb2.NewClient(url, token)

	// User blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// Create point using full params constructor
	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now(),
	)

	// Write point immediately
	writeAPI.WritePoint(context.Background(), p)

	// Ensures background processes finishes
	client.Close()
}
