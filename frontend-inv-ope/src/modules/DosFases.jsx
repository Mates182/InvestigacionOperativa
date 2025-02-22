import React from "react";
import Simplex from "@/modules/Simplex";

function DosFases({ resoluciones }) {
  return (
    <div>
      <h2>Fase 1</h2>
      <Simplex resoluciones={resoluciones.resolucion_fase_1}></Simplex>
      <h2>Fase 2</h2>
      <Simplex resoluciones={resoluciones.resolucion_fase_2}></Simplex>
    </div>
  );
}

export default DosFases;
