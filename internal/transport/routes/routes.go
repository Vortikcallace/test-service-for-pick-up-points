package routes

import (
	"test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/services"
	"test-service-for-pick-up-points/internal/transport/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *database.Database) *gin.Engine {
	router := gin.Default()

	userRepo := database.NewUserRepository(db.DB)

	userService := services.NewUserService(userRepo)

	userHandler := handlers.NewUserHandler(userService)

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
		}
	}

	return router
}
