package controllers

import (
	"github.com/gin-gonic/gin"

	"e-commerce/internal/usecases/rest/services"
)

type Router struct {
	service *services.Service
}

func NewHandler(service *services.Service) *Router {
	return &Router{service: service}
}

// InitHTTPEndpoints will setting endpoints routing.
func (router *Router) InitHTTPEndpoints() *gin.Engine {
	engine := gin.New()

	auth := engine.Group("/auth")
	{
		auth.POST("/sign-up", router.SignUp)
		auth.POST("/sign-in", router.SignIn)
	}

	return engine
}
