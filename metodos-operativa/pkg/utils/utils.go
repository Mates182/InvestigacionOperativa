package utils

// Map aplica una función de transformación a cada elemento de un slice y devuelve un nuevo slice con los valores transformados.
func Map[T any, U any](slice []T, fn func(T, int) U) []U {
	mappedSlice := make([]U, len(slice))
	for i, v := range slice {
		mappedSlice[i] = fn(v, i)
	}
	return mappedSlice
}
