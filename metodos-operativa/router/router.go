package router

import (
	"metodos-operativa/internal/controllers"
	"metodos-operativa/internal/services"

	"metodos-operativa/config/cors"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.GetCORSConfig())
	r.POST("/simplex", controllers.NewProgramacionLinealController(services.NewServices()).ResolverProgramacionLineal)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
