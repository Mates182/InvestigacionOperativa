package controllers

import (
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransporteController struct {
	Service services.Service
}

// Constructor del controlador de Transporte
func NewTransporteController(service services.Service) *TransporteController {
	return &TransporteController{
		Service: service,
	}
}

func (ctrl *TransporteController) ResolverTransporte(c *gin.Context) {
	var request requests.TransporteRequest

	// Intentar parsear el JSON del cuerpo de la solicitud
	if err := c.BindJSON(&request); err != nil {
		// Responder con un estado 400 si los datos no son válidos
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	// Llamar al servicio de transporte para resolver el problema
	status, res := ctrl.Service.Transporte(request)
	c.IndentedJSON(status, res)
}
