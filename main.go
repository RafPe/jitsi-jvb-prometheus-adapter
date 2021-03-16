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
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	srv = &http.Server{
		Addr: fmt.Sprintf("127.0.0.1:%d", 9001),
	}
	interrupt = make(chan os.Signal, 1) // Handle the interrupts with GO routines

)

func init() {
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	// Create a Resty Client
	client := resty.New()

	resp, err := client.R().
		EnableTrace().
		Get("https://google.com")

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	// Explore trace info
	fmt.Println("Request Trace Info:")
	ti := resp.Request.TraceInfo()
	fmt.Println("  DNSLookup     :", ti.DNSLookup)
	fmt.Println("  ConnTime      :", ti.ConnTime)
	fmt.Println("  TCPConnTime   :", ti.TCPConnTime)
	fmt.Println("  TLSHandshake  :", ti.TLSHandshake)
	fmt.Println("  ServerTime    :", ti.ServerTime)
	fmt.Println("  ResponseTime  :", ti.ResponseTime)
	fmt.Println("  TotalTime     :", ti.TotalTime)
	fmt.Println("  IsConnReused  :", ti.IsConnReused)
	fmt.Println("  IsConnWasIdle :", ti.IsConnWasIdle)
	fmt.Println("  ConnIdleTime  :", ti.ConnIdleTime)
	fmt.Println("  RequestAttempt:", ti.RequestAttempt)
	fmt.Println("  RemoteAddr    :", ti.RemoteAddr.String())

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

	log.Fatal(srv.ListenAndServe(), nil)
}
