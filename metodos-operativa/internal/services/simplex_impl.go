package services

import (
	"fmt"
	"metodos-operativa/internal/data/models"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/data/responses"
	"metodos-operativa/pkg/programacion_lineal"
	"metodos-operativa/pkg/utils"
)

func (s *Services) Simplex(r requests.ProgramacionLinealRequest) (int, responses.SimplexResponse) {
	// Inicializa la tabla Simplex con la primera iteración
	tablaSimplex := models.TablaSimplex{Iteracion: 0}

	// Número de restricciones en la función objetivo
	numRestricciones := len(r.FO)

	// Lista de variables de holgura que se agregarán a las restricciones
	var variablesHolgura []models.Termino

	// Generar variables de holgura s1, s2, ..., sn, dependiendo del número de restricciones
	for i := 1; i <= len(r.Restricciones); i++ {
		variablesHolgura = append(variablesHolgura, models.Termino{
			C:  0,                     // Coeficiente de la variable de holgura
			VD: fmt.Sprintf("s%d", i), // Nombre de la variable de holgura (s1, s2, ...)
		})
	}

	// Construcción de la función objetivo agregando las variables de holgura
	funcionObjetivo := programacion_lineal.CrearFuncionObjetivo(r, variablesHolgura, []models.Termino{})
	tablaSimplex.Ecuaciones = append(tablaSimplex.Ecuaciones, funcionObjetivo)

	// Construcción de las restricciones, agregando las variables de holgura a cada una
	for i, restriccion := range r.Restricciones {
		ecuacion := models.Ecuacion{
			Num: i + 1,                                       // Número de la restricción (1, 2, ...)
			VB:  fmt.Sprintf("s%d", i+1),                     // Variable básica de la restricción
			LI:  append(restriccion.LI, variablesHolgura...), // Agregar las variables de holgura
			LD:  restriccion.LD,                              // Lado derecho de la restricción
		}

		// Asignar coeficiente 1 a la variable de holgura correspondiente en la ecuación
		ecuacion.LI[numRestricciones+i].C = 1
		tablaSimplex.Ecuaciones = append(tablaSimplex.Ecuaciones, ecuacion) // Agregar ecuación a la tabla
	}

	// Resolver iterativamente el problema usando el método Simplex
	resolucion := programacion_lineal.ResolverIterativamente(tablaSimplex)

	// Convertir la programación lineal a un formato adecuado para la salida
	modeloFuncionObjetivo, modeloRestricciones := utils.ConvertirProgramacionLinealAFormato(r)
	var modelo []string
	modelo = append(modelo, modeloFuncionObjetivo) // Agregar la función objetivo
	modelo = append(modelo, modeloRestricciones)   // Agregar las restricciones

	// Obtener los resultados finales de la resolución del Simplex
	respuestas := utils.Resultados(resolucion[len(resolucion)-1], numRestricciones)

	// Retornar la solución óptima con la resolución, el modelo y las respuestas
	return 0, responses.SimplexResponse{
		Message:    "Solución óptima encontrada",
		Resolucion: resolucion,
		Metodo:     "simplex",
		Modelo:     modelo,
		Respuestas: respuestas,
	}
}
