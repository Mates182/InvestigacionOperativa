package responses

import (
	"metodos-operativa/internal/data/models"
)

// RutaMasCortaResponse representa la respuesta con la ruta más corta en un grafo,
// incluyendo la distancia mínima y el grafo con la ruta calculada.
type RutaMasCortaResponse struct {
	DistanciaMinima float64      `json:"distanciaMinima"` // Distancia total de la ruta más corta
	GrafoDistancia  models.Grafo `json:"grafo"`           // Grafo que contiene las rutas y distancias
	Mensaje         string       `json:"mensaje"`         // Mensaje indicando el resultado de la operación
}

// FlujoMaximoResponse representa la respuesta con el flujo máximo en un grafo,
// incluyendo el valor del flujo y el grafo con el flujo calculado.
type FlujoMaximoResponse struct {
	Flujo      float64      `json:"flujo"`   // Valor del flujo máximo encontrado
	GrafoFlujo models.Grafo `json:"grafo"`   // Grafo que contiene el flujo calculado
	Mensaje    string       `json:"mensaje"` // Mensaje indicando el resultado de la operación
}
