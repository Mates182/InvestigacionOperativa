import React from "react";
import Link from "next/link";

function Home() {
  return (
    <div>
      <h1>Optimización de Problemas de Investigación Operativa</h1>

      <section>
        <h2>Descripción del Programa</h2>
        <p>
          Este programa resuelve problemas de optimización en los siguientes
          ámbitos:
        </p>
        <ul>
          <li>Programación Lineal</li>
          <li>Transporte</li>
          <li>Redes</li>
        </ul>
        <p>
          Además, integra Inteligencia Artificial para realizar análisis de
          sensibilidad y toma de decisiones asertivas.
        </p>
      </section>

      <section>
        <h2>Integrantes del Proyecto</h2>
        <ul>
          <li>Mateo Bernardo Pillajo López</li>
          <li>Christian Rolando Tapia Diaz</li>
          <li>Lenin Sebastián Serrano Montúfar</li>
        </ul>
      </section>
    </div>
  );
}

export default Home;
