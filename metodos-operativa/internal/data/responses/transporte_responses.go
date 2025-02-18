package responses

import "metodos-operativa/internal/data/requests"

type TransporteResponse struct {
	Request    requests.TransporteRequest `json:"request"`
	Message    string                     `json:"message"`
	Asignacion [][]float64                `json:"asignacion"`
	CostoTotal float64                    `json:"costoTotal"`
	Analisis   string                     `json:"analisis"`
}
