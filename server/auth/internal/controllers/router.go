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
func (r *Router) InitHTTPEndpoints() *gin.Engine {
	engine := gin.New()

	auth := engine.Group("/auth", r.userIdentity)
	{
		auth.POST("/sign-up", r.SignUp)
		auth.POST("/sign-in", r.SignIn)
	}

	return engine
}
