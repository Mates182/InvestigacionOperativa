package transporte

import (
	"fmt"
	"math"
	"metodos-operativa/internal/data/requests"
	"strings"
)

// Función para calcular la diferencia de costos más bajos en cada fila y columna
func calcularDiferencias(costos [][]float64) ([]float64, []float64) {
	numOrigenes := len(costos)    // Número de orígenes
	numDestinos := len(costos[0]) // Número de destinos

	// Diferencias de oferta (filas)
	difOferta := make([]float64, numOrigenes)
	for i := 0; i < numOrigenes; i++ {
		min1, min2 := math.MaxFloat64, math.MaxFloat64
		for j := 0; j < numDestinos; j++ {
			if costos[i][j] < min1 {
				min2 = min1
				min1 = costos[i][j]
			} else if costos[i][j] < min2 {
				min2 = costos[i][j]
			}
		}
		if min2 == math.MaxFloat64 { // Si solo hay un costo en la fila
			min2 = min1
		}
		difOferta[i] = min2 - min1 // Diferencia entre los dos menores costos en la fila
	}

	// Diferencias de demanda (columnas)
	difDemanda := make([]float64, numDestinos)
	for j := 0; j < numDestinos; j++ {
		min1, min2 := math.MaxFloat64, math.MaxFloat64
		for i := 0; i < numOrigenes; i++ {
			if costos[i][j] < min1 {
				min2 = min1
				min1 = costos[i][j]
			} else if costos[i][j] < min2 {
				min2 = costos[i][j]
			}
		}
		if min2 == math.MaxFloat64 { // Si solo hay un costo en la columna
			min2 = min1
		}
		difDemanda[j] = min2 - min1 // Diferencia entre los dos menores costos en la columna
	}

	return difOferta, difDemanda // Retornar las diferencias para oferta y demanda
}

// Función para aplicar el método de Vogel
func Vogel(costosTemp [][]float64, oferta []float64, demanda []float64) [][]float64 {
	costos := costosTemp                         // Copia de los costos de transporte
	numOrigenes := len(costos)                   // Número de orígenes
	numDestinos := len(costos[0])                // Número de destinos
	asignacion := make([][]float64, numOrigenes) // Matriz de asignación vacía

	// Inicializar la matriz de asignación
	for i := range asignacion {
		asignacion[i] = make([]float64, numDestinos)
	}

	// Ciclo para realizar la asignación con el método de Vogel
	for {
		// Calcular diferencias de oferta y demanda
		difOferta, difDemanda := calcularDiferencias(costos)

		// Buscar la mayor penalización entre oferta y demanda
		maxPenalizacion := -1.0
		seleccionI, seleccionJ := -1, -1
		esFila := true

		// Buscar la mayor penalización en las filas (oferta)
		for i := 0; i < numOrigenes; i++ {
			if oferta[i] > 0 && difOferta[i] > maxPenalizacion {
				maxPenalizacion = difOferta[i]
				seleccionI = i
				esFila = true
			}
		}

		// Buscar la mayor penalización en las columnas (demanda)
		for j := 0; j < numDestinos; j++ {
			if demanda[j] > 0 && difDemanda[j] > maxPenalizacion {
				maxPenalizacion = difDemanda[j]
				seleccionI = j
				esFila = false
			}
		}

		// Si no hay más oferta o demanda, terminamos el algoritmo
		if maxPenalizacion == -1 {
			break
		}

		// Encontrar el índice con el menor costo en la fila o columna seleccionada
		if esFila {
			minCosto := math.MaxFloat64
			for j := 0; j < numDestinos; j++ {
				if demanda[j] > 0 && costos[seleccionI][j] < minCosto {
					minCosto = costos[seleccionI][j]
					seleccionJ = j
				}
			}
		} else {
			minCosto := math.MaxFloat64
			for i := 0; i < numOrigenes; i++ {
				if oferta[i] > 0 && costos[i][seleccionI] < minCosto {
					minCosto = costos[i][seleccionI]
					seleccionJ = seleccionI
					seleccionI = i
				}
			}
		}

		// Determinar la cantidad a asignar en la celda seleccionada
		asignacionCantidad := math.Min(oferta[seleccionI], demanda[seleccionJ])
		asignacion[seleccionI][seleccionJ] = asignacionCantidad

		// Reducir la oferta y demanda correspondientes
		oferta[seleccionI] -= asignacionCantidad
		demanda[seleccionJ] -= asignacionCantidad

		// Si la oferta es 0, se anula la fila
		if oferta[seleccionI] == 0 {
			for j := 0; j < numDestinos; j++ {
				costos[seleccionI][j] = math.MaxFloat64
			}
		}

		// Si la demanda es 0, se anula la columna
		if demanda[seleccionJ] == 0 {
			for i := 0; i < numOrigenes; i++ {
				costos[i][seleccionJ] = math.MaxFloat64
			}
		}
	}

	// Retornar la matriz de asignación final
	return asignacion
}

// Función para imprimir la matriz
func ImprimirMatriz(matriz [][]float64) {
	for _, fila := range matriz {
		for _, val := range fila {
			fmt.Printf("%8.2f", val)
		}
		fmt.Println() // Nueva línea después de cada fila
	}
}

// Función para calcular el costo total basado en la asignación
func CalcularCostoTotal(costos [][]float64, asignacion [][]float64) float64 {
	costoTotal := 0.0
	// Calcular el costo total multiplicando los costos por las cantidades asignadas
	for i := 0; i < len(costos); i++ {
		for j := 0; j < len(costos[i]); j++ {
			costoTotal += costos[i][j] * asignacion[i][j]
			fmt.Printf("Costo de (%f * %f): %f\n", costos[i][j], asignacion[i][j], costoTotal) // Mostrar detalle del cálculo
		}
	}
	return costoTotal // Retornar el costo total
}

// Función para formatear la asignación como cadena de texto
func FormatearAsignacion(req requests.TransporteRequest, asignacion [][]float64, costoTotal float64) string {
	var resultado strings.Builder

	// Recorrer las rutas de origen y destino y formatear la asignación
	for i, origen := range req.Origenes {
		for j, destino := range req.Destinos {
			if asignacion[i][j] > 0 {
				costo := req.Costos[i][j]
				resultado.WriteString(fmt.Sprintf(
					"Ruta %s -> %s : Unidades asignadas: %.2f  Costo Unitario: %.2f\n",
					origen.Origen, destino.Destino, asignacion[i][j], costo,
				))
			}
		}
	}

	// Agregar el costo total al final
	resultado.WriteString(fmt.Sprintf("\nCosto Total del Transporte: %.2f\n", costoTotal))

	return resultado.String() // Retornar el resultado formateado
}
