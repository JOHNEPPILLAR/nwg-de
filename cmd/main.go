// Package main - Read data from a glucose sensor api
package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/JOHNEPPILLAR/nwg-de/internal/adapters/handler"
	"github.com/JOHNEPPILLAR/nwg-de/internal/adapters/repository"
	"github.com/JOHNEPPILLAR/nwg-de/internal/core/services"
	"github.com/JOHNEPPILLAR/nwg-de/internal/utility"
)

var (
	httpHandler *handler.HTTPHandler
	apiService  *services.APIService
)

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Setup logger
	logger, logErr := utility.NewLogger()
	if logErr != nil {
		logger.Fatal("Unable to create logger")
		os.Exit(1)
	}

	logger.Info("Starting Car Rental Service API Server...")

	// Setup data repository as mongo
	repository := repository.NewMongoAdaptor(ctx, *logger)

	// Setup http service
	apiService = services.NewAPIService(*logger, repository)

	// Setup http handler
	apiHandler := handler.NewHandler(*logger, *apiService)

	// Setup middleware and routes
	apiHandler.SetupRouter()

	// Start API service
	apiHandler.Start()

	// Listen for interrupt signal
	<-ctx.Done()

	// Restore default behavior on the interrupt signal
	stop()
	logger.Warn("Received interrupt signal, shutting down...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Disconnect database
	repository.Disconnect()

	logger.Info("Shutdown")
}
