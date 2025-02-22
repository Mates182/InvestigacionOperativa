package responses

import (
	"metodos-operativa/internal/data/models"
)

// SimplexResponse representa la respuesta de la resolución de un problema de programación lineal
// utilizando el método Simplex, incluyendo la resolución paso a paso, el método utilizado,
// el modelo matemático y las respuestas obtenidas.
type SimplexResponse struct {
	Resolucion []models.TablaSimplex `json:"resolucion"` // Resolución del problema, paso a paso, utilizando la tabla simplex
	Message    string                `json:"message"`    // Mensaje que describe el resultado del proceso
	Metodo     string                `json:"metodo"`     // El método utilizado, en este caso "Simplex"
	Modelo     []string              `json:"modelo"`     // Representación del modelo de programación lineal
	Respuestas string                `json:"respuestas"` // Respuestas finales obtenidas de la solución
}

// DosFasesResponse representa la respuesta de la resolución de un problema de programación lineal
// utilizando el método de dos fases, con la resolución de cada fase del proceso, el método utilizado,
// el modelo y las respuestas obtenidas.
type DosFasesResponse struct {
	Resolucion DosFasesResolucion `json:"resolucion"` // Resolución de ambas fases del método de dos fases
	Message    string             `json:"message"`    // Mensaje que describe el resultado del proceso
	Metodo     string             `json:"metodo"`     // El método utilizado, en este caso "Dos Fases"
	Modelo     []string           `json:"modelo"`     // Representación del modelo de programación lineal
	Respuestas string             `json:"respuestas"` // Respuestas finales obtenidas de la solución
}

// DosFasesResolucion representa la resolución paso a paso de las dos fases del método de programación
// lineal de dos fases, mostrando las tablas simplex de cada fase.
type DosFasesResolucion struct {
	ResolucionFase1 []models.TablaSimplex `json:"resolucion_fase_1"` // Resolución de la fase 1 del método de dos fases
	ResolucionFase2 []models.TablaSimplex `json:"resolucion_fase_2"` // Resolución de la fase 2 del método de dos fases
}
