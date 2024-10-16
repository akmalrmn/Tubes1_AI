import React, { useState } from 'react';
import { runGeneticAlgorithm } from '../services/api';

function AlgorithmSelector() {
    const [result, setResult] = useState(null);

    const handleRunAlgorithm = async () => {
        const initialState = {};
        const populationSize = 100;
        const generations = 50;
        const result = await runGeneticAlgorithm(initialState, populationSize, generations);
        setResult(result);
    };

    return (
        <div>
            <button onClick={handleRunAlgorithm}>Run Genetic Algorithm</button>
            {result && <pre>{JSON.stringify(result, null, 2)}</pre>}
        </div>
    );
}

export default AlgorithmSelector;