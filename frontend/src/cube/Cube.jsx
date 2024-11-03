import { useState, useEffect, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import '../cube/Cube.css';
import { runSimulatedAnnealing } from '../services/api';

const Cube = () => {
  const [rotation, setRotation] = useState({ x: -30, y: -45 });
  const [isShrinking, setIsShrinking] = useState(false);
  const [cubeData, setCubeData] = useState(null);
  const cubeRef = useRef(null);
  const touchRef = useRef({ x: 0, y: 0 });
  const isDragging = useRef(false);
  const navigate = useNavigate();

  useEffect(() => {
    const fetchData = async () => {
      const data = await runSimulatedAnnealing();
      setCubeData(data.initialState.Cube.Tables);
    };
    fetchData();
  }, []);

  const handleTouchStart = (e) => {
    touchRef.current = {
      x: e.touches[0].clientX,
      y: e.touches[0].clientY
    };
  };

  const handleTouchMove = (e) => {
    const deltaX = e.touches[0].clientX - touchRef.current.x;
    const deltaY = e.touches[0].clientY - touchRef.current.y;
    touchRef.current = {
      x: e.touches[0].clientX,
      y: e.touches[0].clientY
    };

    setRotation((prevRotation) => ({
      x: prevRotation.x - deltaY * 0.5,
      y: prevRotation.y + deltaX * 0.5
    }));
  };

  const handleMouseDown = (e) => {
    isDragging.current = true;
    touchRef.current = {
      x: e.clientX,
      y: e.clientY
    };
  };

  const handleMouseMove = (e) => {
    if (!isDragging.current) return;

    const deltaX = e.clientX - touchRef.current.x;
    const deltaY = e.clientY - touchRef.current.y;
    touchRef.current = {
      x: e.clientX,
      y: e.clientY
    };

    setRotation((prevRotation) => ({
      x: prevRotation.x - deltaY * 0.5,
      y: prevRotation.y + deltaX * 0.5
    }));
  };

  const handleMouseUp = () => {
    isDragging.current = false;
  };

  const handleExpandClick = () => {
    setRotation({ x: -30, y: -45 });
    setTimeout(() => {
      setIsShrinking(true);
      setTimeout(() => {
        navigate('/expand');
      }, 2500);
    }, 0);
  };

  const getFaceData = (face) => {
    if (!cubeData) return null;
    
    const size = cubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return cubeData[0];
      
      case 'back':
        return cubeData[4];
      
      case 'right':
        return cubeData.map(table => 
          table.map(row => row[size - 1])
        );
      
      case 'left':
        return cubeData.map(table => 
          table.map(row => row[0])
        );
      
      case 'top':
        return cubeData.map(table => table[0]);
      
      case 'bottom':
        return cubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const faces = ['front', 'back', 'left', 'right', 'top', 'bottom'];

  return (
    <div className='cube-container'>
      <div
        className={`cube ${isShrinking ? 'shrink' : ''}`}
        onTouchStart={handleTouchStart}
        onTouchMove={handleTouchMove}
        onMouseDown={handleMouseDown}
        onMouseMove={handleMouseMove}
        onMouseUp={handleMouseUp}
        onMouseLeave={handleMouseUp}
      >
        <div
          className='main-cube'
          ref={cubeRef}
          style={{ transform: `rotateX(${rotation.x}deg) rotateY(${rotation.y}deg)` }}
        >
          {cubeData && faces.map((face, faceIdx) => (
            <div key={faceIdx} className={`cube-face cube-face-${face}`}>
              <div className='grid-container'>
                {getFaceData(face)?.map((row, rowIdx) => (
                  <div key={rowIdx} className='grid-row'>
                    {Array.isArray(row) ? row.map((cell, cellIdx) => (
                      <div key={cellIdx} className='grid-cell'>
                        {cell}
                      </div>
                    )) : (
                      <div key={rowIdx} className='grid-cell'>
                        {row}
                      </div>
                    )}
                  </div>
                ))}
              </div>
            </div>
          ))}
        </div>
      </div>
      <div className='expand-button'>
        <button onClick={handleExpandClick}>Expand</button>
      </div>
    </div>
  );
};

export default Cube;
