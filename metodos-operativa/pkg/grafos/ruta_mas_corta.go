package grafos

import (
	"container/heap"
	"math"
	"metodos-operativa/internal/data/models"
)

// Dijkstra que devuelve un grafo con la ruta más corta y permite elegir entre costo o distancia
func DijkstraGrafo(g *models.Grafo, origen, destino string, usarCosto bool) (float64, *models.Grafo) {
	distancia := make(map[string]float64)
	predecesor := make(map[string]string)

	for nodo := range g.Nodos {
		distancia[nodo] = math.Inf(1)
	}
	distancia[origen] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{value: origen, priority: 0})

	for pq.Len() > 0 {
		actual := heap.Pop(pq).(*Item).value
		if actual == destino {
			break
		}

		for _, conexion := range g.Nodos[actual].Salidas {
			valor := conexion.Costo
			if !usarCosto {
				valor = conexion.Distancia
			}

			nuevaDist := distancia[actual] + valor
			if nuevaDist < distancia[conexion.Destino] {
				distancia[conexion.Destino] = nuevaDist
				predecesor[conexion.Destino] = actual
				heap.Push(pq, &Item{value: conexion.Destino, priority: nuevaDist})
			}
		}
	}

	// Construir el grafo de la ruta más corta
	rutaGrafo := models.NuevoGrafo()
	nodo := destino
	for nodo != "" {
		previo, existe := predecesor[nodo]
		if !existe {
			break
		}
		for _, conexion := range g.Nodos[previo].Salidas {
			if conexion.Destino == nodo {
				rutaGrafo.AgregarConexion(previo, nodo, conexion.Costo, conexion.Capacidad, conexion.Distancia)
				break
			}
		}
		nodo = previo
	}

	return distancia[destino], rutaGrafo
}

// Priority Queue para Dijkstra
type Item struct {
	value    string  // Nombre del nodo
	priority float64 // Costo o distancia actual
	index    int     // Índice en la heap
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].priority < pq[j].priority }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
