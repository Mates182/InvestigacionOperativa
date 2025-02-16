package programacion_lineal

import (
	"metodos-operativa/internal/data/models"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/pkg/utils"
)

// Resuelve iterativamente el método Simplex
func ResolverIterativamente(tabla models.TablaSimplex) []models.TablaSimplex {
	var resolucion []models.TablaSimplex
	iteracion := tabla
	esOptimo := false

	resolucion = append(resolucion, iteracion)
	for !esOptimo {
		esOptimo, iteracion = ResolverTablaSimplex(resolucion[len(resolucion)-1])
		if esOptimo {
			break
		}
		iteracion.Iteracion = len(resolucion) - 1
		resolucion = append(resolucion, iteracion)
	}
	return resolucion
}

// Crea la función objetivo original
func CrearFuncionObjetivo(r requests.ProgramacionLinealRequest, variablesHolgura, variablesArtificiales []models.Termino) models.Ecuacion {
	li := append(utils.Map(r.FO, func(termino models.Termino, index int) models.Termino {
		coeficiente := termino.C
		if r.Maximizar {
			coeficiente *= -1
		}
		return models.Termino{C: coeficiente, VD: termino.VD}
	}), variablesHolgura...)

	if len(variablesArtificiales) > 0 {
		li = append(li, variablesArtificiales...)
	}

	return models.Ecuacion{
		Num: 0,
		VB:  "Z",
		LI:  li,
		LD:  0,
	}
}

// Resuelve una tabla de Simplex
func ResolverTablaSimplex(tabla models.TablaSimplex) (bool, models.TablaSimplex) {
	tablaIteracion := models.TablaSimplex{}

	// 1. Seleccionar variable entrante
	indiceVariableEntrante := -1
	variableEntrante := 0.0

	for i, termino := range tabla.Ecuaciones[0].LI {
		if termino.C < variableEntrante {
			variableEntrante = termino.C
			indiceVariableEntrante = i
		}
	}

	if indiceVariableEntrante == -1 {
		return true, tabla // La solución es óptima
	}

	// 2. Seleccionar variable saliente
	indiceVariableSaliente := -1
	minimoCocientePositivo := 0.0
	variableSaliente := 0.0
	var variableBasicaEntrante string

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

	if indiceVariableSaliente == -1 {
		return true, tabla // No hay solución óptima
	}

	// 3. Normalizar ecuación pivote
	ecuacionPivote := models.Ecuacion{
		LI: utils.Map(tabla.Ecuaciones[indiceVariableSaliente].LI, func(termino models.Termino, index int) models.Termino {
			return models.Termino{C: termino.C / variableSaliente, VD: termino.VD}
		}),
		LD:  tabla.Ecuaciones[indiceVariableSaliente].LD / variableSaliente,
		VB:  variableBasicaEntrante,
		Num: tabla.Ecuaciones[indiceVariableSaliente].Num,
	}

	// 4. Actualizar el resto de ecuaciones
	for i, ecuacion := range tabla.Ecuaciones {
		if i != indiceVariableSaliente {
			tablaIteracion.Ecuaciones = append(tablaIteracion.Ecuaciones, models.Ecuacion{
				LI: utils.Map(ecuacion.LI, func(termino models.Termino, index int) models.Termino {
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
			tablaIteracion.Ecuaciones = append(tablaIteracion.Ecuaciones, ecuacionPivote)
		}
	}

	return false, tablaIteracion
}
