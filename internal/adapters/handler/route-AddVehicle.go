// Package handler - module for the Gin frame work
package handler

import (
	"net/http"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"github.com/gin-gonic/gin"
)

// AddVehicle -
func (h *HTTPHandler) AddVehicle(ctx *gin.Context) {

	// Get post data
	var vehicle *domain.Vehicle
	if err := ctx.BindJSON(&vehicle); err != nil {
		h.logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err := h.svc.AddVehicle(vehicle)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
