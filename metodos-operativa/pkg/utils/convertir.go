package utils

import (
	"fmt"
	"metodos-operativa/internal/data/requests"
)

// ConvertirProgramacionLinealAFormato devuelve una representación legible del problema de Programación Lineal.
func ConvertirProgramacionLinealAFormato(req requests.ProgramacionLinealRequest) (string, string) {
	// Construir la función objetivo
	var fo string
	if req.Maximizar {
		fo = "Maximizar: Z = "
	} else {
		fo = "Minimizar: Z = "
	}

	for i, termino := range req.FO {
		if i > 0 {
			fo += fmt.Sprintf(" + %.2fX%d", termino.C, i+1)
		} else {
			fo += fmt.Sprintf("%.2fX%d", termino.C, i+1)
		}
	}

	// Construir las restricciones
	restricciones := ""
	for _, restriccion := range req.Restricciones {
		restriccionStr := ""
		for i, termino := range restriccion.LI {
			if i > 0 {
				restriccionStr += fmt.Sprintf(" + %.2fX%d", termino.C, i+1)
			} else {
				restriccionStr += fmt.Sprintf("%.2fX%d", termino.C, i+1)
			}
		}
		restriccionStr += fmt.Sprintf(" %s %.2f\n", restriccion.Operador, restriccion.LD)
		restricciones += restriccionStr
	}

	// Restricciones de no negatividad
	variables := "X1"
	for i := 1; i < len(req.FO); i++ {
		variables += fmt.Sprintf(", X%d", i+1)
	}
	variables += " ≥ 0"

	// Retornar la representación en formato texto
	return fo, fmt.Sprintf("%s%s", restricciones, variables)
}
