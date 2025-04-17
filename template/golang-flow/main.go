package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync/atomic"
	"syscall"
	"time"

	"handler/function"
)

var (
	acceptingConnections int32
)

const defaultTimeout = 10 * time.Second

func main() {
	readTimeout := parseIntOrDurationValue(os.Getenv("read_timeout"), defaultTimeout)
	writeTimeout := parseIntOrDurationValue(os.Getenv("write_timeout"), defaultTimeout)
	healthInterval := parseIntOrDurationValue(os.Getenv("healthcheck_interval"), writeTimeout)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", 8082),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: 1 << 20, // Max header of 1MB
	}

	http.HandleFunc("/", handle)

	listenUntilShutdown(s, healthInterval, writeTimeout)
}

// Handle is the main entrypoint for the function for Flow execution
func handle(w http.ResponseWriter, r *http.Request) {

	// Create a new FlowInput object
	var input function.FlowInput

	// Read the request body
	if r.Body != nil {
		defer r.Body.Close()

		body, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(body, &input)
		fmt.Printf("%+v\n", input)
	}

	// Execute the Flow
	outputData, err := function.ExecFlow(input)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}

	// Write the output data to the response writer and return it
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(outputData)
	return
}

func listenUntilShutdown(s *http.Server, shutdownTimeout time.Duration, writeTimeout time.Duration) {
	idleConnsClosed := make(chan struct{})
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGTERM)

		<-sig

		log.Printf("[entrypoint] SIGTERM: no connections in: %s", shutdownTimeout.String())
		<-time.Tick(shutdownTimeout)

		ctx, cancel := context.WithTimeout(context.Background(), writeTimeout)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			log.Printf("[entrypoint] Error in Shutdown: %v", err)
		}

		log.Printf("[entrypoint] Exiting.")

		close(idleConnsClosed)
	}()

	// Run the HTTP server in a separate go-routine.
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("[entrypoint] Error ListenAndServe: %v", err)
			close(idleConnsClosed)
		}
	}()

	atomic.StoreInt32(&acceptingConnections, 1)

	<-idleConnsClosed
}

func parseIntOrDurationValue(val string, fallback time.Duration) time.Duration {
	if len(val) > 0 {
		parsedVal, parseErr := strconv.Atoi(val)
		if parseErr == nil && parsedVal >= 0 {
			return time.Duration(parsedVal) * time.Second
		}
	}

	duration, durationErr := time.ParseDuration(val)
	if durationErr != nil {
		return fallback
	}
	return duration
}
