package programacion_lineal

import (
	"fmt"
	"math"
	"metodos-operativa/internal/data/models"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/pkg/utils"
)

// Resuelve iterativamente el método Simplex
func ResolverIterativamente(tabla models.TablaSimplex) []models.TablaSimplex {
	var resolucion []models.TablaSimplex
	iteracion := tabla
	esOptimo := false

	// Añadir la primera tabla a la resolución
	resolucion = append(resolucion, iteracion)

	// Iterar hasta encontrar la solución óptima
	for !esOptimo {
		// Resolver la tabla Simplex y determinar si es óptima
		esOptimo, iteracion = ResolverTablaSimplex(resolucion[len(resolucion)-1])
		if esOptimo {
			break // Si es óptima, terminamos
		}
		// Se marca la iteración actual
		iteracion.Iteracion = len(resolucion) - 1
		// Se añade la tabla resultante de la iteración
		resolucion = append(resolucion, iteracion)
	}

	// Retornar todas las tablas generadas durante el proceso
	return resolucion
}

// Crea la función objetivo original
func CrearFuncionObjetivo(r requests.ProgramacionLinealRequest, variablesHolgura, variablesArtificiales []models.Termino) models.Ecuacion {
	// Modificar los coeficientes de la función objetivo según si es maximización
	li := append(utils.Map(r.FO, func(termino models.Termino, index int) models.Termino {
		coeficiente := termino.C
		if r.Maximizar {
			coeficiente *= -1 // Si es maximización, se invierte el signo
		}
		return models.Termino{C: coeficiente, VD: termino.VD}
	}), variablesHolgura...) // Añadir variables de holgura

	// Añadir las variables artificiales, si las hay
	if len(variablesArtificiales) > 0 {
		li = append(li, variablesArtificiales...)
	}

	// Crear la ecuación de la función objetivo
	return models.Ecuacion{
		Num: 0,   // Número de la ecuación
		VB:  "Z", // Nombre de la variable básica
		LI:  li,  // Lista de términos de la función
		LD:  0,   // Lado derecho (valor de la ecuación)
	}
}

// Resuelve una tabla de Simplex
func ResolverTablaSimplex(tabla models.TablaSimplex) (bool, models.TablaSimplex) {
	tablaIteracion := models.TablaSimplex{}

	// 1. Seleccionar la variable entrante (más negativa)
	indiceVariableEntrante := -1
	variableEntrante := 0.0

	// Buscar el coeficiente más negativo en la fila de la función objetivo
	for i, termino := range tabla.Ecuaciones[0].LI {
		if (math.Trunc(termino.C*100000000) / 100000000) < variableEntrante {
			variableEntrante = termino.C
			indiceVariableEntrante = i
		}
	}

	fmt.Println(variableEntrante)

	// Si no se encuentra un coeficiente negativo, hemos alcanzado la solución óptima
	if indiceVariableEntrante == -1 {
		return true, tabla // La solución es óptima
	}

	// 2. Seleccionar la variable saliente (usando el criterio del cociente mínimo positivo)
	indiceVariableSaliente := -1
	minimoCocientePositivo := 0.0
	variableSaliente := 0.0
	var variableBasicaEntrante string

	// Iterar sobre las filas de las restricciones para calcular los cocientes
	for i := 1; i < len(tabla.Ecuaciones); i++ {
		if tabla.Ecuaciones[i].LI[indiceVariableEntrante].C > 0 {
			cociente := tabla.Ecuaciones[i].LD / tabla.Ecuaciones[i].LI[indiceVariableEntrante].C
			if cociente > 0 && (indiceVariableSaliente == -1 || cociente < minimoCocientePositivo) {
				minimoCocientePositivo = cociente
				variableSaliente = tabla.Ecuaciones[i].LI[indiceVariableEntrante].C
				indiceVariableSaliente = i
				variableBasicaEntrante = tabla.Ecuaciones[i].LI[indiceVariableEntrante].VD
			}
		}
	}

	// Si no se encontró una variable saliente, significa que no hay solución óptima
	if indiceVariableSaliente == -1 {
		return true, tabla // No hay solución óptima
	}

	// 3. Normalizar la ecuación pivote
	ecuacionPivote := models.Ecuacion{
		LI: utils.Map(tabla.Ecuaciones[indiceVariableSaliente].LI, func(termino models.Termino, index int) models.Termino {
			// Normalizar los coeficientes de la ecuación pivote
			return models.Termino{C: termino.C / variableSaliente, VD: termino.VD}
		}),
		LD:  tabla.Ecuaciones[indiceVariableSaliente].LD / variableSaliente,
		VB:  variableBasicaEntrante, // La nueva variable básica
		Num: tabla.Ecuaciones[indiceVariableSaliente].Num,
	}

	// 4. Actualizar el resto de las ecuaciones
	for i, ecuacion := range tabla.Ecuaciones {
		if i != indiceVariableSaliente {
			// Actualizamos las ecuaciones que no son la pivote
			tablaIteracion.Ecuaciones = append(tablaIteracion.Ecuaciones, models.Ecuacion{
				LI: utils.Map(ecuacion.LI, func(termino models.Termino, index int) models.Termino {
					// Aplicamos la eliminación de Gauss-Jordan a los coeficientes
					return models.Termino{
						C:  termino.C - (ecuacionPivote.LI[index].C * ecuacion.LI[indiceVariableEntrante].C),
						VD: termino.VD,
					}
				}),
				LD:  ecuacion.LD - (ecuacionPivote.LD * ecuacion.LI[indiceVariableEntrante].C),
				VB:  ecuacion.VB,
				Num: ecuacion.Num,
			})
		} else {
			// Añadimos la ecuación pivote que ya ha sido normalizada
			tablaIteracion.Ecuaciones = append(tablaIteracion.Ecuaciones, ecuacionPivote)
		}
	}

	// Retornar el resultado de la iteración
	return false, tablaIteracion
}
