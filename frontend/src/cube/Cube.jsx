import React, { useState, useRef } from 'react';
import { useNavigate } from 'react-router-dom';
import '../cube/Cube.css';

const Cube = () => {
  const [rotation, setRotation] = useState({ x: -30, y: -45 });
  const [isShrinking, setIsShrinking] = useState(false);
  const cubeRef = useRef(null);
  const touchRef = useRef({ x: 0, y: 0 });
  const isDragging = useRef(false);
  const navigate = useNavigate();

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
          <div className='cube-face cube-face-front'>
            <div className='grid-container'>
              {[...Array(5)].map((_, gridRowIndex) => (
                <div key={gridRowIndex} className='grid-row'>
                  {[...Array(5)].map((_, gridColIndex) => {
                    const number = 20 + gridRowIndex * 25 + gridColIndex + 1;
                    return (
                      <div key={gridColIndex} className='grid-cell'>
                        {number}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
          <div className='cube-face cube-face-back'>
            <div className='grid-container'>
              {[...Array(5)].map((_, gridRowIndex) => (
                <div key={gridRowIndex} className='grid-row'>
                  {[...Array(5)].map((_, gridColIndex) => {
                    const number = 20 + gridRowIndex * 25 + gridColIndex + 1;
                    return (
                      <div key={gridColIndex} className='grid-cell'>
                        {number}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
          <div className='cube-face cube-face-left'>
            <div className='grid-container'>
              {[...Array(5)].map((_, gridRowIndex) => (
                <div key={gridRowIndex} className='grid-row'>
                  {[...Array(5)].map((_, gridColIndex) => {
                    const number = 1 + gridColIndex * 5 + gridRowIndex * 25;
                    return (
                      <div key={gridColIndex} className='grid-cell'>
                        {number}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
          <div className='cube-face cube-face-right'>
            <div className='grid-container'>
              {[...Array(5)].map((_, gridRowIndex) => (
                <div key={gridRowIndex} className='grid-row'>
                  {[...Array(5)].map((_, gridColIndex) => {
                    const number = (gridRowIndex + 1) * 25 - gridColIndex * 5;
                    return (
                      <div key={gridColIndex} className='grid-cell'>
                        {number}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
          <div className='cube-face cube-face-top'>
            <div className='grid-container'>
              {[...Array(5)].map((_, gridRowIndex) => (
                <div key={gridRowIndex} className='grid-row'>
                  {[...Array(5)].map((_, gridColIndex) => {
                    const number = gridRowIndex * 5 + gridColIndex + 1;
                    return (
                      <div key={gridColIndex} className='grid-cell'>
                        {number}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
          <div className='cube-face cube-face-bottom'>
            <div className='grid-container'>
              {[...Array(5)].map((_, gridRowIndex) => (
                <div key={gridRowIndex} className='grid-row'>
                  {[...Array(5)].map((_, gridColIndex) => {
                    const number = 105 + gridRowIndex * 5 - gridColIndex;
                    return (
                      <div key={gridColIndex} className='grid-cell'>
                        {number}
                      </div>
                    );
                  })}
                </div>
              ))}
            </div>
          </div>
        </div>
      </div>
      <div className='expand-button'>
        <button onClick={handleExpandClick}>Expand</button>
      </div>
    </div>
  );
};

export default Cube;