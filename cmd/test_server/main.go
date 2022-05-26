package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/test_server/internal/domain/event"
	"github.com/test_server/internal/infra/http"
	"github.com/test_server/internal/infra/http/controllers"
)

// @title                       Test Server
// @version                     0.1.0
// @description                 Test Server boilerplate
func main() {
	exitCode := 0
	ctx, cancel := context.WithCancel(context.Background())

	// Recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("The system panicked!: %v\n", r)
			fmt.Printf("Stack trace form panic: %s\n", string(debug.Stack()))
			exitCode = 1
		}
		os.Exit(exitCode)
	}()

	// Signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-c
		fmt.Printf("Received signal '%s', stopping... \n", sig.String())
		cancel()
		fmt.Printf("Sent cancel to all threads...")
	}()

	// Event
	eventRepository := event.NewRepository()
	eventService := event.NewService(&eventRepository)
	eventController := controllers.NewEventController(&eventService)

	// HTTP Server
	err := http.Server(
		ctx,
		http.Router(
			eventController,
		),
	)

	if err != nil {
		fmt.Printf("http server error: %s", err)
		exitCode = 2
		return
	}
}
