package gemini

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func init() {
	// Cargar variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env") // Advertencia si el archivo no está presente
	}
}

func GenerarTexto(prompt string) string {
	ctx := context.Background()

	// Obtener la clave API desde las variables de entorno
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		fmt.Println("GEMINI_API_KEY no está definido en las variables de entorno")
	}

	// Crear un cliente para la API de Gemini
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err) // Terminar ejecución si hay un error al crear el cliente
	}
	defer client.Close() // Cerrar el cliente al finalizar la función

	// Seleccionar el modelo generativo específico
	model := client.GenerativeModel("gemini-1.5-flash")

	// Generar contenido a partir del prompt
	resp, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		log.Fatal(err) // Terminar ejecución si hay un error en la generación de contenido
	}

	return printResponse(resp) // Procesar y devolver la respuesta generada
}

func printResponse(resp *genai.GenerateContentResponse) string {
	response := ""

	// Recorrer las posibles respuestas generadas
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			// Extraer y concatenar los diferentes fragmentos de la respuesta
			for _, part := range cand.Content.Parts {
				response += fmt.Sprint(part)
			}
		}
	}

	return response
}
