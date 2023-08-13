// Package handler - module for the Gin frame work
package handler

import (
	"net/http"
	"strings"
	"time"

	ratelimit "github.com/JGLTechnologies/gin-rate-limit"
	"github.com/JOHNEPPILLAR/nwg-de/internal/utility"
	"github.com/JOHNEPPILLAR/utility/vault"
	"github.com/gin-gonic/gin"
)

func keyFunc(c *gin.Context) string {
	return c.ClientIP()
}

func errorHandler(c *gin.Context, info ratelimit.Info) {
	c.String(429, "Too many requests.")
}

func (h *HTTPHandler) middleware() {
	store := ratelimit.InMemoryStore(&ratelimit.InMemoryOptions{
		Rate:  time.Second,
		Limit: 5})

	rl := ratelimit.RateLimiter(store, &ratelimit.Options{
		ErrorHandler: errorHandler,
		KeyFunc:      keyFunc,
	})

	h.router.SetTrustedProxies(nil)  // Trust no proxies
	h.router.Use(gin.Recovery())     // Add recovery from panics
	h.router.Use(addSecureHeaders()) // Add security headers
	h.router.Use(auth(h.logger))     // Check for valid api auth token
	h.router.Use(rl)                 // Rate limit
}

func auth(logger *utility.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		path := ctx.FullPath()
		if path == "/v1/health" {
			return
		}

		authorization := ctx.Request.Header["Authorization"]
		if authorization == nil {
			ctx.AbortWithStatus(http.StatusBadRequest)
			logger.Error("Missing auth token")
			return
		}

		idToken := strings.TrimSpace(strings.Replace(authorization[0], "Bearer", "", 1))

		if idToken == "" {
			ctx.AbortWithStatus(http.StatusBadRequest)
			logger.Error("Missing auth token")
			return
		}

		validAuthToken, err := vault.GetVaultSecret("CLIENT_ACCESS_KEY")
		if err != nil {
			ctx.AbortWithStatus(http.StatusInternalServerError)
			logger.Error("Unable to get token from vault")
			return
		}

		if validAuthToken != idToken {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			logger.Error("Invalid api auth token")
			return
		}
	}
}

func addSecureHeaders() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		headers := ctx.Writer.Header()

		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Referrer-Policy
		headers.Add("referrer-policy", "strict-origin-when-cross-origin")

		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Content-Type-Options
		headers.Add("x-content-type-options", "nosniff")

		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-Frame-Options
		headers.Add("x-frame-options", "DENY")

		// https://security.stackexchange.com/questions/166024/does-the-x-permitted-cross-domain-policies-header-have-any-benefit-for-my-websit
		headers.Add("X-Permitted-Cross-Domain-Policies", "none")

		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/X-XSS-Protection
		headers.Add("x-xss-protection", "1; mode=block")

		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Expect-CT
		headers.Add("Expect-CT", "max-age=0, enforce, report-uri=\"https://report-uri.com/r/d/ct/enforce\"")

		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Feature-Policy
		// https://github.com/w3c/webappsec-feature-policy/blob/master/features.md
		// https://developers.google.com/web/updates/2018/06/feature-policy
		headers.Add("Feature-Policy",
			"accelerometer 'none';"+
				"ambient-light-sensor 'none';"+
				"autoplay 'none';"+
				"battery 'none';"+
				"camera 'none';"+
				"display-capture 'none';"+
				"document-domain 'none';"+
				"encrypted-media 'none';"+
				"execution-while-not-rendered 'none';"+
				"execution-while-out-of-viewport 'none';"+
				"gyroscope 'none';"+
				"magnetometer 'none';"+
				"microphone 'none';"+
				"midi 'none';"+
				"navigation-override 'none';"+
				"payment 'none';"+
				"picture-in-picture 'none';"+
				"publickey-credentials-get 'none';"+
				"sync-xhr 'none';"+
				"usb 'none';"+
				"wake-lock 'none';"+
				"xr-spatial-tracking 'none';",
		)

		// https://developer.mozilla.org/en-US/docs/Web/HTTP/CSP
		// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Security-Policy
		headers.Add("Content-Security-Policy",
			"base-uri 'none';"+
				"block-all-mixed-content;"+
				"child-src 'none';"+
				"connect-src 'none';"+
				"default-src 'none';"+
				"font-src 'none';"+
				"form-action 'none';"+
				"frame-ancestors 'none';"+
				"frame-src 'none';"+
				"img-src 'none';"+
				"manifest-src 'none';"+
				"media-src 'none';"+
				"object-src 'none';"+
				"sandbox;"+
				"script-src 'none';"+
				"script-src-attr 'none';"+
				"script-src-elem 'none';"+
				"style-src 'none';"+
				"style-src-attr 'none';"+
				"style-src-elem 'none';"+
				"upgrade-insecure-requests;"+
				"worker-src 'none';",
		)
	}
}
