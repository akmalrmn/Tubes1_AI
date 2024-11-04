import React from 'react'
import '../result/Result.css'

const Result = () => {
  return (
    <div className='result'>
        <div className='result-container'>
            <div className='result-title'>
                <h1>Result</h1>
            </div>
            <div className='result-content'>
                <div className='duration'>
                  <p>Duration: </p>
                </div>
                <div className='objective'>
                  <p>Objective: </p>
                </div>
                <div className='plot'>
                  <div className='plot-left'>
                    <p>Plot: </p>
                  </div>
                  <div className='plot-right'>
                    <p>Plot eET: </p>
                  </div>
                </div>
                <div className='iteration'>
                  <p>Iteration: </p>
                </div>
                <div className='frequency'>
                  <p>Frequency:  </p>
                </div>
                <div className='population'>
                  <p>Population: </p>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Result