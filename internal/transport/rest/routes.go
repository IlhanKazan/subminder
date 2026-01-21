package rest

import (
	_ "subminder/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(r *gin.Engine, subscriptionHandler *SubscriptionHandler, categoryHandler *CategoryHandler) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/subscriptions", subscriptionHandler.CreateSubscription)
		v1.GET("/subscriptions", subscriptionHandler.GetAllSubscriptions)
		v1.POST("categories", categoryHandler.CreateCategory)
		v1.GET("/categories", categoryHandler.GetAllCategories)
		r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
