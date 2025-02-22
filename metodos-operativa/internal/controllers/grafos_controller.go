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

// Constructor del controlador de grafos
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
		// Si hay un error en la solicitud, responder con un estado 400
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Agregar conexiones al grafo con los datos proporcionados en la solicitud
	for _, conexion := range request.Conexiones {
		grafo.AgregarConexion(conexion.Origen, conexion.Destino, conexion.Costo, conexion.Capacidad, conexion.Distancia)
	}

	if request.EsRutaCorta {
		// Si se solicita la ruta más corta, aplicar el algoritmo de Dijkstra
		distanciaTotal, rutaDistancia := grafos.DijkstraGrafo(grafo, request.Origen, request.Destino, false)
		fmt.Printf("Ruta más corta basada en Distancia: %f\n", distanciaTotal)

		// Retornar la respuesta con la distancia mínima y el grafo resultante
		c.IndentedJSON(0, &responses.RutaMasCortaResponse{DistanciaMinima: distanciaTotal, GrafoDistancia: *rutaDistancia, Mensaje: "Solución óptima encontrada"})
	} else {
		// Si no es ruta corta, aplicar el algoritmo de Ford-Fulkerson para flujo máximo
		flujoMaximo, flujoGrafo := grafos.FordFulkersonGrafo(grafo, request.Origen, request.Destino)
		fmt.Printf("Flujo Máximo: %.2f\n", flujoMaximo)

		// Retornar la respuesta con el flujo máximo y el grafo resultante
		c.IndentedJSON(0, &responses.FlujoMaximoResponse{Flujo: flujoMaximo, GrafoFlujo: *flujoGrafo, Mensaje: "Solución óptima encontrada"})
	}
}
