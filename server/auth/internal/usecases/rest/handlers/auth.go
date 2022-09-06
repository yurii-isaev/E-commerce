package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"e-commerce/internal/domains"
	response "e-commerce/internal/usecases/common"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input domains.User

	if err := c.BindJSON(&input); err != nil {
		response.NewError(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		response.NewError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
}
