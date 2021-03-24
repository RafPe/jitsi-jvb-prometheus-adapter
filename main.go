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

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const (
	namespace = "jitsi"
	subsystem = "" // can be jvb

	envJVBEndpoint = "JVB_STATS_URL"
)

var (
	srv = &http.Server{
		Addr: fmt.Sprintf(":%d", 9001),
	}
	intervalStats time.Duration = 5 * time.Second         // Default 5 seconds time duration | set accordingly to JVB config
	interrupt                   = make(chan os.Signal, 1) // Handle the interrupts with GO routines
	sugarLogger   *zap.SugaredLogger

	appVersion = "v0.0.3"
)

func init() {
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugarLogger = logger.Sugar()

	sugarLogger.Info("jitsi-jvb-prometheus-adapter is starting")

	sugarLogger.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"version", appVersion,
		"endpoint", os.Getenv(envJVBEndpoint),
		"address", srv.Addr,
	)

	sugarLogger.Infof("Failed to fetch URL: %s", "sss")

	http.Handle("/metrics", promhttp.Handler())

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

	exporter := NewJVBMetricsCollector(os.Getenv(envJVBEndpoint))
	prometheus.MustRegister(exporter)

	sugarLogger.Fatal(srv.ListenAndServe(), nil)
}
