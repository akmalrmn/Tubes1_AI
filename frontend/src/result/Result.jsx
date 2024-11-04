import React from 'react';
import '../result/Result.css';

const Result = ({ algorithmData }) => {
  return (
    <div className='result'>
      <div className='result-container'>
        <div className='result-title'>
          <h1>Result</h1>
        </div>
        <div className='result-content'>
          {algorithmData ? (
            <>
              <div className='duration'>
                <p>Duration: {algorithmData.duration} s</p>
              </div>
              <div className='objective'>
                <p>Objective: {algorithmData.finalObjectiveVal}</p>
              </div>
              <div className='plot'>
                <div className='plot-left'>
                  <p>Plot: {algorithmData.plot}</p>
                </div>
                <div className='plot-right'>
                  <p>Plot eET: {algorithmData.plotEET}</p>
                </div>
              </div>
              <div className='iteration'>
                <p>Iteration: {algorithmData.totalIterations}</p>
              </div>
              <div className='frequency'>
                <p>Frequency: {algorithmData.stuckCount}</p>
              </div>
              <div className='population'>
                <p>Population: {algorithmData.initialEnergy}</p>
              </div>
            </>
          ) : (
            <p>No data available</p>
          )}
        </div>
      </div>
    </div>
  );
};

export default Result;