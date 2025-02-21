package controllers

import (
	"fmt"
	"metodos-operativa/internal/data/models"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/data/responses"
	"metodos-operativa/internal/services"
	"metodos-operativa/pkg/grafos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GrafosController struct {
	Service services.Service
}

func NewGrafosController(service services.Service) *GrafosController {
	return &GrafosController{
		Service: service,
	}
}

func (ctrl *GrafosController) ResolverGrafo(c *gin.Context) {
	grafo := models.NuevoGrafo()
	var request requests.GrafosRequest

	// Intentar vincular el JSON del cuerpo de la solicitud a la estructura Grafo
	if err := c.ShouldBindJSON(&request); err != nil {
		// Si hay un error, responder con un estado 400 y el mensaje de error
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, conexion := range request.Conexiones {
		grafo.AgregarConexion(conexion.Origen, conexion.Destino, conexion.Costo, conexion.Capacidad, conexion.Distancia)
	}

	if request.EsRutaCorta {
		// Calcular la ruta más corta optimizando por Distancia
		distanciaTotal, rutaDistancia := grafos.DijkstraGrafo(grafo, request.Origen, request.Destino, false)
		fmt.Printf("Ruta más corta basada en Distancia: %f\n", distanciaTotal)
		//rutaDistancia.Mostrar()
		c.IndentedJSON(0, &responses.RutaMasCortaResponse{DistanciaMinima: distanciaTotal, GrafoDistancia: *rutaDistancia, Mensaje: "Solución óptima encontrada"})
	} else {
		// Aplicar Ford-Fulkerson
		flujoMaximo, flujoGrafo := grafos.FordFulkersonGrafo(grafo, request.Origen, request.Destino)
		fmt.Printf("Flujo Máximo: %.2f\n", flujoMaximo)

		// Mostrar la red de flujo obtenida
		//flujoGrafo.Mostrar()
		c.IndentedJSON(0, &responses.FlujoMaximoResponse{Flujo: flujoMaximo, GrafoFlujo: *flujoGrafo, Mensaje: "Solución óptima encontrada"})
	}
	//status, res := ctrl.Service.Transporte(request)
}
