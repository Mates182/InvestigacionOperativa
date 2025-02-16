package responses

import (
	"metodos-operativa/internal/data/models"
)

type SimplexResponse struct {
	Resolucion []models.TablaSimplex `json:"resolucion"`
	Message    string                `json:"message"`
	Metodo     string                `json:"metodo"`
	Modelo     []string              `json:"modelo"`
	Algebraico []string              `json:"algebraico"`
}

type DosFasesResponse struct {
	Resolucion DosFasesResolucion `json:"resolucion"`
	Message    string             `json:"message"`
	Metodo     string             `json:"metodo"`
	Modelo     []string           `json:"modelo"`
	Algebraico []string           `json:"algebraico"`
}

type DosFasesResolucion struct {
	ResolucionFase1 []models.TablaSimplex `json:"resolucion_fase_1"`
	ResolucionFase2 []models.TablaSimplex `json:"resolucion_fase_2"`
}
