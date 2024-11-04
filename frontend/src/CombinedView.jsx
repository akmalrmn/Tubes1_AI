import React from 'react';
import Cube from './cube/Cube';
import Result from './result/Result';
import './CombinedView.css';

const CombinedView = () => {
  return (
    <div className="combined-view">
      <Cube />
      <Result />
    </div>
  );
};

export default CombinedView;