"use client";
import React, { useState, useCallback } from "react";
import ReactFlow, {
  addEdge,
  MiniMap,
  Controls,
  Background,
  useNodesState,
  useEdgesState,
} from "reactflow";
import { MarkerType } from "@xyflow/react";
import "reactflow/dist/style.css";

function GrafosForm() {
  const [numNodos, setNumNodos] = useState(2);
  const [nodos, setNodos] = useState([
    { id: "1", data: { label: "Nodo1" }, position: { x: 0, y: 0 } },
    { id: "2", data: { label: "Nodo2" }, position: { x: 100, y: 100 } },
  ]);
  const [numConexiones, setNumConexiones] = useState(1);
  const [procesando, setProcesando] = useState(false);

  const handleSubmit = async (e) => {
    e.preventDefault();
    setProcesando(true);

    // Construir el objeto de nodos
    const conexionesObj = conexiones.map((conexion) => {
      const temp = {
        origen: nodos[parseInt(conexion.edge.source) - 1].data.label,
        destino: nodos[parseInt(conexion.edge.target) - 1].data.label,
        costo: parseFloat(conexion.costo),
        capacidad: parseFloat(conexion.capacidad),
        distancia: parseFloat(conexion.distancia),
      };
      return temp;
    });

    // Crear el cuerpo de la solicitud
    const requestBody = {
      conexiones: conexionesObj,
      origen: nodos[0].data.label,
      destino: nodos[nodos.length-1].data.label,
    };

    try {
      // Enviar los datos al backend
      const response = await fetch("http://localhost:7000/grafos", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(requestBody),
      });

      if (!response.ok) {
        throw new Error("Error en la respuesta del servidor");
      }

      const data = await response.json();
      console.log("Respuesta del servidor:", data);
      // Manejar la respuesta del servidor según sea necesario
    } catch (error) {
      console.error("Error al enviar los datos:", error);
      alert("Hubo un error al enviar los datos");
    } finally {
      setProcesando(false);
    }
  };

  const [conexiones, setConexiones] = useState([
    {
      edge: {
        id: "e1",
        source: "1",
        target: "2",
        label: "Conexión \n1",
        markerEnd: {
          type: MarkerType.ArrowClosed, // Tipo de marcador de flecha cerrada
        },
      },
      costo: 0,
      capacidad: 0,
      distancia: 0,
    },
  ]);

  const [nodes, setNodes, onNodesChange] = useNodesState(nodos);
  const [edges, setEdges, onEdgesChange] = useEdgesState(
    conexiones.map((conexion) => conexion.edge)
  );

  const onConnect = useCallback(
    (params) => setEdges((eds) => addEdge(params, eds)),
    [setEdges]
  );

  const handleNodoChange = (index, value) => {
    const updatedNodos = [...nodos];
    updatedNodos[index] = {
      ...updatedNodos[index],
      data: { label: value },
      id: (index + 1).toString(),
    };
    setNodos(updatedNodos);
    setNodes(updatedNodos);
  };

  const handleAddNodo = (e) => {
    const newNumNodos = Math.max(2, Number(e.target.value));
    setNumNodos(newNumNodos);
    if (nodos.length >= newNumNodos) {
      setNodos((nds) => nds.slice(0, newNumNodos));
      setNodes((nds) => nds.slice(0, newNumNodos));
      return;
    }
    const newNode = {
      id: (nodos.length + 1).toString(),
      data: { label: `Nodo${nodos.length + 1}` },
      position: { x: Math.random() * 400, y: Math.random() * 400 },
    };
    setNodos((nds) => [...nds, newNode]);
    setNodes((nds) => [...nds, newNode]);
  };

  const handleAddConexion = (e) => {
    const newNumConexiones = Math.max(1, Number(e.target.value));
    setNumConexiones(newNumConexiones);

    if (conexiones.length >= newNumConexiones) {
      setConexiones((eds) => eds.slice(0, newNumConexiones));
      setEdges((eds) => eds.slice(0, newNumConexiones));
      return;
    }
    const newEdge = {
      id: `e${numConexiones + 1}`,
      source: "1",
      target: "1",
      label: `Conexión ${numConexiones + 1}`,
      markerEnd: {
        type: MarkerType.ArrowClosed, // Tipo de marcador de flecha cerrada
      },
    };
    let tempConexiones = [
      ...conexiones,
      { edge: newEdge, costo: 0, capacidad: 0, distancia: 0 },
    ];
    let tempEdges = [...edges, newEdge];
    console.log(tempConexiones);
    console.log(tempEdges);
    setConexiones(tempConexiones);
    setEdges(tempEdges);
  };

  const handleConexionChange = (e, i) => {
    let conexionesTemp = [...conexiones];
    console.log(conexionesTemp);
    console.log(i);
    console.log(e.target.name);
    if (e.target.name === "costo") {
      conexionesTemp[i].costo = e.target.value;
    } else if (e.target.name === "capacidad") {
      conexionesTemp[i].capacidad = e.target.value;
    } else if (e.target.name === "distancia") {
      conexionesTemp[i].distancia = e.target.value;
    } else if (e.target.name === "origen") {
      conexionesTemp[i].edge.source = `${e.target.value}`;
      console.log(e.target.value);
    } else if (e.target.name === "destino") {
      conexionesTemp[i].edge.target = `${e.target.value}`;
    }
    conexionesTemp[i].edge.label = `C-${i + 1}\nCosto: ${
      conexionesTemp[i].costo
    }\nCapacidad: ${conexionesTemp[i].capacidad}\nDistancia: ${
      conexionesTemp[i].distancia
    }`;
    setConexiones(conexionesTemp);
    setEdges(conexionesTemp.map((conexion) => conexion.edge));
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
                onChange={handleAddNodo}
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
                onChange={handleAddConexion}
                min="1"
              />
            </div>
          </div>
        </div>
        <h4 className="card-title ">Nodos</h4>
        {nodos.map((nodo, i) => (
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
                value={nodo.data.label}
                onChange={(e) => handleNodoChange(i, e.target.value)}
              />
            </div>
          </div>
        ))}
        <h4 className="card-title ">Conexiones</h4>
        {Array.from({ length: numConexiones }).map((_, i) => (
          <div key={`conexion-${i}`}>
            <div
              className="input-group "
              onChange={(e) => handleConexionChange(e, i)}
            >
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Nodo Origen:
              </label>
              <select
                className="form-select"
                id={`conexion-${i}-origen`}
                defaultValue={nodos[0].id || ""}
                name="origen"
              >
                {nodos.map((nodo, index) => (
                  <option key={index} value={nodo.id}>
                    {nodo.data.label}
                  </option>
                ))}
              </select>
              <label htmlFor={`nodo-${i}`} className="input-group-text">
                Nodo Destino:
              </label>
              <select
                className="form-select"
                id={`conexion-${i}-destino`}
                defaultValue={nodos[1].id || ""}
                name="destino"
              >
                {nodos.map((nodo, index) => (
                  <option key={index} value={nodo.id}>
                    {nodo.data.label}
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
                name="costo"
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
                name="capacidad"
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
                name="distancia"
              />
            </div>
          </div>
        ))}
        <button type="submit" disabled={procesando}>
          <h5 className="mt-2">{procesando ? "Procesando..." : "Calcular"}</h5>
        </button>
      </form>
      <div style={{ width: "100%", height: "500px" }}>
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          onConnect={onConnect}
          fitView
        >
          <MiniMap />
          <Controls />
          <Background />
        </ReactFlow>
      </div>
    </div>
  );
}

export default GrafosForm;
