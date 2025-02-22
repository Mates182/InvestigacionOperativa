package requests

// GrafosRequest representa una solicitud para resolver un problema de grafos,
// incluyendo las conexiones entre nodos y los parámetros de la consulta.
type GrafosRequest struct {
	Conexiones  []Conexion `json:"conexiones"`    // Lista de conexiones entre nodos en el grafo
	Origen      string     `json:"origen"`        // Nodo de inicio para el análisis
	Destino     string     `json:"destino"`       // Nodo de destino para el análisis
	EsRutaCorta bool       `json:"es_ruta_corta"` // Indica si se debe calcular la ruta más corta (true) o flujo máximo (false)
}

// Conexion representa una arista en el grafo con atributos relevantes.
type Conexion struct {
	Origen    string  `json:"origen"`    // Nodo de origen de la conexión
	Destino   string  `json:"destino"`   // Nodo de destino de la conexión
	Costo     float64 `json:"costo"`     // Costo asociado a la conexión
	Capacidad float64 `json:"capacidad"` // Capacidad máxima de flujo en la conexión
	Distancia float64 `json:"distancia"` // Distancia entre los nodos conectados
}
