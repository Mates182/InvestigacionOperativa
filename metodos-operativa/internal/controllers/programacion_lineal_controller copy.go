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

func NewTransporteController(service services.Service) *TransporteController {
	return &TransporteController{
		Service: service,
	}
}

func (ctrl *TransporteController) ResolverTransporte(c *gin.Context) {
	var request requests.TransporteRequest

	// Manejar error al parsear JSON
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"Message": "Error: No se han proporcionado datos válidos"})
		return
	}

	status, res := ctrl.Service.Transporte(request)
	c.IndentedJSON(status, res)
}
