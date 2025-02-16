package services

import (
	"fmt"
	"metodos-operativa/internal/data/models"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/data/responses"
	"metodos-operativa/pkg/programacion_lineal"
	"metodos-operativa/pkg/utils"
)

type Services struct {
}

func NewServices() Service {
	return &Services{}
}

// Método para resolver el problema utilizando el método de Dos Fases
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
			variablesHolgura = append(variablesHolgura, models.Termino{
				C:  0,
				VD: fmt.Sprintf("s%d", len(variablesHolgura)+1),
			})
			indicesHolgura = append(indicesHolgura, i)
		}
	}

	for i, restriccion := range r.Restricciones {
		if restriccion.Operador == "=" || restriccion.Operador == "\u2265" {
			variablesArtificiales = append(variablesArtificiales, models.Termino{
				C:  0,
				VD: fmt.Sprintf("a%d", len(variablesArtificiales)+1),
			})
			indicesArtificiales = append(indicesArtificiales, i)
		}
	}

	// Crear función objetivo para la segunda fase
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

		// Asignar coeficientes a variables de holgura y artificiales
		if restriccion.Operador == "\u2264" {
			ecuacion.LI[restriccionesLen+indicesHolgura[contadorHolgura]].C = 1
			contadorHolgura++
		} else if restriccion.Operador == "\u2265" {
			ecuacion.LI[restriccionesLen+indicesHolgura[contadorHolgura]].C = -1
			contadorHolgura++
			ecuacion.LI[restriccionesLen+len(indicesHolgura)+contadorArtificiales].C = 1
			contadorArtificiales++
		} else {
			ecuacion.LI[restriccionesLen+len(indicesHolgura)+contadorArtificiales].C = 1
			contadorArtificiales++
		}

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

	for i, ecuacion := range tablaFase2.Ecuaciones {
		tablaFase2.Ecuaciones[i].LI = ecuacion.LI[:len(ecuacion.LI)-len(indicesArtificiales)]
		for k := 1; k <= len(tablaFase2.Ecuaciones); k++ {
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

	respuestas := utils.Resultados(resolucionFase2[len(resolucionFase2)-1], restriccionesLen)

	return 0, responses.DosFasesResponse{Message: "Solución óptima encontrada", Resolucion: responses.DosFasesResolucion{ResolucionFase1: resolucionFase1, ResolucionFase2: resolucionFase2}, Metodo: "dos fases", Modelo: modelo, Respuestas: respuestas}
}

func (s *Services) Simplex(r requests.ProgramacionLinealRequest) (int, responses.SimplexResponse) {
	// Inicializa la tabla Simplex con la primera iteración
	tablaSimplex := models.TablaSimplex{Iteracion: 0}

	// Número de restricciones en la función objetivo
	numRestricciones := len(r.FO)

	// Lista de variables de holgura
	var variablesHolgura []models.Termino

	// Generar variables de holgura s1, s2, ..., sn
	for i := 1; i <= len(r.Restricciones); i++ {
		variablesHolgura = append(variablesHolgura, models.Termino{
			C:  0,
			VD: fmt.Sprintf("s%d", i),
		})
	}

	// Construcción de la función objetivo
	funcionObjetivo := programacion_lineal.CrearFuncionObjetivo(r, variablesHolgura, []models.Termino{})
	tablaSimplex.Ecuaciones = append(tablaSimplex.Ecuaciones, funcionObjetivo)

	// Construcción de las restricciones
	for i, restriccion := range r.Restricciones {
		ecuacion := models.Ecuacion{
			Num: i + 1,
			VB:  fmt.Sprintf("s%d", i+1),
			LI:  append(restriccion.LI, variablesHolgura...), // Agregar variables de holgura
			LD:  restriccion.LD,
		}

		// Asignar coeficiente 1 a la variable de holgura correspondiente
		ecuacion.LI[numRestricciones+i].C = 1
		tablaSimplex.Ecuaciones = append(tablaSimplex.Ecuaciones, ecuacion)
	}

	// Resolver iterativamente usando el método Simplex
	resolucion := programacion_lineal.ResolverIterativamente(tablaSimplex)
	modeloFuncionObjetivo, modeloRestricciones := utils.ConvertirProgramacionLinealAFormato(r)
	var modelo []string
	modelo = append(modelo, modeloFuncionObjetivo)
	modelo = append(modelo, modeloRestricciones)
	respuestas := utils.Resultados(resolucion[len(resolucion)-1], numRestricciones)
	// Retornar la solución óptima
	return 0, responses.SimplexResponse{Message: "Solución óptima encontrada", Resolucion: resolucion, Metodo: "simplex", Modelo: modelo, Respuestas: respuestas}
}
