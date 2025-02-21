package requests

type GrafosRequest struct {
	Conexiones  []Conexion `json:"conexiones"`
	Origen      string     `json:"origen"`
	Destino     string     `json:"destino"`
	EsRutaCorta bool       `json:"es_ruta_corta"`
}

type Conexion struct {
	Origen    string  `json:"origen"`
	Destino   string  `json:"destino"`
	Costo     float64 `json:"costo"`
	Capacidad float64 `json:"capacidad"`
	Distancia float64 `json:"distancia"`
}
