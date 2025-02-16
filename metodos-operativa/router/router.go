package router

import (
	"metodos-operativa/internal/controllers"
	"metodos-operativa/internal/services"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.POST("/simplex", controllers.NewProgramacionLinealController(services.NewServices()).ResolverProgramacionLineal)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}
