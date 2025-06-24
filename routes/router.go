package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/viniciuscg/gin-api/internal/handle"
	"github.com/viniciuscg/gin-api/internal/handler"
	"github.com/viniciuscg/gin-api/internal/services"
)

func SetupRouter(dbConn *sql.DB) *gin.Engine {
	router := gin.Default()

	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.UserHandle(userService)

	user := router.Group("/users")
	{
		user.GET("/", handle.HandleGetUsers())
	}

	return router
}
