package messages

func PromptProgramacionLineal() string {
	return `Realiza un análisis de sensibilidad utilizando los siguientes datos. Incluye:
	- Listado de resultados y a que parte del problema representan, si no hay enunciado del problema interpreta solo como datos
	- Interpretación de los valores de las variables de decisión "xn" y su impacto en la función objetivo "z".
	- Determinación de si existen holguras en las restricciones y la razón detrás de ello ("sn").
	- Interpretación de los precios sombra ("yn") y cómo afectan las restricciones.
	- Evaluación de si es necesario aumentar o disminuir alguna restricción y la justificación de esta modificación.
	`
}
