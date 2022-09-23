package router

import (
	"go-merchant/controller"
	"go-merchant/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/ping", controller.Ping)
		api.POST("/signin", controller.SignIn)
		secured := api.Group("/secured").Use(middleware.Auth())
		{
			secured.GET("/pong", controller.Ping)
		}
		transaction := api.Group("/transaction").Use(middleware.Auth())
		{
			transaction.GET("/merchant", controller.GetMerchantTransaction)
			transaction.GET("/outlet", controller.GetOutletTransaction)
		}
	}
	return router
}
