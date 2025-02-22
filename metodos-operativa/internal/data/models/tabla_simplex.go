package models

// TablaSimplex representa una iteración del método simplex,
// conteniendo el conjunto de ecuaciones y el número de iteración.
type TablaSimplex struct {
	Ecuaciones []Ecuacion `json:"ecuaciones"` // Lista de ecuaciones de la iteración
	Iteracion  int        `json:"iteracion"`  // Número de la iteración actual
}

// Ecuacion representa una restricción o ecuación dentro de la tabla simplex.
type Ecuacion struct {
	Num int       `json:"num"` // Número identificador de la ecuación
	VB  string    `json:"vb"`  // Variable básica en la ecuación
	LI  []Termino `json:"li"`  // Lado izquierdo de la ecuación (coeficientes y variables)
	LD  float64   `json:"ld"`  // Lado derecho de la ecuación (resultado)
}

// Termino representa un coeficiente con su variable correspondiente en una ecuación.
type Termino struct {
	C  float64 `json:"c"`  // Coeficiente del término
	VD string  `json:"vd"` // Nombre de la variable
}
