"use client";
import React, { useState } from "react";
import Simplex from "@/modules/Simplex";
import DosFases from "@/modules/DosFases";
import ReactMarkdown from "react-markdown";

function ProgramacionLinealForm() {
  const [numVariables, setNumVariables] = useState(2);
  const [numRestricciones, setNumRestricciones] = useState(2);
  const [procesando, setProcesando] = useState(false);
  const [resultado, setResultado] = useState({});
  const [metodo, setMetodo] = useState("");
  const [modelo, setModelo] = useState([]);
  const [analisis, setAnalisis] = useState("");
  const [analisisSensibilidad, setAnalisisSensibilidad] = useState(false);

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
      const restricciones = Array.from({ length: numRestricciones }).map(
        (_, i) => {
          const li = Array.from({ length: numVariables }).map((_, j) => ({
            C: Number(document.getElementById(`rest-${i}-coef-${j}`).value),
            VD: `x${j + 1}`,
          }));
          const operador = document.getElementById(`rest-${i}-operador`).value;
          const ld = Number(document.getElementById(`rest-${i}-ld`).value);
          return { li, operador, ld };
        }
      );

      // Obtener si es maximización o minimización
      const maximizar =
        document.getElementById("rest-objetivo").value === "true";

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

      const { resolucion, metodo, modelo, respuestas } = await response.json();

      setResultado(resolucion);
      setMetodo(metodo);
      setModelo(modelo);

      if (analisisSensibilidad) {
        const prompt = {
          content: `Enunciado: ${document.getElementById("content").value}
        Funcion Objetivo: ${modelo[0]}
        Restricciones: ${modelo[1]}
        Respuestas: ${respuestas}`,
        };
        const responseGemini = await fetch("http://localhost:7000/analisispl", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(prompt),
        });
        if (!responseGemini.ok)
          throw new Error("Error al procesar la solicitud");
        const { Message } = await responseGemini.json();
        setAnalisis(Message);
      }
    } catch (error) {
      console.error("Error al enviar los datos:", error);
      alert("Hubo un error al enviar los datos");
    } finally {
      setProcesando(false);
    }
  };

  return (
    <>
      <form
        onSubmit={handleSubmit}
        className="row g-3 card border-dark mb-3 ms-0 me-0"
      >
        <div className="card-header mt-0">Planteamiento del problema</div>
        {/* Input para modificar el número de variables */}
        <div className="row gy-2 gx-3 align-items-center">
          <div className="col-auto">
            <div className="input-group">
              <label htmlFor="numVariables" className="input-group-text">
                Número de Variables:
              </label>
              <input
                type="number"
                step="any"
                className="form-control"
                id="numVariables"
                value={numVariables}
                onChange={(e) =>
                  setNumVariables(Math.max(1, Number(e.target.value)))
                }
                min="1"
              />
            </div>
          </div>
          <div className="col-auto">
            {/* Input para modificar el número de restricciones */}
            <div className="input-group">
              <label htmlFor="numRestricciones" className="input-group-text">
                Número de Restricciones:
              </label>
              <input
                type="number"
                step="any"
                className="form-control"
                id="numRestricciones"
                value={numRestricciones}
                onChange={(e) =>
                  setNumRestricciones(Math.max(1, Number(e.target.value)))
                }
                min="1"
              />
            </div>
          </div>
        </div>

        <h4 className="card-title ">Función Objetivo</h4>
        <div className="input-group">
          <label htmlFor="rest-objetivo" className="input-group-text">
            Objetivo:
          </label>
          <select
            className="form-select"
            id="rest-objetivo"
            defaultValue="true"
          >
            <option value="true">Maximizar</option>
            <option value="false">Minimizar</option>
          </select>
        </div>
        <div className="row gy-2 gx-3 align-items-center">
          <h5 className="mb-1 ms-2 col-auto input-group-text">Z</h5>
          <h4 className="mb-1 ms-2 col-auto">=</h4>
          {Array.from({ length: numVariables }).map((_, i) => (
            <div key={`fo-${i}`} className="col-auto">
              <div className="input-group ">
                <input
                  className="form-control"
                  type="number"
                  step="any"
                  id={`fo-coef-${i}`}
                  placeholder="Coeficiente"
                  required
                />
                <span className="input-group-text">x{i + 1}</span>
                {i < numVariables - 1 && <h4 className="ms-3">+</h4>}
              </div>
            </div>
          ))}
        </div>

        <h4 className="card-title ">Restricciones</h4>
        {Array.from({ length: numRestricciones }).map((_, i) => (
          <div
            key={`rest-${i}`}
            style={{
              display: "grid",
              alignItems: "center",
              gap: "10px",
              marginBottom: "10px",
            }}
          >
            <h5>Restricción {i + 1}:</h5>
            <div className="row gy-2 gx-3 align-items-center">
              {Array.from({ length: numVariables }).map((_, j) => (
                <div key={`rest-${i}-var-${j}`} className="col-auto">
                  <div className="input-group ">
                    <input
                      type="number"
                      step="any"
                      className="form-control"
                      id={`rest-${i}-coef-${j}`}
                      placeholder="Coeficiente"
                      defaultValue={0}
                      required
                    />
                    <span className="input-group-text">x{j + 1}</span>
                    {j < numVariables - 1 && <h4 className="ms-3">+</h4>}
                  </div>
                </div>
              ))}
              <div className="col-auto">
                <div className="input-group ">
                  <select
                    className="form-select"
                    id={`rest-${i}-operador`}
                    defaultValue="≤"
                  >
                    <option value="≤">≤</option>
                    <option value="=">=</option>
                    <option value="≥">≥</option>
                  </select>
                </div>
              </div>
              <div className="col-auto">
                <div className="input-group ">
                  <input
                    className="form-control"
                    type="number"
                    step="any"
                    id={`rest-${i}-ld`}
                    placeholder="LD"
                    defaultValue={0}
                    required
                  />
                </div>
              </div>
            </div>
          </div>
        ))}

        <div className="card-header mt-0">Análisis de sensiblilidad</div>
        <div className="input-group">
          <div className="form-check">
            <input
              className="form-check-input"
              type="checkbox"
              value=""
              id="flexCheckDefault"
              onChange={(e) => {
                setAnalisisSensibilidad(e.target.checked);
              }}
            />
            <label className="form-check-label" htmlFor="flexCheckDefault">
              Incluir análisis de sensibilidad
            </label>
          </div>
        </div>
        {analisisSensibilidad && (
          <div className="input-group">
            <span className="input-group-text">
              Enunciado del <br /> problema:
            </span>
            <textarea
              placeholder="Ingrese el enunciado del problema (opcional)"
              className="form-control"
              aria-label="With textarea"
              id="content"
            ></textarea>
          </div>
        )}

        <button type="submit" disabled={procesando}>
          <h5 className="mt-2">{procesando ? "Procesando..." : "Calcular"}</h5>
        </button>
      </form>
      {(metodo == "simplex" || metodo == "dos fases") && (
        <div>
          <div className="card border-success mb-3">
            <div className="card-header text-success">Modelo Matemático</div>
            <div className="card-body">
              <h5 className="card-title text-success">Función Objetivo</h5>
              <p className="card-text">{modelo[0]}</p>
              <h5 className="card-title text-success">Sujeto a:</h5>
              <p className="card-text">{modelo[1]}</p>
            </div>
          </div>
        </div>
      )}
      {metodo == "simplex" && <Simplex resoluciones={resultado} />}
      {metodo == "dos fases" && <DosFases resoluciones={resultado} />}
      {(metodo == "simplex" || metodo == "dos fases") &&
        analisisSensibilidad && (
          <div className="card border-info mb-3">
            <div className="card-header">Interpretación de Resultados</div>
            <div className="card-body">
              <h5 className="card-title">
                Interpretación generada por Gemini:
              </h5>
              <ReactMarkdown className="card-text">{analisis}</ReactMarkdown>
            </div>
          </div>
        )}
    </>
  );
}

export default ProgramacionLinealForm;
