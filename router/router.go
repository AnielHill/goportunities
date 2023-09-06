package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Initialize router default gin
	router := gin.Default()
	//Defined route
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Estamos rodando nossa api
	router.Run() // listen and serve on 0.0.0.0:8080
}
