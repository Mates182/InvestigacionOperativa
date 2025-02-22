package requests

import "metodos-operativa/internal/data/models"

// ProgramacionLinealRequest representa una solicitud para resolver un problema de programación lineal.
type ProgramacionLinealRequest struct {
	FO            []models.Termino `json:"fo"` // Función objetivo, representada como una lista de términos
	Restricciones []struct {
		LI       []models.Termino `json:"li"`       // Lado izquierdo de la restricción (expresión lineal)
		Operador string           `json:"operador"` // Operador de la restricción (≤, =, ≥)
		LD       float64          `json:"ld"`       // Lado derecho de la restricción
	} `json:"restricciones"`
	Maximizar bool `json:"maximizar"` // Indica si la función objetivo debe maximizarse (true) o minimizarse (false)
}
