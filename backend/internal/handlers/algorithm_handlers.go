package handlers

import (
    "encoding/json"
    "net/http"
    "magic-cube-solver/internal/algorithms/geneticalgorithm"
)

func GeneticAlgorithmHandler(w http.ResponseWriter, r *http.Request) {
    cube := geneticalgorithm.GenerateCube()

    response := struct {
        Cube            geneticalgorithm.Cube `json:"cube"`
        ColumnSums      []int                  `json:"columnSums"`
        RowSums         []int                  `json:"rowSums"`
        PoleSums        []int                  `json:"poleSums"`
    }{
        Cube:       cube,
        ColumnSums: geneticalgorithm.SumColumns(cube.Tables),
        RowSums:    geneticalgorithm.SumRows(cube.Tables),
        PoleSums:   geneticalgorithm.SumPoles(cube.Tables),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}