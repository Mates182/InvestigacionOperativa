package services

// Services es una estructura que representa un conjunto de servicios.
type Services struct {
}

// NewServices crea e inicializa una nueva instancia de Services.
// Devuelve un puntero a la instancia de Services.
func NewServices() Service {
	return &Services{}
}
