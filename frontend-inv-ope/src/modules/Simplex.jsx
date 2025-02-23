import React from "react";

function Simplex({ resoluciones }) {
  return (
    <div>
      {resoluciones.map((resolucion, i) => {
        return (
          <div key={i}>
            <h3>Iteración {i}</h3>
            <table className="table">
              <thead className="table-dark">
                <tr>
                  <th>V.B</th>
                  <th>Z</th>

                  {resolucion.ecuaciones[0].li.map((termino, j) => {
                    return <th key={j}>{termino.vd}</th>;
                  })}
                  <th>L.D</th>
                </tr>
              </thead>
              <tbody>
                {resolucion.ecuaciones.map((ecuacion, j) => (
                  <tr key={j}>
                    <td>{ecuacion.vb}</td>
                    <td>{j == 0 ? 1 : 0}</td>

                    {ecuacion.li.map((termino, k) => (
                      <td key={k}>{Math.round(termino.c * 10000) / 10000}</td>
                    ))}
                    <td>{Math.round(ecuacion.ld * 10000) / 10000}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        );
      })}
    </div>
  );
}

export default Simplex;
