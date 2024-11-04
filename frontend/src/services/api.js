export async function runGeneticAlgorithm(populationSize, generations) {
    const response = await fetch('/api/genetic-algorithm', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ populationSize, generations }),
    });
    return response.json();
}

export async function runSimulatedAnnealing() {
    const response = await fetch('/api/simulated-annealing', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    });
    return response.json();
}

export async function runSteepestAscent(maxIterations, targetSum) {
    const response = await fetch('/api/steepest-ascent', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ maxIterations, targetSum }),
    });
    return response.json();
}