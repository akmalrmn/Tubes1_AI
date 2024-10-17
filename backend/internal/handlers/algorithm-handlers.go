package handlers

import (
    "encoding/json"
    "net/http"
    "time"
    "magic-cube-solver/internal/algorithms/genetic_algorithm"
    "magic-cube-solver/internal/algorithms/simulated_annealing"
)

func GeneticAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
    cube := genetic_algorithm.GenerateCube()

    response := struct {
        Cube       genetic_algorithm.Cube `json:"cube"`
        ColumnSums []int                  `json:"columnSums"`
        RowSums    []int                  `json:"rowSums"`
        PoleSums   []int                  `json:"poleSums"`
    }{
        Cube:       cube,
        ColumnSums: genetic_algorithm.SumColumns(cube.Tables),
        RowSums:    genetic_algorithm.SumRows(cube.Tables),
        PoleSums:   genetic_algorithm.SumPoles(cube.Tables),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func SimulatedAnnealingHandler(w http.ResponseWriter, r *http.Request) {
    initialTemp := 100.0
    coolingRate := 0.003
    maxIterations := 10000

    initialState, finalState, energyHistory, tempHistory, duration, stuckCount := simulated_annealing.SimulatedAnnealing(initialTemp, coolingRate, maxIterations)

    response := struct {
        InitialState  simulated_annealing.State `json:"initialState"`
        FinalState    simulated_annealing.State `json:"finalState"`
        EnergyHistory []float64                 `json:"energyHistory"`
        TempHistory   []float64                 `json:"tempHistory"`
        Duration      time.Duration             `json:"duration"`
        StuckCount    int                       `json:"stuckCount"`
    }{
        InitialState:  initialState,
        FinalState:    finalState,
        EnergyHistory: energyHistory,
        TempHistory:   tempHistory,
        Duration:      duration,
        StuckCount:    stuckCount,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}