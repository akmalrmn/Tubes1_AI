import React from 'react';
import { useNavigate } from 'react-router-dom';
import './Expand.css';

const Expand = () => {
  const navigate = useNavigate();

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
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {21 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {25 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className='cubes-container' id='cube2'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {46 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {50 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {25 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className='cubes-container' id='cube3'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {71 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {75 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {50 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className='cubes-container' id='cube4'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {96 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {100 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {75 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className  ='cubes-container' id='cube5'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {121 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {125 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {100 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
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
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {21 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {25 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className='cubes-container' id='cube2'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {46 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {50 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {25 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className='cubes-container' id='cube3'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {71 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {75 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {50 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className='cubes-container' id='cube4'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {96 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {100 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {75 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
              </div>
            </div>
            <div className  ='cubes-container' id='cube5'>
              <div className='expand-cube'>
                <div className="face front">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {121 + colIndex}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face back"></div>
                <div className="face left"></div>
                <div className="face right">
                  <div className='grid-container'>
                    <div className='grid-row'>
                      {[...Array(5)].map((_, colIndex) => (
                        <div key={colIndex} className='grid-cell'>
                          {125 - colIndex * 5}
                        </div>
                      ))}
                    </div>
                  </div>
                </div>
                <div className="face top">
                  <div className='grid-container'>
                    {[...Array(5)].map((_, rowIndex) => (
                      <div key={rowIndex} className='grid-row'>
                        {[...Array(5)].map((_, colIndex) => (
                          <div key={colIndex} className='grid-cell'>
                            {100 + rowIndex * 5 + colIndex + 1}
                          </div>
                        ))}
                      </div>
                    ))}
                  </div>
                </div>
                <div className="face bottom"></div>
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