package responses

import "metodos-operativa/internal/data/requests"

// TransporteResponse representa la respuesta de un problema de transporte resuelto,
// incluyendo la solicitud original, el mensaje con información adicional, la asignación
// de recursos a los orígenes y destinos, el costo total de la asignación y un análisis adicional.
type TransporteResponse struct {
	Request    requests.TransporteRequest `json:"request"`    // La solicitud original del problema de transporte
	Message    string                     `json:"message"`    // Mensaje que describe el resultado del proceso
	Asignacion [][]float64                `json:"asignacion"` // La matriz de asignación de recursos entre los orígenes y destinos
	CostoTotal float64                    `json:"costoTotal"` // El costo total de la asignación
	Analisis   string                     `json:"analisis"`   // Análisis adicional del resultado obtenido
}
