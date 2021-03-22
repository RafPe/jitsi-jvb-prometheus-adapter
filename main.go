package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	namespace = "jitsi"
	subsystem = "" // can be jvb
)

var (
	srv = &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d", 9001),
	}
	intervalStats time.Duration = 5 * time.Second         // Default 5 seconds time duration | set accordingly to JVB config
	interrupt                   = make(chan os.Signal, 1) // Handle the interrupts with GO routines

	promTags []PrometheusTag
)

func init() {
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	promTags = listPrometheusTags()
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	// Create a Resty Client
	client := resty.New()

	_, _ = client.R().
		Get("https://google.com")

	go func() {
		for {
			select {
			case s := <-interrupt:
				log.Printf("Signal (%d) received, stopping", s)
				time.Sleep(time.Duration(2 * time.Second)) // wait specific amount seconds to close all requests ...

				srv.Shutdown(context.Background())
			}
		}
	}()

	prometheus.NewDesc(
		prometheus.BuildFQName(namespace, "", "messages_filtered_total"),
		"How many messages have been filtered (per channel).",
		[]string{"channel"}, nil)

	exporter := NewExporter("whatever")
	prometheus.MustRegister(exporter)

	log.Fatal(srv.ListenAndServe(), nil)
}
