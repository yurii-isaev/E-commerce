package common

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func NewError(c *gin.Context, statusCode int, message string) {
	log.Fatal(message)
	c.AbortWithStatusJSON(statusCode, Error{message})
}
