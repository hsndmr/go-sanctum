package routers

import (
	"github.com/gin-gonic/gin"
	controllerv1 "github.com/hsndmr/go-sanctum/controllers/v1"
	docs "github.com/hsndmr/go-sanctum/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	initV1(r)
	return r
}

// initV1 initializes the v1 routes
func initV1(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	apiv1 := r.Group("/api/v1")
	// users 
	controllerv1.RegisterUserRoutes(apiv1)

	//swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}