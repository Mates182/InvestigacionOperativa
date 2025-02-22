package messages

// Genera el prompt para el análisis de sensibilidad en Programación Lineal
func PromptProgramacionLineal() string {
	return `Realiza un análisis de sensibilidad utilizando los siguientes datos. Incluye:
	- Listado de resultados y a qué parte del problema representan, si no hay enunciado del problema interpreta solo como datos.
	- Interpretación de los valores de las variables de decisión "xn" y su impacto en la función objetivo "z".
	- Determinación de si existen holguras en las restricciones y la razón detrás de ello ("sn").
	- Interpretación de los precios sombra ("yn") y cómo afectan las restricciones.
	- Evaluación de si es necesario aumentar o disminuir alguna restricción y la justificación de esta modificación.
	`
}

// Genera el prompt para el análisis de sensibilidad en problemas de Transporte
func PromptTransporte() string {
	return `Realiza un análisis de sensibilidad utilizando los siguientes datos del problema de transporte. Incluye:
	- Listado de asignaciones con su interpretación en términos de oferta y demanda, si no hay enunciado del problema interpreta solo como datos.
	- Análisis del costo total del transporte y su impacto en la eficiencia del sistema logístico.
	- Si existen orígenes o destinos ficticios, explica el por qué.
	- No incluyas tablas.
	`
}

// Genera el prompt para el análisis de sensibilidad en problemas de Grafos
func PromptGrafos() string {
	return `Realiza un análisis de sensibilidad utilizando los siguientes datos del problema de grafos. Incluye:
	- Si no hay enunciado del problema, interpreta solo como datos.
	- Si es Flujo Máximo, explicar el flujo máximo; si es distancia más corta, explicar la distancia mínima.
	- Listado de las rutas / ruta obtenidas, basado en los nodos del grafo en el JSON que contiene el grafo respuesta.
	- Análisis del costo total (cálculalo en base a las respuestas) y su impacto en la eficiencia.
	- No incluyas tablas.
	`
}
