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

func NewProgramacionLinealController(service services.Service) *ProgramacionLinealController {
	return &ProgramacionLinealController{
		Service: service,
	}
}

func (ctrl *ProgramacionLinealController) ResolverProgramacionLineal(c *gin.Context) {
	var request requests.ProgramacionLinealRequest

	// Manejar error al parsear JSON
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Verificar que los datos sean válidos
	if len(request.FO) == 0 || len(request.Restricciones) == 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Verificar operadores en restricciones
	for _, restriccion := range request.Restricciones {
		operador := strings.TrimSpace(restriccion.Operador) // Quitar espacios
		if operador != "\u2264" {                           // Comparar con Unicode para '≤'
			status, res := ctrl.Service.DosFases(request)
			c.IndentedJSON(status, res)
			return
		}
	}

	// Si todas las restricciones usan '≤', usar Simplex
	status, res := ctrl.Service.Simplex(request)
	c.IndentedJSON(status, res)
}
