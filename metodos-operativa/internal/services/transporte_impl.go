package services

import (
	"fmt"
	"metodos-operativa/internal/data/requests"
	"metodos-operativa/internal/data/responses"
	"metodos-operativa/pkg/transporte"
	"metodos-operativa/pkg/utils"
)

func (s *Services) Transporte(r requests.TransporteRequest) (int, responses.TransporteResponse) {
	// Obtener los datos de costos, orígenes y destinos de la solicitud
	datos := r.Costos
	origenes := r.Origenes
	destinos := r.Destinos

	// Calcular la diferencia entre la oferta total y la demanda total
	diferencia := func() float64 {
		var sumaOferta, sumaDemanda float64
		// Sumar la oferta de todos los orígenes
		for _, origen := range r.Origenes {
			sumaOferta += origen.Oferta
		}
		// Sumar la demanda de todos los destinos
		for _, destino := range r.Destinos {
			sumaDemanda += destino.Demanda
		}
		// Retornar la diferencia entre oferta y demanda
		return sumaOferta - sumaDemanda
	}()

	// Si la oferta es mayor que la demanda, agregar un destino ficticio
	if diferencia > 0 {
		datos = utils.Map(datos, func(costosOrigen []float64, i int) []float64 {
			// Agregar una columna extra con costo 0 para el destino ficticio
			return append(costosOrigen, 0)
		})
		destinos = append(destinos, requests.Destino{Demanda: diferencia, Destino: "Destino Ficticio"})
	} else if diferencia < 0 {
		// Si la demanda es mayor que la oferta, agregar un origen ficticio
		datos = append(datos, utils.Map(datos[0], func(costoDestino float64, i int) float64 {
			// Agregar una fila extra con costo 0 para el origen ficticio
			return 0
		}))
		origenes = append(origenes, requests.Origen{Oferta: -diferencia, Origen: "Origen Ficticio"})
	} else {
		// Si no hay diferencia, imprimir un mensaje
		fmt.Println("No hay diferencia")
	}

	// Aplicar el algoritmo de Vogel para encontrar la asignación inicial
	ofertas := utils.Map(origenes, func(origen requests.Origen, i int) float64 {
		// Extraer las ofertas de los orígenes
		return origen.Oferta
	})
	demandas := utils.Map(destinos, func(destino requests.Destino, i int) float64 {
		// Extraer las demandas de los destinos
		return destino.Demanda
	})

	// Crear una copia de los datos para evitar modificaciones directas en los originales
	datosTemp := utils.CopiarMatriz(datos)

	// Resolver el problema usando el algoritmo de Vogel
	asignacion := transporte.Vogel(datosTemp, ofertas, demandas)

	// Calcular el costo total de la asignación obtenida
	costoTotal := transporte.CalcularCostoTotal(datos, asignacion)

	// Crear un nuevo objeto de solicitud con los datos modificados
	newReq := requests.TransporteRequest{
		Costos:   datos,
		Origenes: origenes,
		Destinos: destinos,
	}

	// Formatear la asignación y el análisis de costos
	analisis := transporte.FormatearAsignacion(newReq, asignacion, costoTotal)

	// Retornar la solución óptima con la asignación, el costo total y el análisis
	return 0, responses.TransporteResponse{
		Message:    "Solución óptima encontrada",
		Request:    newReq,
		Asignacion: asignacion,
		CostoTotal: costoTotal,
		Analisis:   analisis,
	}
}
