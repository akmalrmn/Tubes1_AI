import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import './Expand.css';
import { runSimulatedAnnealing } from '../services/api';

const Expand = () => {
  const navigate = useNavigate();
  const [initialCubeData, setInitialCubeData] = useState(null);
  const [finalCubeData, setFinalCubeData] = useState(null);

  useEffect(() => {
    const fetchData = async () => {
      const data = await runSimulatedAnnealing();
      setInitialCubeData(data.initialState.Cube.Tables);
      setFinalCubeData(data.finalState.Cube.Tables);
    };
    fetchData();
  }, []);

  const getInitialFaceData1 = (face) => {
    if (!initialCubeData) return null;
    
    const size = initialCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return initialCubeData[0].slice(0, 1);
      
      case 'back':
        return initialCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => initialCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][size - 1])
        ).slice(0, 1);
      
      case 'left':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return initialCubeData.map(table => table[0]).reverse();
      
      case 'bottom':
        return initialCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getInitialFaceData2 = (face) => {
    if (!initialCubeData) return null;
    
    const size = initialCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return initialCubeData[0].slice(1, 2);
      
      case 'back':
        return initialCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => initialCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][size - 1])
        ).slice(1, 2);
      
      case 'left':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return initialCubeData.map(table => table[1]).reverse();
      
      case 'bottom':
        return initialCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getInitialFaceData3 = (face) => {
    if (!initialCubeData) return null;
    
    const size = initialCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return initialCubeData[0].slice(2, 3);
      
      case 'back':
        return initialCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => initialCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][size - 1])
        ).slice(2, 3);
      
      case 'left':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return initialCubeData.map(table => table[2]).reverse();
      
      case 'bottom':
        return initialCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getInitialFaceData4 = (face) => {
    if (!initialCubeData) return null;
    
    const size = initialCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return initialCubeData[0].slice(3, 4);
      
      case 'back':
        return initialCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => initialCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][size - 1])
        ).slice(3, 4);
      
      case 'left':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return initialCubeData.map(table => table[3]).reverse();
      
      case 'bottom':
        return initialCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getInitialFaceData5 = (face) => {
    if (!initialCubeData) return null;
    
    const size = initialCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return initialCubeData[0].slice(4, 5);
      
      case 'back':
        return initialCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => initialCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][size - 1])
        ).slice(4, 5);
      
      case 'left':
        return initialCubeData[0].map((_, colIndex) => 
          initialCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return initialCubeData.map(table => table[4]).reverse();
      
      case 'bottom':
        return initialCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getFinalFaceData1 = (face) => {
    if (!finalCubeData) return null;
    
    const size = finalCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return finalCubeData[0].slice(0, 1);
      
      case 'back':
        return finalCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => finalCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][size - 1])
        ).slice(0, 1);
      
      case 'left':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return finalCubeData.map(table => table[0]).reverse();
      
      case 'bottom':
        return finalCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getFinalFaceData2 = (face) => {
    if (!finalCubeData) return null;
    
    const size = finalCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return finalCubeData[0].slice(1, 2);
      
      case 'back':
        return finalCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => finalCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][size - 1])
        ).slice(1, 2);
      
      case 'left':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return finalCubeData.map(table => table[1]).reverse();
      
      case 'bottom':
        return finalCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getFinalFaceData3 = (face) => {
    if (!finalCubeData) return null;
    
    const size = finalCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return finalCubeData[0].slice(2, 3);
      
      case 'back':
        return finalCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => finalCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][size - 1])
        ).slice(2, 3);
      
      case 'left':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return finalCubeData.map(table => table[2]).reverse();
      
      case 'bottom':
        return finalCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getFinalFaceData4 = (face) => {
    if (!finalCubeData) return null;
    
    const size = finalCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return finalCubeData[0].slice(3, 4);
      
      case 'back':
        return finalCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => finalCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][size - 1])
        ).slice(3, 4);
      
      case 'left':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return finalCubeData.map(table => table[3]).reverse();
      
      case 'bottom':
        return finalCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const getFinalFaceData5 = (face) => {
    if (!finalCubeData) return null;
    
    const size = finalCubeData[0][0].length;
    
    switch(face) {
      case 'front':
        return finalCubeData[0].slice(4, 5);
      
      case 'back':
        return finalCubeData[4].map((row, rowIndex) => 
          row.map((_, colIndex) => finalCubeData[4][colIndex][rowIndex])
        );
      
      case 'right':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][size - 1])
        ).slice(4, 5);
      
      case 'left':
        return finalCubeData[0].map((_, colIndex) => 
          finalCubeData.map((table) => table[colIndex][0])
        );
      
      case 'top':
        return finalCubeData.map(table => table[4]).reverse();
      
      case 'bottom':
        return finalCubeData.map(table => table[size - 1]);
      
      default:
        return null;
    }
  };

  const faces = ['front', 'back', 'left', 'right', 'top', 'bottom'];

  const handleMergeClick = () => {
    document.querySelectorAll('.cubes-container').forEach((cube) => {
      cube.classList.add('move-back');
    });

    setTimeout(() => {
      navigate('/');
    }, 2500);
  };

  return (
    <div className='expand-container'>
      <div className='expand-initial'>
        <div className='initial-container'>
          <div className='title'>
            <h1>Initial Cube</h1>
          </div>
          <div className='expand-box'>
            <div className='cubes-container' id='cube1'>
              <div className='expand-cube'>
                {faces.map((face, faceIdx) => (
                  <div key={faceIdx} className={`face ${face}`}>
                    <div className='grid-container'>
                      {getInitialFaceData1(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube2'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getInitialFaceData2(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube3'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getInitialFaceData3(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube4'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getInitialFaceData4(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube5'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getInitialFaceData5(face)?.map((row, rowIdx) => (
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
          </div>
        </div>
      </div>
      <div className='expand-final'>
        <div className='final-container'>
          <div className='title'>
            <h1>Final Cube</h1>
          </div>
          <div className='expand-box'>
            <div className='cubes-container' id='cube1'>
              <div className='expand-cube'>
                {faces.map((face, faceIdx) => (
                  <div key={faceIdx} className={`face ${face}`}>
                    <div className='grid-container'>
                      {getFinalFaceData1(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube2'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getFinalFaceData2(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube3'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getFinalFaceData3(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube4'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getFinalFaceData4(face)?.map((row, rowIdx) => (
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
            <div className='cubes-container' id='cube5'>
              <div className='expand-cube'>
                  {faces.map((face, faceIdx) => (
                    <div key={faceIdx} className={`face ${face}`}>
                      <div className='grid-container'>
                        {getFinalFaceData5(face)?.map((row, rowIdx) => (
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
          </div>
        </div>
      </div>
      <div className='merge-button'>
        <button onClick={handleMergeClick}>Merge</button>
      </div>
    </div>
  );
};

export default Expand;