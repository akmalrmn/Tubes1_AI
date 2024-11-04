import { useState, useEffect, useRef } from "react";
import { useNavigate } from "react-router-dom";
import "../cube/Cube.css";
import {
  runSimulatedAnnealing,
  runGeneticAlgorithm,
  runSteepestAscent,
} from "../services/api";

const Cube = () => {
  const [rotation, setRotation] = useState({ x: -30, y: -45 });
  const [initialCubeData, setInitialCubeData] = useState(null);
  const [finalCubeData, setFinalCubeData] = useState(null);
  const [algorithmData, setAlgorithmData] = useState(null);
  const cubeRef = useRef(null);
  const touchRef = useRef({ x: 0, y: 0 });
  const isDragging = useRef(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      const data = await runSimulatedAnnealing();
      setInitialCubeData(data.initialState.Cube.Tables);
      setFinalCubeData(data.finalState.Cube.Tables);
      setAlgorithmData(data);
    };
    fetchData();
  }, []);

  const handleTouchStart = (e) => {
    touchRef.current = {
      x: e.touches[0].clientX,
      y: e.touches[0].clientY,
    };
  };

  const handleTouchMove = (e) => {
    const deltaX = e.touches[0].clientX - touchRef.current.x;
    const deltaY = e.touches[0].clientY - touchRef.current.y;
    touchRef.current = {
      x: e.touches[0].clientX,
      y: e.touches[0].clientY,
    };

    setRotation((prevRotation) => ({
      x: prevRotation.x - deltaY * 0.5,
      y: prevRotation.y + deltaX * 0.5,
    }));
  };

  const handleMouseDown = (e) => {
    isDragging.current = true;
    touchRef.current = {
      x: e.clientX,
      y: e.clientY,
    };
  };

  const handleMouseMove = (e) => {
    if (!isDragging.current) return;

    const deltaX = e.clientX - touchRef.current.x;
    const deltaY = e.clientY - touchRef.current.y;
    touchRef.current = {
      x: e.clientX,
      y: e.clientY,
    };

    setRotation((prevRotation) => ({
      x: prevRotation.x - deltaY * 0.5,
      y: prevRotation.y + deltaX * 0.5,
    }));
  };

  const handleMouseUp = () => {
    isDragging.current = false;
  };

  const handleExpandClick = () => {
    setRotation({ x: -30, y: -45 });
    setTimeout(() => {
      setTimeout(() => {
        navigate("/expand");
      }, 1000);
    }, 0);
  };

  const getFaceData = (cubeData, face) => {
    if (!cubeData) return null;

    const size = cubeData[0][0].length;

    switch (face) {
      case "front":
        return cubeData[0];
      case "back":
        return cubeData[4].map((row, rowIndex) =>
          row.map((_, colIndex) => cubeData[4][rowIndex][colIndex]).reverse()
        );
      case "right":
        return cubeData[0].map((_, colIndex) =>
          cubeData.map((table) => table[colIndex][size - 1])
        );
      case "left":
        return cubeData[0].map((_, colIndex) =>
          cubeData.map((table) => table[colIndex][0]).reverse()
        );
      case "top":
        return cubeData.map((table) => table[0]).reverse();
      case "bottom":
        return cubeData.map((table) => table[size - 1]);
      default:
        return null;
    }
  };

  const faces = ["front", "back", "left", "right", "top", "bottom"];

  const handleAlgClick = async (e) => {
    document.querySelectorAll(".alg-button button").forEach((button) => {
      button.classList.remove("active");
    });
    e.target.classList.add("active");

    let data;
    if (e.target.id === "geneticAlgorithm") {
      data = await runGeneticAlgorithm(50, 100);
      if (data) {
        setInitialCubeData(data.highestIndividual.Tables);
        setFinalCubeData(data.lowestIndividual.Tables);
        setAlgorithmData(data);
      }
    } else if (e.target.id === "simulatedAnnealing") {
      data = await runSimulatedAnnealing();
      if (data) {
        setInitialCubeData(data.initialState.Cube.Tables);
        setFinalCubeData(data.finalState.Cube.Tables);
        setAlgorithmData(data);
      }
    } else if (e.target.id === "steepestAscent") {
      data = await runSteepestAscent(1000, 25);
      if (data) {
        setInitialCubeData(data.initialState.Cube.Tables);
        setFinalCubeData(data.finalState.Cube.Tables);
        setAlgorithmData(data);
      }
    }
  };

  return (
    <div className="cube-container">
      <div className="container">
        <div className="top-container">
          <div className="left-container">
            <div className="title">
              <h1>Initial Cube</h1>
            </div>
            <div className="cube-box">
              <div className="cube-wrapper">
                <div
                  className={`cube`}
                  onTouchStart={handleTouchStart}
                  onTouchMove={handleTouchMove}
                  onMouseDown={handleMouseDown}
                  onMouseMove={handleMouseMove}
                  onMouseUp={handleMouseUp}
                  onMouseLeave={handleMouseUp}
                >
                  <div
                    className="main-cube"
                    ref={cubeRef}
                    style={{
                      transform: `rotateX(${rotation.x}deg) rotateY(${rotation.y}deg)`,
                    }}
                  >
                    {initialCubeData &&
                      faces.map((face, faceIdx) => (
                        <div
                          key={faceIdx}
                          className={`cube-face cube-face-${face}`}
                        >
                          <div className="grid-container">
                            {getFaceData(initialCubeData, face)?.map(
                              (row, rowIdx) => (
                                <div key={rowIdx} className="grid-row">
                                  {Array.isArray(row) ? (
                                    row.map((cell, cellIdx) => (
                                      <div key={cellIdx} className="grid-cells">
                                        {cell}
                                      </div>
                                    ))
                                  ) : (
                                    <div key={rowIdx} className="grid-cells">
                                      {row}
                                    </div>
                                  )}
                                </div>
                              )
                            )}
                          </div>
                        </div>
                      ))}
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="right-container">
            <div className="title">
              <h1>Final Cube</h1>
            </div>
            <div className="cube-box">
              <div className="cube-wrapper">
                <div
                  className={`cube`}
                  onTouchStart={handleTouchStart}
                  onTouchMove={handleTouchMove}
                  onMouseDown={handleMouseDown}
                  onMouseMove={handleMouseMove}
                  onMouseUp={handleMouseUp}
                  onMouseLeave={handleMouseUp}
                >
                  <div
                    className="main-cube"
                    ref={cubeRef}
                    style={{
                      transform: `rotateX(${rotation.x}deg) rotateY(${rotation.y}deg)`,
                    }}
                  >
                    {finalCubeData &&
                      faces.map((face, faceIdx) => (
                        <div
                          key={faceIdx}
                          className={`cube-face cube-face-${face}`}
                        >
                          <div className="grid-container">
                            {getFaceData(finalCubeData, face)?.map(
                              (row, rowIdx) => (
                                <div key={rowIdx} className="grid-row">
                                  {Array.isArray(row) ? (
                                    row.map((cell, cellIdx) => (
                                      <div key={cellIdx} className="grid-cells">
                                        {cell}
                                      </div>
                                    ))
                                  ) : (
                                    <div key={rowIdx} className="grid-cells">
                                      {row}
                                    </div>
                                  )}
                                </div>
                              )
                            )}
                          </div>
                        </div>
                      ))}
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="bottom-container">
          <div className="alg-button">
            <button id="geneticAlgorithm" onClick={handleAlgClick}>
              Genetic Algorithm
            </button>
            <button id="simulatedAnnealing" onClick={handleAlgClick}>
              Simulated Annealing
            </button>
            <button id="steepestAscent" onClick={handleAlgClick}>
              Steepest Ascent
            </button>
          </div>
          <div className="expand-button">
            <button onClick={handleExpandClick}>Expand</button>
          </div>
        </div>
        {algorithmData && (
          <div className="algorithm-data">
            <h2>Algorithm Data</h2>
            <p>Duration: {algorithmData.duration}</p>
            <p>Total Iterations: {algorithmData.totalIterations}</p>
            <p>Final Objective Value: {algorithmData.finalObjectiveVal}</p>
            <p>Stuck Count: {algorithmData.stuckCount}</p>
            <p>Initial Energy: {algorithmData.initialEnergy}</p>
          </div>
        )}
      </div>
    </div>
  );
};

export default Cube;