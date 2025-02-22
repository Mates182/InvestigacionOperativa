package grafos

import (
	"container/heap"
	"math"
	"metodos-operativa/internal/data/models"
)

// Dijkstra que devuelve un grafo con la ruta más corta y permite elegir entre costo o distancia
func DijkstraGrafo(g *models.Grafo, origen, destino string, usarCosto bool) (float64, *models.Grafo) {
	// Inicializa las distancias de todos los nodos a infinito
	distancia := make(map[string]float64)
	// Mapa para almacenar el nodo predecesor de cada nodo en la ruta más corta
	predecesor := make(map[string]string)

	// Se asigna infinito a todas las distancias
	for nodo := range g.Nodos {
		distancia[nodo] = math.Inf(1)
	}
	// La distancia al nodo origen es 0
	distancia[origen] = 0

	// Inicializa la cola de prioridad (min-heap)
	pq := &PriorityQueue{}
	heap.Init(pq)
	// Se agrega el nodo origen con prioridad 0
	heap.Push(pq, &Item{value: origen, priority: 0})

	// Mientras haya nodos en la cola de prioridad
	for pq.Len() > 0 {
		// Se toma el nodo con la menor prioridad (distancia más corta)
		actual := heap.Pop(pq).(*Item).value
		// Si hemos llegado al nodo destino, terminamos
		if actual == destino {
			break
		}

		// Se recorren las conexiones del nodo actual
		for _, conexion := range g.Nodos[actual].Salidas {
			// Se elige el valor de la conexión basado en si se utiliza costo o distancia
			valor := conexion.Costo
			if !usarCosto {
				valor = conexion.Distancia
			}

			// Calculamos la nueva distancia hasta el nodo de destino a través de esta conexión
			nuevaDist := distancia[actual] + valor
			// Si la nueva distancia es menor a la ya conocida, se actualiza
			if nuevaDist < distancia[conexion.Destino] {
				distancia[conexion.Destino] = nuevaDist
				predecesor[conexion.Destino] = actual
				// Se agrega el nodo vecino con su nueva prioridad (distancia)
				heap.Push(pq, &Item{value: conexion.Destino, priority: nuevaDist})
			}
		}
	}

	// Construir el grafo de la ruta más corta
	rutaGrafo := models.NuevoGrafo()
	nodo := destino
	// Seguimos el camino desde el destino hasta el origen, siguiendo los predecesores
	for nodo != "" {
		// Obtenemos el nodo predecesor
		previo, existe := predecesor[nodo]
		if !existe {
			break // Si no hay predecesor, significa que no hay ruta
		}
		// Se agrega la conexión correspondiente al grafo de la ruta más corta
		for _, conexion := range g.Nodos[previo].Salidas {
			if conexion.Destino == nodo {
				rutaGrafo.AgregarConexion(previo, nodo, conexion.Costo, conexion.Capacidad, conexion.Distancia)
				break
			}
		}
		// Actualizamos el nodo al predecesor
		nodo = previo
	}

	// Retornamos la distancia más corta al destino y el grafo con la ruta más corta
	return distancia[destino], rutaGrafo
}

// Priority Queue para Dijkstra
type Item struct {
	value    string  // Nombre del nodo
	priority float64 // Costo o distancia actual
	index    int     // Índice en la heap
}

type PriorityQueue []*Item

// Métodos necesarios para implementar la interfaz heap en Go
func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Método para agregar un item a la cola de prioridad
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Método para remover el item con la menor prioridad de la cola
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
