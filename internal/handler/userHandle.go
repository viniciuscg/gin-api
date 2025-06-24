package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/viniciuscg/gin-api/internal/services"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (h *UserHandler) GetUser(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, users)
}
