"use client";
import React, { useState } from "react";
import Simplex from '@/modules/Simplex'
import DosFases from '@/modules/DosFases'

function ProgramacionLinealForm() {
  const [numVariables, setNumVariables] = useState(2);
  const [numRestricciones, setNumRestricciones] = useState(2);
  const [procesando, setProcesando] = useState(false);
  const [resultado, setResultado] = useState({});
  const [metodo, setMetodo] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setProcesando(true);

    try {
      // Obtener datos de la función objetivo
      const fo = Array.from({ length: numVariables }).map((_, i) => ({
        C: Number(document.getElementById(`fo-coef-${i}`).value),
        VD: `x${i + 1}`,
      }));

      // Obtener restricciones
      const restricciones = Array.from({ length: numRestricciones }).map((_, i) => {
        const li = Array.from({ length: numVariables }).map((_, j) => ({
          C: Number(document.getElementById(`rest-${i}-coef-${j}`).value),
          VD: `x${j + 1}`,
        }));
        const operador = document.getElementById(`rest-${i}-operador`).value;
        const ld = Number(document.getElementById(`rest-${i}-ld`).value);
        return { li, operador, ld };
      });

      // Obtener si es maximización o minimización
      const maximizar = document.getElementById("rest-objetivo").value === "true";

      // Estructurar datos para enviar a la API
      const body = { fo, restricciones, maximizar };

      // Llamada a la API en localhost:7000
      const response = await fetch("http://localhost:7000/simplex", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });

      if (!response.ok) throw new Error("Error al procesar la solicitud");

      alert("¡Datos enviados con éxito!");
      const {resolucion, metodo} = await response.json();
      setResultado(resolucion)
      setMetodo(metodo)
    } catch (error) {
      console.error("Error al enviar los datos:", error);
      alert("Hubo un error al enviar los datos");
    } finally {
      setProcesando(false);
    }
  };

  return (
    <>
    <form onSubmit={handleSubmit}>
      <h1>Formulario de Programación Lineal</h1>

      {/* Input para modificar el número de variables */}
      <label htmlFor="numVariables">Número de Variables:</label>
      <input
        type="number"
        id="numVariables"
        value={numVariables}
        onChange={(e) => setNumVariables(Math.max(1, Number(e.target.value)))}
        min="1"
      />

      {/* Input para modificar el número de restricciones */}
      <label htmlFor="numRestricciones">Número de Restricciones:</label>
      <input
        type="number"
        id="numRestricciones"
        value={numRestricciones}
        onChange={(e) => setNumRestricciones(Math.max(1, Number(e.target.value)))}
        min="1"
      />

      <h2>Función Objetivo</h2>
      <div style={{ display: "grid", alignItems: "center", gap: "10px" }}>
        {Array.from({ length: numVariables }).map((_, i) => (
          <div key={`fo-${i}`} style={{ display: "flex", alignItems: "center", gap: "5px" }}>
            <input type="number" id={`fo-coef-${i}`} placeholder="Coeficiente" required />
            <span>x{i + 1}</span>
            {i < numVariables - 1 && <span>+</span>}
          </div>
        ))}
      </div>
      <select id="rest-objetivo" defaultValue="true">
        <option value="true">Maximizar</option>
        <option value="false">Minimizar</option>
      </select>

      <h2>Restricciones</h2>
      {Array.from({ length: numRestricciones }).map((_, i) => (
        <div key={`rest-${i}`} style={{ display: "grid", alignItems: "center", gap: "10px", marginBottom: "10px" }}>
          {Array.from({ length: numVariables }).map((_, j) => (
            <div key={`rest-${i}-var-${j}`} style={{ display: "flex", alignItems: "center", gap: "5px" }}>
              <input type="number" id={`rest-${i}-coef-${j}`} placeholder="Coeficiente" required />
              <span>x{j + 1}</span>
              {j < numVariables - 1 && <span>+</span>}
            </div>
          ))}
          <select id={`rest-${i}-operador`} defaultValue="≤">
            <option value="≤">≤</option>
            <option value="=">=</option>
            <option value="≥">≥</option>
          </select>
          <input type="number" id={`rest-${i}-ld`} placeholder="LD" required />
        </div>
      ))}

      <button type="submit" disabled={procesando}>
        {procesando ? "Procesando..." : "Enviar"}
      </button>
    </form>
    {metodo == "simplex" && (<Simplex resoluciones={resultado}/>)}
    {metodo == "dos fases" && (<DosFases resoluciones={resultado}/>)}

    </>
  );
}

export default ProgramacionLinealForm;
