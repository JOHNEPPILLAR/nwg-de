// Package handler - module for the Gin frame work
package handler

import (
	"os"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/services"
	"github.com/JOHNEPPILLAR/nwg-de/internal/utility"

	"net/http"

	"github.com/gin-gonic/gin"
)

// HTTPHandler -
type HTTPHandler struct {
	svc    services.APIService
	router *gin.Engine
	logger *utility.Logger
}

// NewHandler -
func NewHandler(logger utility.Logger, APIService services.APIService) *HTTPHandler {
	return &HTTPHandler{
		svc:    APIService,
		logger: &logger,
	}
}

// SetupRouter -
func (h *HTTPHandler) SetupRouter() {

	h.logger.Info("Setting up routing and middleware...")

	h.router = gin.New()

	environment := os.Getenv("ENVIRONMENT")
	if environment == "Development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup middleware
	h.middleware()

	// Setup routes
	h.routes()
}

// Start -
func (h *HTTPHandler) Start() {

	h.logger.Info("Starting API server...")

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	err := h.router.Run(":" + httpPort)
	if err != nil && err != http.ErrServerClosed {
		h.logger.Fatal(err.Error())
		os.Exit(1)
	}

	h.logger.Info("Started api Server on port " + httpPort)
}
