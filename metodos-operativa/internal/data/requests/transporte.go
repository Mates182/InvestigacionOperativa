package requests

// TransporteRequest representa una solicitud para resolver un problema de transporte,
// incluyendo los orígenes, destinos, costos y el objetivo de maximizar o minimizar.
type TransporteRequest struct {
	Origenes  []Origen    `json:"origenes"`  // Lista de orígenes con su oferta
	Destinos  []Destino   `json:"destinos"`  // Lista de destinos con su demanda
	Costos    [][]float64 `json:"costos"`    // Matriz de costos de transporte entre orígenes y destinos
	Maximizar bool        `json:"maximizar"` // Indica si se debe maximizar el costo (true) o minimizarlo (false)
}

// Origen representa un origen con su nombre y la cantidad de oferta disponible.
type Origen struct {
	Origen string  `json:"origen"` // Nombre del origen
	Oferta float64 `json:"oferta"` // Oferta disponible en el origen
}

// Destino representa un destino con su nombre y la cantidad de demanda que debe ser satisfecha.
type Destino struct {
	Destino string  `json:"destino"` // Nombre del destino
	Demanda float64 `json:"demanda"` // Demanda en el destino
}
