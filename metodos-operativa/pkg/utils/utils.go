package utils

// Map aplica una función de transformación a cada elemento de un slice y devuelve un nuevo slice con los valores transformados.
func Map[T any, U any](slice []T, fn func(T, int) U) []U {
	mappedSlice := make([]U, len(slice))
	for i, v := range slice {
		mappedSlice[i] = fn(v, i)
	}
	return mappedSlice
}

// Función para hacer una copia profunda de una matriz [][]float64
func CopiarMatriz(matriz [][]float64) [][]float64 {
	copia := make([][]float64, len(matriz)) // Crear un nuevo slice de slices
	for i := range matriz {
		copia[i] = make([]float64, len(matriz[i])) // Crear una nueva fila
		copy(copia[i], matriz[i])                  // Copiar los valores de la fila original
	}
	return copia
}
