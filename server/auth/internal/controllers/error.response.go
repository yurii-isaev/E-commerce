package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, ErrorResponse{message})
}
