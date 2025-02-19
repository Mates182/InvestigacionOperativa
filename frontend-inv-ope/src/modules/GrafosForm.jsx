"use client";
import React, { useState } from "react";
import ReactMarkdown from "react-markdown";

function GrafosForm() {
  const [numNodos, setNumNodos] = useState(2);
  const [nodos, setNodos] = useState(["Nodo1", "Nodo2"]);
  const [numConexiones, setNumConexiones] = useState(1);
  const [procesando, setProcesando] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setProcesando(true);

    try {
    } catch (error) {
      console.error("Error al enviar los datos:", error);
      alert("Hubo un error al enviar los datos");
    } finally {
      setProcesando(false);
    }
  };

  return (
    <div>
      <form
        onSubmit={handleSubmit}
        className="row g-3 card border-dark mb-3 ms-0 me-0"
      >
        <div className="card-header mt-0">Planteamiento del problema</div>
        <div className="row gy-2 gx-3 align-items-center">
          <div className="col-auto">
            <div className="input-group">
              <label htmlFor="numNodos" className="input-group-text">
                Número de Nodos:
              </label>
              <input
                type="number"
                step="any"
                className="form-control"
                id="numNodos"
                value={numNodos}
                onChange={(e) => {
                  const newNumNodos = Math.max(2, Number(e.target.value));
                  setNumNodos(newNumNodos);
                  setNodos((prevNodos) => {
                    const newNodos = [...prevNodos];
                    while (newNodos.length < newNumNodos)
                      newNodos.push(`Nodo${newNodos.length + 1}`);
                    return newNodos.slice(0, newNumNodos);
                  });
                }}
                min="2"
              />
            </div>
          </div>
          <div className="col-auto">
            <div className="input-group">
              <label htmlFor="numConexiones" className="input-group-text">
                Número de Conexiones:
              </label>
              <input
                type="number"
                step="any"
                className="form-control"
                id="numConexiones"
                value={numConexiones}
                onChange={(e) =>
                  setNumConexiones(Math.max(1, Number(e.target.value)))
                }
                min="1"
              />
            </div>
          </div>
        </div>
        <h4 className="card-title ">Nodos</h4>
        {Array.from({ length: numNodos }).map((_, i) => (
          <div key={`nodo-${i}`}>
            <div className="input-group ">
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Nodo {i + 1}:
              </label>
              <input
                type="text"
                className="form-control"
                placeholder="Nombre"
                required
                value={nodos[i]}
                onChange={(e) => {
                  const nuevosNodos = [...nodos];
                  nuevosNodos[i] = e.target.value;
                  setNodos(nuevosNodos);
                }}
              />
            </div>
          </div>
        ))}
        <h4 className="card-title ">Conexiones</h4>
        {Array.from({ length: numConexiones }).map((_, i) => (
          <div key={`conexion-${i}`}>
            <div className="input-group ">
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Nodo Origen:
              </label>
              <select
                className="form-select"
                id={`conexion-${i}-origen`}
                defaultValue={nodos[0] || ""}
              >
                {nodos.map((nodo, index) => (
                  <option key={index} value={nodo}>
                    {nodo}
                  </option>
                ))}
              </select>
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Nodo Destino:
              </label>
              <select
                className="form-select"
                id={`conexion-${i}-destino`}
                defaultValue={nodos[1] || ""}
              >
                {nodos.map((nodo, index) => (
                  <option key={index} value={nodo}>
                    {nodo}
                  </option>
                ))}
              </select>
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Costo:
              </label>
              <input
                type="number"
                step="any"
                className="form-control"
                placeholder="Costo"
                required
              />
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Capacidad:
              </label>
              <input
                type="number"
                step="any"
                className="form-control"
                placeholder="Capacidad"
                required
              />
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Distancia:
              </label>
              <input
                type="number"
                step="any"
                className="form-control"
                placeholder="Distancia"
                required
              />
            </div>
          </div>
        ))}
      </form>
      <div>
        {/* Vista previa del nodo / grafico */}
      </div>
    </div>
  );
}

export default GrafosForm;
