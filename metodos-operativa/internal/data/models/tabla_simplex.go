package models

type TablaSimplex struct {
	Ecuaciones []Ecuacion `json:"ecuaciones"`
	Iteracion  int        `json:"iteracion"`
}

type Ecuacion struct {
	Num int       `json:"num"`
	VB  string    `json:"vb"`
	LI  []Termino `json:"li"`
	LD  float64   `json:"ld"`
}

type Termino struct {
	C  float64 `json:"c"`
	VD string  `json:"vd"`
}
