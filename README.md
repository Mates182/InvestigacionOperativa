# Proyecto Final - Investigación Operativa

Este es el proyecto final de la materia de Investigación Operativa en la Universidad Central del Ecuador, Facultad de Ingeniería y Ciencias Aplicadas, Carrera de Sistemas de Información.
---

## Integrantes

- Mateo Bernardo Pillajo Lopez
- Christian Rolando Tapia Diaz
- Lenin Sebastián Serrano Montúfar

### Descripción General

El proyecto consiste en desarrollar una aplicación computacional que resuelva problemas de **Programación Lineal**, **Transporte** y **Redes**, integrando **Inteligencia Artificial (IA)** para realizar el análisis de sensibilidad.

Además, se planteó un problema empresarial real e integral que abarca los temas previamente mencionados, permitiendo tomar decisiones asertivas basadas en los resultados obtenidos.

### Objetivos Específicos

- Desarrollar un programa capaz de resolver problemas de optimización en los siguientes ámbitos:
  - Programación Lineal
  - Transporte
  - Redes

- Integrar con **Inteligencia Artificial (IA)** para realizar análisis de sensibilidad.

---

## Requisitos

Para ejecutar este proyecto, se necesita tener instalado:

- **Node.js v22.11.0 o superior** (para el frontend)
- **Go 1.23.4 o superior** (para el backend)

---

## Instrucciones para Correr el Proyecto

1. **Clonar el repositorio:**
   
   Clonar el repositorio desde GitHub:

   ```bash
   git clone https://github.com/Mates182/InvestigacionOperativa
   ```

2. **Instalar dependencias del frontend:**

   Ir a la carpeta del frontend (`frontend-inv-ope`) y ejecutar el siguiente comando para instalar las dependencias:

   ```bash
   cd frontend-inv-ope
   npm install
   ```

3. **Correr el frontend:**

   Después de instalar las dependencias, ejecutar el siguiente comando para correr el proyecto en modo desarrollo:

   ```bash
   npm run dev
   ```

   El frontend estará disponible en `http://localhost:3000`.

4. **Correr el backend:**

   Ir a la carpeta del backend (`metodos-operativa`) 
   ```bash
   cd metodos-operativa
   ```
   y crear un archivo `.env` con una clave de API para Gemini, por ejemplo:

   ```bash
   GEMINI_API_KEY=APIKEYEJEMPLO123
   ```

   Luego, ejecutar el siguiente comando para correr el backend:

   ```bash
   go run ./app
   ```

---

## Estructura del Proyecto

El proyecto está dividido en dos partes principales: el frontend y el backend.

- **frontend-inv-ope**: Contiene la aplicación construida en React.
- **metodos-operativa**: Contiene la API construida en Go.




## Backend - Métodos de Investigación Operativa

