// Package handler - module for the Gin frame work
package handler

import (
	"errors"
	"net/http"
	"time"

	"github.com/JOHNEPPILLAR/nwg-de/internal/core/domain"
	"github.com/buger/jsonparser"
	"github.com/gin-gonic/gin"
)

// AddBooking -
func (h *HTTPHandler) AddBooking(ctx *gin.Context) {

	// Get post data
	rawBooking, err := ctx.GetRawData()
	if err != nil {
		h.logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	strStartDate, err := jsonparser.GetString(rawBooking, "StartDate")
	if err != nil {
		h.logger.Error(err.Error())
		userError := errors.New("Invalid start date")
		ctx.JSON(http.StatusBadRequest, userError.Error())
	}
	startDate, err := time.Parse("02-01-2006", strStartDate)
	if err != nil {
		h.logger.Error(err.Error())
		userError := errors.New("Invalid start date")
		ctx.JSON(http.StatusBadRequest, userError.Error())
	}

	strEndDate, err := jsonparser.GetString(rawBooking, "EndDate")
	if err != nil {
		h.logger.Error(err.Error())
		userError := errors.New("Invalid end date")
		ctx.JSON(http.StatusBadRequest, userError.Error())
	}
	endDate, err := time.Parse("02-01-2006", strEndDate)
	if err != nil {
		h.logger.Error(err.Error())
		userError := errors.New("Invalid end date")
		ctx.JSON(http.StatusBadRequest, userError.Error())
	}
	endDate = endDate.AddDate(0, 0, 1) // Make sure booking ends at midnight of that end date

	licenseNumber, err := jsonparser.GetString(rawBooking, "LicenseNumber")
	if err != nil {
		h.logger.Error(err.Error())
		userError := errors.New("Invalid License Number")
		ctx.JSON(http.StatusBadRequest, userError.Error())
	}

	var booking domain.Booking
	booking.LicenseNumber = licenseNumber
	booking.StartDate = startDate.UTC()
	booking.EndDate = endDate.UTC()
	booking.Cost = 0

	err = h.svc.AddBooking(&booking)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
