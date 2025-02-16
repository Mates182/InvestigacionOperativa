package requests

import "metodos-operativa/internal/data/models"

type ProgramacionLinealRequest struct {
	FO            []models.Termino `json:"fo"`
	Restricciones []struct {
		LI       []models.Termino `json:"li"`
		Operador string           `json:"operador"`
		LD       float64          `json:"ld"`
	} `json:"restricciones"`
	Maximizar bool `json:"maximizar"`
}
