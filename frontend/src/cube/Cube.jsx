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
  const cubeRef = useRef(null);
  const touchRef = useRef({ x: 0, y: 0 });
  const isDragging = useRef(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      const data = await runSimulatedAnnealing();
      setInitialCubeData(data.initialState.Cube.Tables);
      setFinalCubeData(data.finalState.Cube.Tables);
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

    const size = cubeData[0].length;

    switch (face) {
      case "front":
        return cubeData[0];
      case "back":
        return cubeData[4];
      case "left":
        return cubeData.map((row) => row[0]);
      case "right":
        return cubeData.map((row) => row[size - 1]);
      case "top":
        return cubeData.map((row, idx) => cubeData[idx][0]);
      case "bottom":
        return cubeData.map((row, idx) => cubeData[idx][size - 1]);
      default:
        return null;
    }
  };

  const faces = ["front", "back", "left", "right", "top", "bottom"];

  const handleAlgClick = async (e) => {
    // Remove 'active' class from all buttons and add to clicked button
    document.querySelectorAll(".alg-button button").forEach((button) => {
      button.classList.remove("active");
    });
    e.target.classList.add("active");

    
    let data;
    if (e.target.id === "geneticAlgorithm") {
      data = await runGeneticAlgorithm(50, 100);
    } else if (e.target.id === "simulatedAnnealing") {
      data = await runSimulatedAnnealing();
    } else if (e.target.id === "steepestAscent") {
      data = await runSteepestAscent(1000, 25);
    }

    // Update cube data if data is successfully fetched
    if (data) {
      setInitialCubeData(data.initialState.Cube.Tables);
      setFinalCubeData(data.finalState.Cube.Tables);
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
                                      <div key={cellIdx} className="grid-cell">
                                        {cell}
                                      </div>
                                    ))
                                  ) : (
                                    <div key={rowIdx} className="grid-cell">
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
                                      <div key={cellIdx} className="grid-cell">
                                        {cell}
                                      </div>
                                    ))
                                  ) : (
                                    <div key={rowIdx} className="grid-cell">
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
      </div>
    </div>
  );
};

export default Cube;
