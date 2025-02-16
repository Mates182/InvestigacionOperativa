package router

import (
	"metodos-operativa/internal/controllers"
	"metodos-operativa/internal/data/messages"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/gemini"
	"metodos-operativa/internal/services"
	"net/http"

	"metodos-operativa/config/cors"

	"github.com/gin-gonic/gin"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.GetCORSConfig())
	r.POST("/simplex", controllers.NewProgramacionLinealController(services.NewServices()).ResolverProgramacionLineal)
	r.POST("/analisispl", GenerarAnalisis)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	return r
}

func GenerarAnalisis(c *gin.Context) {
	var request requests.PromptRequest

	// Manejar error al parsear JSON
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}
	prompt := messages.PromptProgramacionLineal() + request.Content

	res := gemini.GenerarTexto(prompt)
	// Enviar respuesta
	c.IndentedJSON(http.StatusOK, gin.H{"Message": res})
}
