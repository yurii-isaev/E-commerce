package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (r *Router) userIdentity(c *gin.Context) {
	// Header check for empty -> 401.
	header := c.GetHeader(authHeader)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	// Header = bearer[0] + token[1] check -> 401.
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	// Operation call check -> 401.
	userId, err := r.service.Authentication.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	// Record value ID in context -> context.
	c.Set(userCtx, userId)
}
