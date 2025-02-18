"use client";
import React, { useState } from "react";
import ReactMarkdown from "react-markdown";

function TransporteForm() {
  const [numOrigenes, setNumOrigenes] = useState(3);
  const [numDestinos, setNumDestinos] = useState(3);
  const [procesando, setProcesando] = useState(false);
  const [asignacion, setAsignacion] = useState();
  const [request, setRequest] = useState();
  const [message, setMessage] = useState("");
  const [analisis, setAnalisis] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    setProcesando(true);

    try {
      // Obtener si es maximización o minimización
      const maximizar =
        document.getElementById("rest-objetivo").value === "true";

      // Obtener los nombres y ofertas de los orígenes
      const origenes = Array.from({ length: numOrigenes }).map((_, i) => ({
        origen:
          document.getElementById(`origen-${i + 1}`).value.trim() ||
          `O${i + 1}`,
        oferta: Number(document.getElementById(`oferta-${i + 1}`).value),
      }));

      // Obtener los nombres y demandas de los destinos
      const destinos = Array.from({ length: numDestinos }).map((_, i) => ({
        destino:
          document.getElementById(`destino-${i + 1}`).value.trim() ||
          `D${i + 1}`,
        demanda: Number(document.getElementById(`demanda-${i + 1}`).value),
      }));

      // Obtener la matriz de costos
      const costos = Array.from({ length: numOrigenes }).map((_, i) =>
        Array.from({ length: numDestinos }).map((_, j) =>
          Number(
            document.getElementById(`origen-${i + 1}-destino-${j + 1}`).value
          )
        )
      );

      // Construcción del JSON a enviar
      const body = { simplex: [], origenes, destinos, costos, maximizar };

      // Llamada a la API en localhost:7000/simplex
      const response = await fetch("http://localhost:7000/transporte", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(body),
      });

      if (!response.ok) throw new Error("Error al procesar la solicitud");

      const data = await response.json();
      const { asignacion, request, message, analisis } = data;
      setAsignacion(asignacion);
      setRequest(request);
      setMessage(message);
      const prompt = {
        content: `Enunciado: ${document.getElementById("content").value}
        Resultados: 
        ${analisis}`,
      };
      const responseGemini = await fetch(
        "http://localhost:7000/analisistransporte",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify(prompt),
        }
      );
      if (!responseGemini.ok) throw new Error("Error al procesar la solicitud");
      const { Message } = await responseGemini.json();
      setAnalisis(Message);
      console.log(data);
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
              <label htmlFor="numOrigenes" className="input-group-text">
                Número de Origenes:
              </label>
              <input
                type="number"
                className="form-control"
                id="numOrigenes"
                value={numOrigenes}
                onChange={(e) =>
                  setNumOrigenes(Math.max(1, Number(e.target.value)))
                }
                min="1"
              />
            </div>
          </div>
          <div className="col-auto">
            <div className="input-group">
              <label htmlFor="numDestinos" className="input-group-text">
                Número de Destinos:
              </label>
              <input
                type="number"
                className="form-control"
                id="numDestinos"
                value={numDestinos}
                onChange={(e) =>
                  setNumDestinos(Math.max(1, Number(e.target.value)))
                }
                min="1"
              />
            </div>
          </div>
        </div>
        <div className="input-group">
          <label htmlFor="rest-objetivo" className="input-group-text">
            Objetivo:
          </label>
          <select
            className="form-select"
            id="rest-objetivo"
            defaultValue="true"
          >
            {/*<option value="true">Maximizar</option>*/}
            <option value="false">Minimizar</option>
          </select>
        </div>
        <div className="card-header mt-0">Tabla de costos</div>
        <div className="p-2">
          <table className="table">
            <thead className="table-dark">
              <tr>
                <th>Origenes\Destinos</th>
                {Array.from({ length: numDestinos }).map((_, i) => (
                  <th key={`dest-${i}`}>
                    <input
                      type="text"
                      id={`destino-${i + 1}`}
                      placeholder={`Destino ${i + 1}`}
                      className="w-100"
                    />
                  </th>
                ))}
                <th>Oferta</th>
              </tr>
            </thead>
            <tbody>
              {Array.from({ length: numOrigenes }).map((_, i) => (
                <tr key={`origen-${i}`}>
                  <th className="table-dark">
                    <input
                      type="text"
                      id={`origen-${i + 1}`}
                      placeholder={`Origen ${i + 1}`}
                      className="w-100"
                    />
                  </th>
                  {Array.from({ length: numDestinos }).map((_, j) => (
                    <td key={`origen-${i}-destino-${j}`}>
                      <input
                        type="number"
                        id={`origen-${i + 1}-destino-${j + 1}`}
                        placeholder="Costo"
                        className="w-100 form-control"
                      />
                    </td>
                  ))}
                  <td className="table-active">
                    <input
                      type="number"
                      id={`oferta-${i + 1}`}
                      placeholder="Oferta"
                      className="w-100 form-control"
                    />
                  </td>
                </tr>
              ))}
              <tr>
                <th className="table-dark">Demanda</th>
                {Array.from({ length: numDestinos }).map((_, i) => (
                  <td key={`demanda-${i}`} className="table-active">
                    <input
                      type="number"
                      id={`demanda-${i + 1}`}
                      placeholder="Demanda"
                      className="w-100 form-control"
                    />
                  </td>
                ))}
                <td className="table-active"></td>
              </tr>
            </tbody>
          </table>
        </div>

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

        <button type="submit" disabled={procesando}>
          <h5 className="mt-2">{procesando ? "Procesando..." : "Calcular"}</h5>
        </button>
      </form>
      {message == "Solución óptima encontrada" && (
        <>
          <div className="card border-success mb-3 ms-0 me-0">
            <div className="card-header mt-0">Solución óptima</div>
            <div className="p-2">
              <table className="table">
                <thead className="table-dark">
                  <tr>
                    <th>Origenes\Destinos</th>
                    {request.destinos.map((destino, i) => (
                      <th key={`dest-${i}`}>{destino.destino}</th>
                    ))}
                    <th>Oferta</th>
                  </tr>
                </thead>
                <tbody>
                  {request.origenes.map((origen, i) => (
                    <tr key={`origen-${i}`}>
                      <th className="table-dark">{origen.origen}</th>
                      {request.destinos.map((_, j) => (
                        <td key={`origen-${i}-destino-${j}`}>
                          {
                            <ReactMarkdown className="card-text">{`**Costo:** ${request.costos[i][j]}\n**Asignación**: ${asignacion[i][j]}`}</ReactMarkdown>
                          }
                        </td>
                      ))}
                      <td className="table-active">
                        <ReactMarkdown>{`**${origen.oferta}**`}</ReactMarkdown>
                      </td>
                    </tr>
                  ))}
                  <tr>
                    <th className="table-dark">Demanda</th>
                    {request.destinos.map((destino, i) => (
                      <td key={`demanda-${i}`} className="table-active">
                        <ReactMarkdown>{`**${destino.demanda}**`}</ReactMarkdown>
                      </td>
                    ))}
                    <td className="table-active"></td>
                  </tr>
                </tbody>
              </table>
            </div>
          </div>
          <div className="card border-info mb-3">
            <div className="card-header">Interpretación de Resultados</div>
            <div className="card-body">
              <h5 className="card-title">
                Interpretación generada por Gemini:
              </h5>
              <ReactMarkdown className="card-text">{analisis}</ReactMarkdown>
            </div>
          </div>
        </>
      )}
    </div>
  );
}

export default TransporteForm;
