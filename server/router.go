package server

import (
	"sample-go-app/controllers"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	gin.Default()
	// router.Use(gin.Logger())
	// router.Use(gin.Recovery())
	gin.SetMode(gin.ReleaseMode)

	health := new(controllers.HealthController)
	router.GET("/health", health.Status)

	v1 := router.Group("v1")
	{
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/", user.Retrieve)
			userGroup.GET("/:id", user.Retrieve)
			userGroup.POST("/", user.Create)
		}
	}
	return router

}
