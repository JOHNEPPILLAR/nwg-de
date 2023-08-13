// Package handler - module for the Gin frame work
package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// GetAvailableVehicles -
func (h *HTTPHandler) GetAvailableVehicles(ctx *gin.Context) {

	strStartDate := ctx.Param("startDate")
	startDate, err := time.Parse("02-01-2006", strStartDate)
	if err != nil {
		h.logger.Error(err.Error())
		userError := errors.New("Invalid start date")
		ctx.JSON(http.StatusBadRequest, userError.Error())
	}

	strEndDate := ctx.Param("endDate")
	endDate, err := time.Parse("02-01-2006", strEndDate)
	if err != nil {
		h.logger.Error(err.Error())
		userError := errors.New("Invalid end date")
		ctx.JSON(http.StatusBadRequest, userError.Error())
	}

	results, err := h.svc.GetAvailableVehicles(startDate, endDate)

	if err != nil {
		h.logger.Error(err.Error())
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": results})
}
