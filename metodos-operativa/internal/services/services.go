package services

import (
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/data/responses"
)

type Service interface {
	Simplex(r requests.ProgramacionLinealRequest) (int, responses.SimplexResponse)
	DosFases(r requests.ProgramacionLinealRequest) (int, responses.DosFasesResponse)
}
