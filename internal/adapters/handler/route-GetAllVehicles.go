// Package handler - module for the Gin frame work
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllVehicles -
func (h *HTTPHandler) GetAllVehicles(ctx *gin.Context) {

	results, err := h.svc.GetAllVehicles()

	if err != nil {
		h.logger.Error(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": results})
}