```bash
metodos-operativa/
├── app/
│   └── main.go  # Punto de entrada de la aplicación, donde se inicia el servidor web.
├── config/
│   └── cors/
│       └── cors.go  # Configuración de CORS para permitir solicitudes desde orígenes específicos.
├── internal/
│   ├── controllers/
│   │   ├── grafos_controller.go  # Controlador para manejar las solicitudes relacionadas con grafos.
│   │   ├── programacion_lineal_controller.go  # Controlador para manejar las solicitudes de programación lineal.
│   │   └── transporte_controller.go  # Controlador para manejar las solicitudes de problemas de transporte.
│   ├── data/
│   │   ├── messages/
│   │   │   └── messages.go  # Contiene los mensajes que se usarán en los análisis (ej. prompt para Gemini).
│   │   ├── models/
│   │   │   ├── grafos.go  # Modelos relacionados con grafos, como flujos máximos o rutas más cortas.
│   │   │   └── tabla_simplex.go  # Modelos para representar la tabla de Simplex en programación lineal.
│   │   ├── requests/
│   │   │   ├── chat.go  # Estructuras para las solicitudes de chat (posiblemente para interacción).
│   │   │   ├── grafos_request.go  # Estructura de datos para solicitudes relacionadas con grafos.
│   │   │   ├── programacion_lineal_request.go  # Estructura de datos para solicitudes de programación lineal.
│   │   │   └── transporte.go  # Estructura de datos para solicitudes relacionadas con problemas de transporte.
│   │   └── responses/
│   │       ├── grafos_responses.go  # Respuestas para las solicitudes de grafos.
│   │       ├── simplex_responses.go  # Respuestas para las solicitudes de programación lineal.
│   │       └── transporte_responses.go  # Respuestas para las solicitudes de transporte.
│   ├── gemini/
│   │   └── gemini.go  # Funciones que interactúan con Gemini para generar texto a partir de los prompts.
│   └── services/
│       ├── dos_fases_impl.go  # Implementación del método de dos fases para programación lineal.
│       ├── services.go  # Interfaz general para los servicios de la aplicación.
│       ├── services_impl.go  # Implementación de los servicios generales.
│       ├── simplex_impl.go  # Implementación específica para resolver problemas de programación lineal con Simplex.
│       └── transporte_impl.go  # Implementación para resolver problemas de transporte.
├── pkg/
│   ├── grafos/
│   │   ├── flujo_maximo.go  # Algoritmos para resolver problemas de flujo máximo en grafos.
│   │   └── ruta_mas_corta.go  # Algoritmos para encontrar la ruta más corta en un grafo.
│   ├── programacion_lineal/
│   │   └── programacion_lineal.go  # Implementación de algoritmos para resolver problemas de programación lineal.
│   ├── transporte/
│   │   └── transporte.go  # Implementación de algoritmos para resolver problemas de transporte.
│   └── utils/
│       ├── convertir.go  # Funciones auxiliares para convertir y dar formato a datos de programación lineal y transporte.
│       └── utils.go  # Funciones generales auxiliares para la aplicación.
├── router/
│   └── router.go  # Configura las rutas del servidor web y los controladores asociados.
├── test/
│   └── main.go  # Archivos de pruebas (unitarias o de integración).
├── .env  # Archivo de configuración para variables de entorno (como claves de API, base de datos, etc.).
├── .gitignore  # Archivos y carpetas que deben ser ignorados por Git.
├── go.mod  # Archivo de módulos Go para la gestión de dependencias.
└── go.sum  # Archivo que contiene el hash de las dependencias para garantizar su integridad.
```

## Frontend

```bash
frontend-inv-ope/
├── public/ # Iconos de la aplicación
│   ├── file.svg 
│   ├── globe.svg 
│   ├── next.svg  
│   ├── vercel.svg 
│   └── window.svg  
├── src/
│   ├── app/
│   │   ├── grafos/
│   │   │   └── page.jsx  # Componente para manejar las vistas de grafos en la interfaz.
│   │   ├── programacion-lineal/
│   │   │   └── page.jsx  # Componente para manejar las vistas de programación lineal en la interfaz.
│   │   ├── transporte/
│   │   │   └── page.jsx  # Componente para manejar las vistas de transporte en la interfaz.
│   │   ├── favicon.ico  # Icono que aparece en la pestaña del navegador.
│   │   ├── globals.css  # Estilos globales que se aplican a toda la aplicación.
│   │   ├── layout.jsx  # Componente que define el layout general de la aplicación.
│   │   ├── page.jsx  # Página principal que maneja la vista general.
│   │   └── page.module.css  # Estilos específicos para la página principal.
│   └── modules/
│       ├── DosFases.jsx  # Componente que maneja el formulario y los datos para el método de dos fases.
│       ├── GrafosForm.jsx  # Componente que maneja el formulario para problemas de grafos.
│       ├── ProgramacionLinealForm.jsx  # Componente que maneja el formulario para programación lineal.
│       ├── Simplex.jsx  # Componente que gestiona la lógica y vista para el algoritmo Simplex.
│       └── TransporteForm.jsx  # Componente que maneja el formulario para problemas de transporte.
├── .gitignore  # Archivos y carpetas que deben ser ignorados por Git.
├── README.md  # Documento que contiene la descripción del proyecto y la guía de uso.
├── eslint.config.mjs  # Configuración de ESLint para el análisis de código.
├── jsconfig.json  # Configuración de JavaScript, útil para autocompletado y rutas en el proyecto.
├── next.config.mjs  # Configuración de Next.js para personalizar el comportamiento del proyecto.
├── package-lock.json  # Archivo que bloquea las versiones exactas de las dependencias.
└── package.json  # Archivo de configuración de dependencias y scripts de la aplicación.
```