package controllers

import (
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type ProgramacionLinealController struct {
	Service services.Service
}

// Constructor del controlador de Programación Lineal
func NewProgramacionLinealController(service services.Service) *ProgramacionLinealController {
	return &ProgramacionLinealController{
		Service: service,
	}
}

func (ctrl *ProgramacionLinealController) ResolverProgramacionLineal(c *gin.Context) {
	var request requests.ProgramacionLinealRequest

	// Intentar parsear el JSON del cuerpo de la solicitud
	if err := c.BindJSON(&request); err != nil {
		// Responder con un estado 400 si los datos no son válidos
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Validar que se haya proporcionado una función objetivo y restricciones
	if len(request.FO) == 0 || len(request.Restricciones) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Verificar si alguna restricción usa un operador distinto de '≤'
	for _, restriccion := range request.Restricciones {
		operador := strings.TrimSpace(restriccion.Operador) // Eliminar espacios innecesarios
		if operador != "\u2264" {                           // Comparar con Unicode para '≤'
			// Si hay una restricción con otro operador, aplicar el método de Dos Fases
			status, res := ctrl.Service.DosFases(request)
			c.IndentedJSON(status, res)
			return
		}
	}

	// Si todas las restricciones usan '≤', aplicar el método Simplex
	status, res := ctrl.Service.Simplex(request)
	c.IndentedJSON(status, res)
}
