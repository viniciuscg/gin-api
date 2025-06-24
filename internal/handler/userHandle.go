package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandle struct {
	service *service.UserService
}

func UserHandle(service *service.UserService) *UserHandle {
	return &UserHandle{
		service: service,
	}
}

func (h *UserHandle) GetUser(c *gin.Context) {
	users, err := h.service.GetUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, users)
}
