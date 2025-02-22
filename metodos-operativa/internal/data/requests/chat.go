package requests

// PromptRequest representa una solicitud que contiene un mensaje o contenido de entrada
// utilizado para generar respuestas en modelos de lenguaje o procesamiento de datos.
type PromptRequest struct {
	Content string `json:"content"` // Texto de entrada proporcionado en la solicitud
}
