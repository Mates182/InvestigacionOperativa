package requests

type TransporteRequest struct {
	Origenes  []Origen    `json:"origenes"`
	Destinos  []Destino   `json:"destinos"`
	Costos    [][]float64 `json:"costos"`
	Maximizar bool        `json:"maximizar"`
}

type Origen struct {
	Origen string  `json:"origen"`
	Oferta float64 `json:"oferta"`
}

type Destino struct {
	Destino string  `json:"destino"`
	Demanda float64 `json:"demanda"`
}
