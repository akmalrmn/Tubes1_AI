export async function runGeneticAlgorithm(initialState, populationSize, generations) {
    const response = await fetch('http://localhost:8070/api/genetic-algorithm', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ initialState, populationSize, generations }),
    });
    return response.json();
}

export async function runSimulatedAnnealing() {
    const response = await fetch('http://localhost:8070/api/anies', {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
        },
    });
    return response.json();
}