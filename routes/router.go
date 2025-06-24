package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/viniciuscg/gin-api/internal/handler"
	"github.com/viniciuscg/gin-api/internal/repository"
	"github.com/viniciuscg/gin-api/internal/services"
)

func SetupRouter(dbConn *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(dbConn)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	user := router.Group("/users")
	{
		user.GET("/", userHandler.GetUser)
	}

	return router
}
