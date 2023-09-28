package router

import (
	"github.com/AnielHill/goportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", handler.ShowOpeningHangler)
		v1.POST("/opening", handler.CreateOpeningHangler)
		v1.DELETE("/opening", handler.DeleteOpeningHangler)
		v1.PUT("/opening", handler.UpdateOpeningHangler)
		v1.GET("/openings", handler.ListOpeningHangler)
	}
}
