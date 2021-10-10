// Copyright (c) 2021 teal.finance
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.

package main

import (
	"log"
	"net"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/armon/go-metrics"
	"github.com/armon/go-metrics/prometheus"
	"github.com/go-chi/chi/v5"
	"github.com/justinas/alice"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpConn     int64 // gauge
	httpActive   int64 // counter
	httpIdle     int64 // counter
	httpHijacked int64 // counter
)

// metricsServer creates and starts the Prometheus export server.
func metricsServer() (alice.Chain, func(net.Conn, http.ConnState)) {
	if exportPort == "" || exportPort == "no" {
		log.Print("Disable Prometheus because export port=", exportPort)
	} else if handler := metricsHandler(); handler != nil {
		addr := ":" + exportPort

		go func() {
			err := http.ListenAndServe(addr, handler)
			log.Fatal(err)
		}()

		log.Print("Enable Prometheus export listening on http://localhost", addr)

		middlewares := alice.New(countRED, logRequests, limitReqRate, setServerHeader)

		// connState counts the HTTP client connections.
		// In prod mode, we do not care about minor counting errors, we use the unsafe-thread version.
		// In dev mode we use the atomic version to avoid warnings from "go build -race".
		connState := updateConnCounters
		if *dev {
			connState = updateConnCountersAtomic
		}

		return middlewares, connState
	}

	log.Print("Disable Prometheus export")

	return alice.New(logRequests, limitReqRate, setServerHeader), nil
}

// metricsHandler returns the endpoints "/metrics/xxx".
func metricsHandler() http.Handler {
	sink, err := prometheus.NewPrometheusSink()
	if err != nil {
		log.Print("ERROR: NewPrometheusSink cannot register sink because ", err)
		return nil
	}

	if _, err := metrics.NewGlobal(metrics.DefaultConfig("Rainbow"), sink); err != nil {
		log.Print("ERROR: Prometheus export is not able to provide metrics because ", err)
		return nil
	}

	handler := chi.NewRouter()
	handler.Handle("/metrics", promhttp.Handler())

	return handler
}

// updateConnCounters increments/decrements the number of connections.
func updateConnCounters(nc net.Conn, cs http.ConnState) {
	switch cs {
	// StateNew: the client just connects to TealAPI expecting a request.
	// Transition to either StateActive or StateClosed.
	case http.StateNew:
		httpConn++
	// StateActive when 1 or more bytes of a request has been read.
	// After the request is handled, transitions to StateClosed, StateHijacked, or StateIdle.
	// HTTP/2: StateActive only transitions away once all active requests are complete.
	case http.StateActive:
		httpActive++
	// StateIdle when handling a request is finished and is in the keep-alive state,
	// waiting for a new request, then transitions to either StateActive or StateClosed.
	case http.StateIdle:
		httpIdle++
	// StateHijacked is a terminal state: does not transition to StateClosed.
	case http.StateHijacked:
		httpHijacked++
		httpConn--
	// StateClosed is a terminal state.
	case http.StateClosed:
	default:
		httpConn--
	}
}

func updateConnCountersAtomic(nc net.Conn, cs http.ConnState) {
	switch cs {
	case http.StateNew:
		atomic.AddInt64(&httpConn, 1)
	case http.StateActive:
		atomic.AddInt64(&httpActive, 1)
	case http.StateIdle:
		atomic.AddInt64(&httpIdle, 1)
	case http.StateHijacked:
		atomic.AddInt64(&httpHijacked, 1)
		atomic.AddInt64(&httpConn, -1)
	case http.StateClosed:
	default:
		atomic.AddInt64(&httpConn, -1)
	}
}

// countRED increments/decrements the RED metrics depending on incoming requests and outgoing responses.
func countRED(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		record := &statusRecorder{ResponseWriter: w, Status: "success"}

		next.ServeHTTP(record, r)

		labels := []metrics.Label{
			{Name: "method", Value: r.Method},
			{Name: "route", Value: r.RequestURI},
			{Name: "status", Value: record.Status},
		}

		duration := time.Since(start)
		metrics.AddSampleWithLabels([]string{"request_duration"}, float32(duration.Milliseconds()), labels)

		log.Printf("%v %v %v\n", r.Method, r.RequestURI, duration)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	Status string
}

func (r *statusRecorder) WriteHeader(status int) {
	if status != http.StatusOK {
		r.Status = "error"
	}

	r.ResponseWriter.WriteHeader(status)
}

// setServerHeader sets the Server HTTP header in the response.
func setServerHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Server", version)
			next.ServeHTTP(w, r)
		})
}

// logRequests logs the incoming HTTP requests.
func logRequests(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("Received: %+v %+v", r.Method, r.URL)
			next.ServeHTTP(w, r)
		})
}