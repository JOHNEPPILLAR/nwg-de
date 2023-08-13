// Package handler - module for the Gin frame work
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck -
func (h *HTTPHandler) HealthCheck(ctx *gin.Context) {

	ping := h.svc.HealthCheck()

	if ping != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.AbortWithStatus(http.StatusOK)
}
