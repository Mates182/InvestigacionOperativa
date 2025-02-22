package services

import (
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/data/responses"
)

// Interface Service que define los métodos disponibles para los servicios de programación operativa
type Service interface {
	// Método para resolver el problema de programación lineal usando el método Simplex
	Simplex(r requests.ProgramacionLinealRequest) (int, responses.SimplexResponse)

	// Método para resolver el problema de programación lineal utilizando el método de Dos Fases
	DosFases(r requests.ProgramacionLinealRequest) (int, responses.DosFasesResponse)

	// Método para resolver el problema de transporte usando la técnica de transporte
	Transporte(r requests.TransporteRequest) (int, responses.TransporteResponse)
}
