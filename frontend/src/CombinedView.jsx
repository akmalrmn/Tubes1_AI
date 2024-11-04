import React, { useState } from 'react';
import Cube from './cube/Cube';
import Result from './result/Result';
import './CombinedView.css';

const CombinedView = () => {
  const [algorithmData, setAlgorithmData] = useState(null);

  return (
    <div className="combined-view">
      <Cube setAlgorithmData={setAlgorithmData} />
      <Result algorithmData={algorithmData} />
    </div>
  );
};

export default CombinedView;