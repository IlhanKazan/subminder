package main

import (
	"log"
	"strings"
	"subminder/internal/domain"
	"subminder/internal/repository"
	"subminder/internal/service"
	"subminder/internal/transport/rest"
	"subminder/internal/worker"
	"subminder/pkg/database"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

// @title           SubMinder API
// @version         1.0
// @description     A RESTful API to track subscriptions and calculate renewal dates.

// @host      localhost:8080
// @BasePath  /api/v1
// @schemes   http
func main() {
	viper.SetConfigFile("configs/config.yaml")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading config file: ", err)
	}

	db := database.ConnectToDB()

	err := db.AutoMigrate(&domain.Category{}, &domain.Subscription{})
	if err != nil {
		log.Fatal("Migration error: ", err)
	}
	log.Println("Tables successfully migrated")
	database.SeedDatabase(db)

	subscriptionRepo := repository.NewSubscriptionRepository(db)
	subscriptionService := service.NewSubscriptionService(subscriptionRepo)
	subscriptionHandler := rest.NewSubscriptionHandler(subscriptionService)
	categoryRepo := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := rest.NewCategoryHandler(categoryService)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Subminder is healthy"})
	})

	rest.RegisterRoutes(r, subscriptionHandler, categoryHandler)

	myWorker := worker.NewRenewalWorker(subscriptionRepo)
	log.Println("Starting background worker")
	myWorker.Start()

	port := viper.GetString("server.port")
	r.Run(":" + port)
}
