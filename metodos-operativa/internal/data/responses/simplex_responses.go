package responses

import (
	"metodos-operativa/internal/data/models"
)

type SimplexResponse struct {
	Resolucion []models.TablaSimplex `json:"resolucion"`
	Message    string                `json:"message"`
	Metodo     string                `json:"metodo"`
}

type DosFasesResponse struct {
	Resolucion DosFasesResolucion `json:"resolucion"`
	Message    string             `json:"message"`
	Metodo     string             `json:"metodo"`
}

type DosFasesResolucion struct {
	ResolucionFase1 []models.TablaSimplex `json:"resolucion_fase_1"`
	ResolucionFase2 []models.TablaSimplex `json:"resolucion_fase_2"`
}
