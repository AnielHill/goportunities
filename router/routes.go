package router

import (
	docs "github.com/AnielHill/goportunities/docs"
	"github.com/AnielHill/goportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		v1.GET("/opening", handler.ShowOpeningHangler)
		v1.POST("/opening", handler.CreateOpeningHangler)
		v1.DELETE("/opening", handler.DeleteOpeningHangler)
		v1.PUT("/opening", handler.UpdateOpeningHangler)
		v1.GET("/openings", handler.ListOpeningHangler)
	}
	//Initialize Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
