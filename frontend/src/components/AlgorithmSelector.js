import React, { useState } from 'react';
import { runSimulatedAnnealing, runGeneticAlgorithm, runSteepestAscent } from '../services/api';

const AlgorithmSelector = () => {
    const [result, setResult] = useState(null);

    const handleRunSimulatedAnnealing = async () => {
        const data = await runSimulatedAnnealing();
        setResult(data);
    };

    const handleRunGeneticAlgorithm = async () => {
        const populationSize = 50; // Example value
        const generations = 100; // Example value
        const data = await runGeneticAlgorithm(populationSize, generations);
        setResult(data);
    };

    const handleRunSteepestAscent = async () => {
        const maxIterations = 1000; // Example value
        const targetSum = 25; // Example value
        const data = await runSteepestAscent(maxIterations, targetSum);
        setResult(data);
    };

    return (
        <div>
            <button onClick={handleRunSimulatedAnnealing}>Run Simulated Annealing</button>
            <button onClick={handleRunGeneticAlgorithm}>Run Genetic Algorithm</button>
            <button onClick={handleRunSteepestAscent}>Run Steepest Ascent</button>
            {result && (
                <div>
                    <h3>Algorithm Results</h3>
                    <p>Initial State Energy: {result.initialState.ObjectiveValue}</p>
                    <p>Final State Energy: {result.finalObjectiveVal}</p>
                    <p>Total Iterations: {result.totalIterations}</p>
                    <p>Duration: {result.duration}</p>
                    <div>
                        <h4>Initial Cube State:</h4>
                        {result.initialState.Cube.Tables.map((table, tableIdx) => (
                            <div key={tableIdx}>
                                <h5>Table {tableIdx + 1}</h5>
                                <table>
                                    <tbody>
                                        {table.map((row, rowIdx) => (
                                            <tr key={rowIdx}>
                                                {row.map((cell, cellIdx) => (
                                                    <td key={cellIdx}>{cell}</td>
                                                ))}
                                            </tr>
                                        ))}
                                    </tbody>
                                </table>
                            </div>
                        ))}
                    </div>
                    <div>
                        <h4>Final Cube State:</h4>
                        {result.finalState.Cube.Tables.map((table, tableIdx) => (
                            <div key={tableIdx}>
                                <h5>Table {tableIdx + 1}</h5>
                                <table>
                                    <tbody>
                                        {table.map((row, rowIdx) => (
                                            <tr key={rowIdx}>
                                                {row.map((cell, cellIdx) => (
                                                    <td key={cellIdx}>{cell}</td>
                                                ))}
                                            </tr>
                                        ))}
                                    </tbody>
                                </table>
                            </div>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
};

export default AlgorithmSelector;