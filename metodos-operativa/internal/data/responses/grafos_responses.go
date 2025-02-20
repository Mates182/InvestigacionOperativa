package responses

import (
	"metodos-operativa/internal/data/models"
)

type GrafosResponse struct {
	Grafo           models.Grafo `json:"grafo"`
	Flujo           float64      `json:"flujo"`
	DistanciaMinima float64      `json:"distanciaMinima"`
	GrafoFlujo      models.Grafo `json:"grafoFlujo"`
	GrafoDistancia  models.Grafo `json:"grafoDistancia"`
}
