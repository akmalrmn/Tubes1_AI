export async function runGeneticAlgorithm(initialState, populationSize, generations) {
    const response = await fetch('http://localhost:8080/api/genetic-algorithm', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ initialState, populationSize, generations }),
    });
    return response.json();
}