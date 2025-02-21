package responses

import (
	"metodos-operativa/internal/data/models"
)

type RutaMasCortaResponse struct {
	DistanciaMinima float64      `json:"distanciaMinima"`
	GrafoDistancia  models.Grafo `json:"grafo"`
	Mensaje         string       `json:"mensaje"`
}
type FlujoMaximoResponse struct {
	Flujo      float64      `json:"flujo"`
	GrafoFlujo models.Grafo `json:"grafo"`
	Mensaje    string       `json:"mensaje"`
}
