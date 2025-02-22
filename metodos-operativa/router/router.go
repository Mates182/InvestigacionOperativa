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

// SetRouter configura las rutas del servidor web y los controladores asociados.
func SetRouter() *gin.Engine {
	r := gin.Default()          // Inicializa el router de Gin con los valores predeterminados.
	r.Use(cors.GetCORSConfig()) // Configura las políticas de CORS para permitir solicitudes desde orígenes específicos.

	// Define las rutas POST para los diferentes métodos y asigna los controladores correspondientes.
	r.POST("/simplex", controllers.NewProgramacionLinealController(services.NewServices()).ResolverProgramacionLineal)
	r.POST("/transporte", controllers.NewTransporteController(services.NewServices()).ResolverTransporte)
	r.POST("/grafos", controllers.NewGrafosController(services.NewServices()).ResolverGrafo)
	r.POST("/analisispl", GenerarAnalisisPL)                 // Ruta para generar análisis de programación lineal.
	r.POST("/analisistransporte", GenerarAnalisisTransporte) // Ruta para generar análisis de transporte.
	r.POST("/analisisgrafos", GenerarAnalisisGrafos)         // Ruta para generar análisis de grafos.

	// Define la ruta GET para verificar si el servidor está funcionando (ping).
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong", // Respuesta sencilla para verificar la conexión.
		})
	})
	return r // Retorna el router configurado.
}

// GenerarAnalisisPL maneja la solicitud para generar un análisis de programación lineal.
func GenerarAnalisisPL(c *gin.Context) {
	var request requests.PromptRequest

	// Manejar error al parsear JSON de la solicitud
	if err := c.BindJSON(&request); err != nil {
		// Si el cuerpo de la solicitud no contiene datos válidos, se responde con un error.
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Generar el prompt para el análisis de programación lineal usando el contenido de la solicitud.
	prompt := messages.PromptProgramacionLineal() + request.Content

	// Llamar a Gemini para generar el texto basado en el prompt.
	res := gemini.GenerarTexto(prompt)

	// Enviar la respuesta con el resultado generado por Gemini.
	c.IndentedJSON(http.StatusOK, gin.H{"Message": res})
}

// GenerarAnalisisTransporte maneja la solicitud para generar un análisis de transporte.
func GenerarAnalisisTransporte(c *gin.Context) {
	var request requests.PromptRequest

	// Manejar error al parsear JSON de la solicitud
	if err := c.BindJSON(&request); err != nil {
		// Si el cuerpo de la solicitud no contiene datos válidos, se responde con un error.
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Generar el prompt para el análisis de transporte usando el contenido de la solicitud.
	prompt := messages.PromptTransporte() + request.Content

	// Llamar a Gemini para generar el texto basado en el prompt.
	res := gemini.GenerarTexto(prompt)

	// Enviar la respuesta con el resultado generado por Gemini.
	c.IndentedJSON(http.StatusOK, gin.H{"Message": res})
}

// GenerarAnalisisGrafos maneja la solicitud para generar un análisis de grafos.
func GenerarAnalisisGrafos(c *gin.Context) {
	var request requests.PromptRequest

	// Manejar error al parsear JSON de la solicitud
	if err := c.BindJSON(&request); err != nil {
		// Si el cuerpo de la solicitud no contiene datos válidos, se responde con un error.
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Generar el prompt para el análisis de grafos usando el contenido de la solicitud.
	prompt := messages.PromptGrafos() + request.Content

	// Llamar a Gemini para generar el texto basado en el prompt.
	res := gemini.GenerarTexto(prompt)

	// Enviar la respuesta con el resultado generado por Gemini.
	c.IndentedJSON(http.StatusOK, gin.H{"Message": res})
}
