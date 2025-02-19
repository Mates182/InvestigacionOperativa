package grafos

import (
	"math"
	"metodos-operativa/internal/data/models"
)

// DFS modificado para encontrar un camino aumentante con capacidad disponible
func dfs(g *models.Grafo, origen, destino string, flujoActual float64, visitados map[string]bool, camino map[string]string) float64 {
	if origen == destino {
		return flujoActual
	}

	visitados[origen] = true
	for _, conexion := range g.Nodos[origen].Salidas {
		if !visitados[conexion.Destino] && conexion.Capacidad > 0 {
			minFlujo := math.Min(flujoActual, conexion.Capacidad)
			camino[conexion.Destino] = origen
			flujo := dfs(g, conexion.Destino, destino, minFlujo, visitados, camino)
			if flujo > 0 {
				return flujo
			}
		}
	}
	return 0
}

// Ford-Fulkerson que devuelve un grafo con la red de flujo máximo
func FordFulkersonGrafo(g *models.Grafo, origen, destino string) (float64, *models.Grafo) {
	flujoMaximo := 0.0
	flujoGrafo := models.NuevoGrafo() // Grafo para almacenar el flujo máximo

	for {
		visitados := make(map[string]bool)
		camino := make(map[string]string)

		flujo := dfs(g, origen, destino, math.MaxFloat64, visitados, camino)
		if flujo == 0 {
			break // No hay más caminos con capacidad disponible
		}

		// Construcción de la red de flujo máximo
		nodo := destino
		for nodo != origen {
			previo := camino[nodo]
			for i, conexion := range g.Nodos[previo].Salidas {
				if conexion.Destino == nodo {
					// Reducir capacidad en la arista utilizada
					g.Nodos[previo].Salidas[i].Capacidad -= flujo

					// Agregar conexión a la red de flujo con la capacidad utilizada
					flujoGrafo.AgregarConexion(previo, nodo, conexion.Costo, flujo, conexion.Distancia)
					break
				}
			}
			nodo = previo
		}

		flujoMaximo += flujo
	}

	return flujoMaximo, flujoGrafo
}
