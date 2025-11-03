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
	orderRepo := database.NewOrderRepository(db.DB)
	productRepo := database.NewProductRepository(db.DB)
	pointRepo := database.NewPointRepository(db.DB)

	userService := services.NewUserService(userRepo)
	orderService := services.NewOrderService(orderRepo, userRepo, productRepo, pointRepo)
	productService := services.NewProductService(productRepo)
	pointService := services.NewPointService(pointRepo)

	userHandler := handlers.NewUserHandler(userService)
	orderHandler := handlers.NewOrderHandler(orderService)
	productHandler := handlers.NewProductHandler(productService)
	pointHandler := handlers.NewPointHandler(pointService)

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("/", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		orders := api.Group("/orders")
		{
			orders.POST("/", orderHandler.CreateOrder)
			orders.GET("/user/:user_id", orderHandler.GetUserOrders)
			orders.PATCH("/:id/readiness", orderHandler.UpdateOrderReadiness)
			orders.PATCH("/:id/access", orderHandler.UpdateOrderAccess)
			orders.PATCH("/:id/active", orderHandler.UpdateOrderActive)
		}

		products := api.Group("/products")
		{
			products.POST("/", productHandler.CreateProduct)
			products.GET("/", productHandler.GetProducts)
			products.GET("/:id", productHandler.GetProduct)
			products.PUT("/products/:id", productHandler.UpdateProduct)
		}

		points := api.Group("/points")
		{
			points.POST("/", pointHandler.CreatePoint)
			points.GET("/", pointHandler.GetPoints)
			points.GET("/:id", pointHandler.GetPoint)
		}
	}

	return router
}
