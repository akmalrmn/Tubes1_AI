package handlers

import (
    "encoding/json"
    "net/http"
    "magic-cube-solver/internal/algorithms/genetic_algorithm"
)

func GeneticAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
    cube := genetic_algorithm.GenerateCube()

    response := struct {
        Cube            genetic_algorithm.Cube `json:"cube"`
        ColumnSums      []int                  `json:"columnSums"`
        RowSums         []int                  `json:"rowSums"`
        PoleSums        []int                  `json:"poleSums"`
    }{
        Cube:       cube,
        ColumnSums: genetic_algorithm.SumColumns(cube.Tables),
        RowSums:    genetic_algorithm.SumRows(cube.Tables),
        PoleSums:   genetic_algorithm.SumPoles(cube.Tables),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}