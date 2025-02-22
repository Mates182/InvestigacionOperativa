package grafos

import (
	"math"
	"metodos-operativa/internal/data/models"
)

// DFS (Búsqueda en Profundidad) para encontrar un camino aumentante con capacidad disponible
func dfs(g *models.Grafo, origen, destino string, flujoActual float64, visitados map[string]bool, camino map[string]string) float64 {
	// Si hemos llegado al destino, devolvemos el flujo actual
	if origen == destino {
		return flujoActual
	}

	// Marcamos el nodo actual como visitado
	visitados[origen] = true

	// Iteramos por las conexiones del nodo de origen
	for _, conexion := range g.Nodos[origen].Salidas {
		// Si no se ha visitado el nodo de destino y la capacidad de la conexión es mayor que 0
		if !visitados[conexion.Destino] && conexion.Capacidad > 0 {
			// Calculamos el flujo mínimo entre el flujo actual y la capacidad de la conexión
			minFlujo := math.Min(flujoActual, conexion.Capacidad)
			// Registramos el camino de origen a destino
			camino[conexion.Destino] = origen
			// Llamada recursiva para buscar el camino aumentante
			flujo := dfs(g, conexion.Destino, destino, minFlujo, visitados, camino)
			// Si encontramos un flujo positivo, lo retornamos
			if flujo > 0 {
				return flujo
			}
		}
	}
	// Si no encontramos un flujo válido, retornamos 0
	return 0
}

// Ford-Fulkerson que devuelve un grafo con la red de flujo máximo
func FordFulkersonGrafo(g *models.Grafo, origen, destino string) (float64, *models.Grafo) {
	flujoMaximo := 0.0
	flujoGrafo := models.NuevoGrafo() // Grafo para almacenar el flujo máximo

	// Continuamos buscando caminos aumentantes mientras los haya
	for {
		visitados := make(map[string]bool)
		camino := make(map[string]string)

		// Buscamos un camino aumentante desde el origen hasta el destino
		flujo := dfs(g, origen, destino, math.MaxFloat64, visitados, camino)
		// Si no encontramos un flujo válido, terminamos el algoritmo
		if flujo == 0 {
			break // No hay más caminos con capacidad disponible
		}

		// Construcción de la red de flujo máximo
		nodo := destino
		// Retrocedemos por el camino encontrado
		for nodo != origen {
			previo := camino[nodo]
			// Buscamos la conexión correspondiente entre los nodos
			for i, conexion := range g.Nodos[previo].Salidas {
				if conexion.Destino == nodo {
					// Reducir capacidad en la arista utilizada
					g.Nodos[previo].Salidas[i].Capacidad -= flujo

					// Agregar la conexión con la capacidad utilizada al grafo de flujo
					flujoGrafo.AgregarConexion(previo, nodo, conexion.Costo, flujo, conexion.Distancia)
					break
				}
			}
			nodo = previo
		}

		// Aumentamos el flujo máximo con el flujo encontrado
		flujoMaximo += flujo
	}

	// Retornamos el flujo máximo y el grafo con las conexiones de flujo
	return flujoMaximo, flujoGrafo
}
