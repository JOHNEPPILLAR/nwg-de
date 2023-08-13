// Package handler - module for the Gin frame work
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindVehicle -
func (h *HTTPHandler) FindVehicle(ctx *gin.Context) {

	licenseNumber := ctx.Param("licenseNumber")

	results, err := h.svc.FindVehicle(licenseNumber)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": results})
}
