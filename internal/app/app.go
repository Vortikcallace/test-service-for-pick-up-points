package app

import (
	"log"
	"test-service-for-pick-up-points/internal/config"
	"test-service-for-pick-up-points/internal/database"
	"test-service-for-pick-up-points/internal/transport/routes"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type App struct {
	config *config.Config
	db     *database.Database
	redis  *redis.Client
	router *gin.Engine
}

func NewApp() *App {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	v := config.SetUpViper()
	redisConfig := config.NewRedisConfig(v)
	redisClient := config.InitRedisClient(redisConfig)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	router := routes.SetupRoutes(db)

	return &App{
		config: cfg,
		db:     db,
		redis:  redisClient,
		router: router,
	}
}

func (a *App) Run() {
	log.Printf("Server is running on port %s", a.config.ServerPort)
	a.router.Run(":" + a.config.ServerPort)
}
