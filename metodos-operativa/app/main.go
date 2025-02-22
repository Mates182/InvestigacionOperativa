package main

import (
	"fmt"
	"metodos-operativa/router"
)

func main() {
	// Imprime un mensaje indicando que la API ha iniciado
	fmt.Println("Api para Metodos de investigacion operativa iniciada!")

	// Configura las rutas de la API utilizando el paquete 'router'
	r := router.SetRouter()

	// Inicia el servidor en 'localhost' en el puerto 7000
	r.Run("localhost:7000")
}
