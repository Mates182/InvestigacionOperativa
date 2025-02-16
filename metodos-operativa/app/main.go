package main

import (
	"fmt"
	"metodos-operativa/router"
)

func main() {
	fmt.Println("Api para Metodos de investigacion operativa iniciada!")
	r := router.SetRouter()

	r.Run("localhost:7000")
}
