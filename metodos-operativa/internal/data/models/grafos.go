package models

import "fmt"

// Representa la ruta entre dos nodos con sus datos de costo, capacidad y distancia
type Conexion struct {
	Destino   string  `json:"destino"`   // Nodo destino
	Costo     float64 `json:"costo"`     // Costo de la conexión
	Capacidad float64 `json:"capacidad"` // Capacidad máxima de flujo
	Distancia float64 `json:"distancia"` // Distancia entre los nodos
}

// Estructura de un nodo en el grafo
type Nodo struct {
	Nombre   string     `json:"nombre"`   // Identificador del nodo
	Entradas []string   `json:"entradas"` // Nodos de donde recibe flujo
	Salidas  []Conexion `json:"salidas"`  // Conexiones hacia otros nodos o destinos
}

// Grafo dirigido con costos y capacidades
type Grafo struct {
	Nodos map[string]*Nodo `json:"nodos"` // Mapa de nodos por nombre
}

// Crear un nuevo grafo vacío
func NuevoGrafo() *Grafo {
	return &Grafo{Nodos: make(map[string]*Nodo)}
}

// Agregar un nodo al grafo
func (g *Grafo) AgregarNodo(nombre string) {
	if _, existe := g.Nodos[nombre]; !existe {
		g.Nodos[nombre] = &Nodo{Nombre: nombre}
	}
}

// Agregar una conexión entre nodos con costo y capacidad
func (g *Grafo) AgregarConexion(origen, destino string, costo, capacidad, distancia float64) {
	g.AgregarNodo(origen)  // Asegurar que el nodo origen exista
	g.AgregarNodo(destino) // Asegurar que el nodo destino exista

	// Agregar salida al nodo origen con distancia incluida
	g.Nodos[origen].Salidas = append(g.Nodos[origen].Salidas, Conexion{
		Destino:   destino,
		Costo:     costo,
		Capacidad: capacidad,
		Distancia: distancia,
	})

	// Agregar entrada al nodo destino
	g.Nodos[destino].Entradas = append(g.Nodos[destino].Entradas, origen)
}

// Mostrar el grafo con sus conexiones
func (g *Grafo) Mostrar() {
	for _, nodo := range g.Nodos {
		fmt.Printf("Nodo: %s\n", nodo.Nombre)
		if len(nodo.Entradas) > 0 {
			fmt.Printf("  Entradas: %v\n", nodo.Entradas)
		}
		if len(nodo.Salidas) > 0 {
			fmt.Println("  Salidas:")
			for _, salida := range nodo.Salidas {
				fmt.Printf("    -> %s (Costo: %.2f, Capacidad: %.2f)\n",
					salida.Destino, salida.Costo, salida.Capacidad)
			}
		}
		fmt.Println()
	}
}
