// Package handler - module for the Gin frame work
package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (h *HTTPHandler) routes() {

	// Set API ver
	v1 := h.router.Group("/v1")

	// Add Prometheus metrics
	v1.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Add health check
	v1.GET("/health", h.HealthCheck)

	// Vehicles api group
	vehicle := v1.Group("/vehicle")
	vehicle.POST("", h.AddVehicle)
	vehicle.GET(":licenseNumber", h.FindVehicle)
	vehicle.GET("/all", h.GetAllVehicles)
	vehicle.GET("/available/:startDate/:endDate", h.GetAvailableVehicles)

	// Booking api group
	booking := v1.Group("/booking")
	booking.POST("", h.AddBooking)
}
