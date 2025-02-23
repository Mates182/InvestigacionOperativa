package services

import (
	"fmt"
	"metodos-operativa/internal/data/models"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/data/responses"
	"metodos-operativa/pkg/programacion_lineal"
	"metodos-operativa/pkg/utils"
)

// DosFases resuelve un problema de programación lineal usando el método de Dos Fases.
// La función maneja restricciones y variables artificiales o de holgura para formular y resolver el problema.
func (s *Services) DosFases(r requests.ProgramacionLinealRequest) (int, responses.DosFasesResponse) {
	tablaFase1 := models.TablaSimplex{Iteracion: 0}
	tablaFase2 := models.TablaSimplex{Iteracion: 0}
	restriccionesLen := len(r.FO)

	var variablesHolgura []models.Termino
	var indicesHolgura []int
	var variablesArtificiales []models.Termino
	var indicesArtificiales []int

	// Identificar variables de holgura y artificiales según los operadores de restricción
	for i, restriccion := range r.Restricciones {
		if restriccion.Operador != "=" {
			// Se añade una variable de holgura si la restricción no es de igualdad
			variablesHolgura = append(variablesHolgura, models.Termino{
				C:  0,
				VD: fmt.Sprintf("s%d", len(variablesHolgura)+1),
			})
			indicesHolgura = append(indicesHolgura, i)
		}
	}

	// Identificar variables artificiales según las restricciones de tipo "=" o ">="
	for i, restriccion := range r.Restricciones {
		if restriccion.Operador == "=" || restriccion.Operador == "\u2265" {
			variablesArtificiales = append(variablesArtificiales, models.Termino{
				C:  0,
				VD: fmt.Sprintf("a%d", len(variablesArtificiales)+1),
			})
			indicesArtificiales = append(indicesArtificiales, i)
		}
	}

	// Crear la función objetivo para la segunda fase
	funcionObjetivoFase2 := programacion_lineal.CrearFuncionObjetivo(r, variablesHolgura, variablesArtificiales)

	contadorHolgura := 0
	contadorArtificiales := 0

	// Construcción de las restricciones para la primera fase
	for i, restriccion := range r.Restricciones {
		ecuacion := models.Ecuacion{
			Num: i + 1,
			VB:  fmt.Sprintf("s%d", i+1),
			LI:  append(append(restriccion.LI, variablesHolgura...), variablesArtificiales...),
			LD:  restriccion.LD,
		}

		// Asignar coeficientes a variables de holgura y artificiales según el operador de la restricción
		if restriccion.Operador == "\u2264" {
			ecuacion.LI[restriccionesLen+indicesHolgura[contadorHolgura]].C = 1
			contadorHolgura++
		} else if restriccion.Operador == "\u2265" {
			// Ajustar variables de holgura y artificiales para restricciones ">="
			ecuacion.LI[restriccionesLen+indicesHolgura[contadorHolgura]].C = -1
			contadorHolgura++
			ecuacion.LI[restriccionesLen+len(indicesHolgura)+contadorArtificiales].C = 1
			contadorArtificiales++
		} else {
			// Para restricciones de igualdad
			ecuacion.LI[restriccionesLen+len(indicesHolgura)+contadorArtificiales].C = 1
			contadorArtificiales++
		}

		// Añadir la ecuación a la tabla de la fase 1
		tablaFase1.Ecuaciones = append(tablaFase1.Ecuaciones, ecuacion)
	}

	// Construcción de la función objetivo para la primera fase
	contadorArtificiales = 0
	funcionObjetivoFase1 := models.Ecuacion{
		Num: 0,
		VB:  "Z",
		LI: utils.Map(funcionObjetivoFase2.LI, func(termino models.Termino, index int) models.Termino {
			return models.Termino{
				C: func() float64 {
					// Asignar coeficiente de 1 a las variables artificiales en la primera fase
					if contadorArtificiales < len(indicesArtificiales) && index == restriccionesLen+len(indicesHolgura)+contadorArtificiales {
						contadorArtificiales++
						return 1
					}
					return 0
				}(),
				VD: termino.VD,
			}
		}),
	}

	// Ajuste de coeficientes en la función objetivo de la primera fase
	if len(variablesArtificiales) > 0 {
		for i := 0; i < len(funcionObjetivoFase2.LI); i++ {
			funcionObjetivoFase1.LI[i] = models.Termino{
				C: func() float64 {
					// Ajustar los coeficientes de las variables artificiales
					suma := funcionObjetivoFase1.LI[i].C
					for _, index := range indicesArtificiales {
						ec := tablaFase1.Ecuaciones[index]
						suma -= ec.LI[i].C
					}
					return suma
				}(),
				VD: funcionObjetivoFase2.LI[i].VD,
			}
		}
		funcionObjetivoFase1.LD = func() float64 {
			// Ajustar el término independiente de la función objetivo
			suma := 0.0
			for _, index := range indicesArtificiales {
				ec := tablaFase1.Ecuaciones[index]
				suma -= ec.LD
			}
			return suma
		}()
	}

	// Agregar la función objetivo ajustada a la tabla de la primera fase
	tablaTemporal := models.TablaSimplex{}
	tablaTemporal.Ecuaciones = append(tablaTemporal.Ecuaciones, funcionObjetivoFase1)
	tablaFase1.Ecuaciones = append(tablaTemporal.Ecuaciones, tablaFase1.Ecuaciones...)

	// Resolver la primera fase del método de Dos Fases
	resolucionFase1 := programacion_lineal.ResolverIterativamente(tablaFase1)

	// Preparar la segunda fase con la solución obtenida en la primera fase
	tablaFase2.Ecuaciones = append(tablaFase2.Ecuaciones, resolucionFase1[len(resolucionFase1)-1].Ecuaciones...)
	tablaFase2.Ecuaciones[0] = funcionObjetivoFase2

	// Ajustar las ecuaciones de la segunda fase
	for i, ecuacion := range tablaFase2.Ecuaciones {
		tablaFase2.Ecuaciones[i].LI = ecuacion.LI[:len(ecuacion.LI)-len(indicesArtificiales)]
		for k := 1; k <= len(tablaFase2.Ecuaciones); k++ {
			// Realizar pivoteo en la segunda fase
			if ecuacion.VB == fmt.Sprintf("x%d", k) {
				cpivot := tablaFase2.Ecuaciones[0].LI[k-1].C
				tablaFase2.Ecuaciones[0].LI = utils.Map(tablaFase2.Ecuaciones[0].LI, func(termino models.Termino, index int) models.Termino {
					return models.Termino{
						C:  termino.C - (cpivot * ecuacion.LI[index].C),
						VD: termino.VD,
					}
				})
				tablaFase2.Ecuaciones[0].LD -= cpivot * ecuacion.LD
			}
		}
	}

	// Resolver la segunda fase
	resolucionFase2 := programacion_lineal.ResolverIterativamente(tablaFase2)
	modeloFuncionObjetivo, modeloRestricciones := utils.ConvertirProgramacionLinealAFormato(r)
	var modelo []string
	modelo = append(modelo, modeloFuncionObjetivo)
	modelo = append(modelo, modeloRestricciones)
	if !r.Maximizar {
		for _, iteracion := range resolucionFase2 {
			iteracion.Ecuaciones[0].LD = iteracion.Ecuaciones[0].LD * -1
			/*for i, termino := range iteracion.Ecuaciones[0].LI {
				iteracion.Ecuaciones[0].LI[i].C = termino.C * -1
			}*/
		}

	}

	// Obtener los resultados finales
	respuestas := utils.Resultados(resolucionFase2[len(resolucionFase2)-1], restriccionesLen)

	// Retornar la respuesta final con la resolución de ambas fases
	return 0, responses.DosFasesResponse{Message: "Solución óptima encontrada", Resolucion: responses.DosFasesResolucion{ResolucionFase1: resolucionFase1, ResolucionFase2: resolucionFase2}, Metodo: "dos fases", Modelo: modelo, Respuestas: respuestas}
}
