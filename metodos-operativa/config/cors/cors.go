package cors

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GetCORSConfig devuelve una configuración de CORS para la aplicación
func GetCORSConfig() gin.HandlerFunc {
	// Crear una nueva configuración de CORS
	corsConfig := cors.New(cors.Config{
		// Permitir todos los orígenes (remover si se desean orígenes específicos)
		AllowAllOrigins: true,

		// Definir los métodos HTTP permitidos
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},

		// Especificar los encabezados permitidos
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},

		// Permitir credenciales como cookies o encabezados de autorización
		AllowCredentials: true,
	})

	// Retornar la configuración de CORS
	return corsConfig
}
