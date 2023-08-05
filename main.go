package main

import (
	"context"
	"log"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	bucket := "sensor"
	org := "fandogh"
	token := "secret"
	url := "http://127.0.0.1:8086"

	client := influxdb2.NewClient(url, token)

	if _, err := client.Ping(context.Background()); err != nil {
		log.Printf("pinging influx failed %s", err)
	}

	writeAPI := client.WriteAPIBlocking(org, bucket)

	p := influxdb2.NewPoint("stat",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 24.5, "max": 45},
		time.Now(),
	)

	if err := writeAPI.WritePoint(context.Background(), p); err != nil {
		log.Printf("writing into influx failed %s", err)
	}

	client.Close()
}
