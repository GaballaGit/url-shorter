package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const YELLOW = "\x1b[33m"
const END_COLOR = "\x1b[0m"
const DEV_PORT = "8080"

func main() {
	fmt.Println(YELLOW + "[ STARTING SERVER ]" + END_COLOR)

	// logging
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})

	// TODO: Middleware for mux

	port := os.Getenv("PORT")
	if port == "" {
		port = DEV_PORT
	}

	// the server struct has timeouts, pretty cool. Maybe explore more fields later, but good against slowloris attacks for now :)
	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  2 * time.Minute,
	}

	// graceful shutdown yaya
	graceError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		s := <-quit

		logger.Info("shutting down server" + "signal:" + s.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30+time.Second)
		defer cancel()

		// shutdown is like a gentle Close() that first closes for new connections and then waits for the current ones to stop.
		// I chose to wait 30 seconds with a ctx
		graceError <- server.Shutdown(ctx)
	}()

	// jarvis, start the server
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		logger.Error("server failed to start:" + err.Error())
		os.Exit(1)
	}

	if err := <-graceError; err != nil {
		logger.Error("graceful error was not very graceful:" + err.Error())
		os.Exit(1)
	}

	fmt.Println(YELLOW + "[ SERVER STOPPED ]" + END_COLOR)
}
