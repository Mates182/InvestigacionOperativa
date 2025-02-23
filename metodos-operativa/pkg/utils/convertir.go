package utils

import (
	"fmt"
	"math"
	"metodos-operativa/internal/data/models"
	"metodos-operativa/internal/data/requests"
)

// ConvertirProgramacionLinealAFormato devuelve una representación legible del problema de Programación Lineal.
func ConvertirProgramacionLinealAFormato(req requests.ProgramacionLinealRequest) (string, string) {
	// Construir la función objetivo
	var fo string
	// Determinar si se está maximizando o minimizando la función objetivo
	if req.Maximizar {
		fo = "Maximizar: Z = "
	} else {
		fo = "Minimizar: Z = "
	}

	// Agregar los términos de la función objetivo
	for i, termino := range req.FO {
		if i > 0 {
			fo += fmt.Sprintf(" + %.2fX%d", termino.C, i+1) // Si no es el primer término, agregar el signo de adición
		} else {
			fo += fmt.Sprintf("%.2fX%d", termino.C, i+1) // Para el primer término no agregar el signo de adición
		}
	}

	// Construir las restricciones
	restricciones := ""
	for _, restriccion := range req.Restricciones {
		restriccionStr := ""
		// Agregar los términos de cada restricción
		for i, termino := range restriccion.LI {
			tempC := math.Round(termino.C * 10000)
			if i > 0 {
				restriccionStr += fmt.Sprintf(" + "+func() string {
					if tempC > termino.C*1000 {
						return fmt.Sprintf(`%.3f`, termino.C)
					} else {
						return fmt.Sprintf(`%d`, int(termino.C))
					}
				}()+"X%d", i+1)
			} else {
				restriccionStr += fmt.Sprintf(func() string {
					if tempC > termino.C*1000 {
						return fmt.Sprintf(`%.3f`, termino.C)
					} else {
						return fmt.Sprintf(`%d`, int(termino.C))
					}
				}()+"X%d", i+1)
			}
		}
		// Agregar el operador y el límite de la restricción
		restriccionStr += fmt.Sprintf(" %s "+func() string {
			tempC := restriccion.LD
			if tempC > restriccion.LD*1000 {
				return fmt.Sprintf(`%.3f`, restriccion.LD)
			} else {
				return fmt.Sprintf(`%d`, int(restriccion.LD))
			}
		}()+"\n", restriccion.Operador)
		restricciones += restriccionStr
	}

	// Restricciones de no negatividad para las variables
	variables := "X1"
	// Agregar las variables X2, X3, ..., Xn si es necesario
	for i := 1; i < len(req.FO); i++ {
		variables += fmt.Sprintf(", X%d", i+1)
	}
	// Asegurarse de que las variables sean mayores o iguales a 0
	variables += " ≥ 0"

	// Retornar la representación en formato texto de la función objetivo y las restricciones
	return fo, fmt.Sprintf("%s%s", restricciones, variables)
}

// Resultados genera los resultados de la solución del método Simplex.
func Resultados(sof models.TablaSimplex, numVD int) string {
	var resultados string
	// Iterar sobre las ecuaciones de la tabla Simplex
	for i, ecuacion := range sof.Ecuaciones {
		// Mostrar el valor de la variable básica (VB) y el lado derecho (LD) de cada ecuación
		resultados += fmt.Sprintf("%s = %f\n", ecuacion.VB, ecuacion.LD)
		// Si es la primera ecuación (ecuación de la base)
		if i == 0 {
			// Mostrar los costos reducidos de las variables de decisión (VD) y las variables artificiales
			for j, termino := range ecuacion.LI {
				if j < numVD {
					resultados += fmt.Sprintf("Costo reducido %s = %f\n", termino.VD, math.Round(termino.C*10000)/10000)
				} else {
					resultados += fmt.Sprintf("y%d = %f\n", j-numVD+1, math.Round(termino.C*10000)/10000)
				}
			}
		}
	}

	// Retornar los resultados como una cadena formateada
	return resultados
}
