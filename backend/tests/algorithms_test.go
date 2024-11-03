package tests

import (
    "fmt"
    "magic-cube-solver/internal/algorithms/simulated_annealing"
    "magic-cube-solver/internal/models"
    "testing"
)

func formatCube(cube models.Cube) string {
    result := ""

    // Create slices for rows, columns, and heights
    rows := make([][][]string, 5)
    columns := make([][][]string, 5)
    heights := make([][][]string, 5)

    for i := 0; i < 5; i++ {
        rows[i] = make([][]string, 5)
        columns[i] = make([][]string, 5)
        heights[i] = make([][]string, 5)
        for j := 0; j < 5; j++ {
            rows[i][j] = make([]string, 5)
            heights[i][j] = make([]string, 5)
            columns[i][j] = make([]string, 5)
        }
    }

    // Fill rows
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            for k := 0; k < 5; k++ {
                rows[i][j][k] = cube.Tables[i][j][k]
            }
        }
    }

    // Fill columns
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            for k := 0; k < 5; k++ {
                columns[i][j][k] = cube.Tables[j][i][k]
            }
        }
    }

    // Fill heights
    for i := 0; i < 5; i++ {
        for j := 0; j < 5; j++ {
            for k := 0; k < 5; k++ {
                heights[i][j][k] = cube.Tables[j][k][i]
            }
        }
    }

    // Format rows
    for i := 0; i < 5; i++ {
        result += fmt.Sprintf("Row %d:\n", i+1)
        for _, row := range rows[i] {
            result += fmt.Sprintf("%v\n", row)
        }
        result += "\n"
    }

    // Format columns
    for i := 0; i < 5; i++ {
        result += fmt.Sprintf("Column %d:\n", i+1)
        for _, column := range columns[i] {
            result += fmt.Sprintf("%v\n", column)
        }
        result += "\n"
    }

    // Format heights
    for i := 0; i < 5; i++ {
        result += fmt.Sprintf("Height %d:\n", i+1)
        for _, height := range heights[i] {
            result += fmt.Sprintf("%v\n", height)
        }
        result += "\n"
    }

    return result
}

func formatInitialCube(cube models.Cube) string {
    result := ""
    for tableIdx, table := range cube.Tables {
        result += fmt.Sprintf("Table %d:\n", tableIdx+1)
        for _, row := range table {
            result += fmt.Sprintf("%v\n", row)
        }
        result += "\n"
    }
    return result
}

func TestSimulatedAnnealing(t *testing.T) {
    // Create an initial cube
    initialCube := models.GenerateCube() 

    _, finalState, energyHistory, tempHistory, duration, stuckCount, initialEnergy, totalIterations := simulated_annealing.SimulatedAnnealing(initialCube)

    if initialEnergy == 0 {
        t.Errorf("Initial state energy should not be zero")
    }

    if finalState.Energy == 0 {
        t.Errorf("Final state energy should not be zero")
    }

    if len(energyHistory) < 1 {
        t.Errorf("Energy history length should be at least 1, got %d", len(energyHistory))
    }

    if len(tempHistory) < 1 {
        t.Errorf("Temperature history length should be at least 1, got %d", len(tempHistory))
    }

    if stuckCount < 0 {
        t.Errorf("Stuck count should not be negative")
    }

    fmt.Printf("Initial Cube (before partitioning):\n%s\n", formatInitialCube(initialCube))
    fmt.Printf("Initial Cube (after partitioning):\n%s\n", formatCube(initialCube))
    fmt.Printf("Initial State Energy: %f\n", initialEnergy)
    fmt.Printf("Final State Energy: %f\n", finalState.Energy)
    fmt.Printf("Best Energy Achieved: %f\n", finalState.Energy)
    fmt.Printf("Duration: %v\n", duration)
    fmt.Printf("Stuck Count: %d\n", stuckCount)
    fmt.Printf("Total Iterations: %d\n", totalIterations)

    fmt.Printf("Energy History (first 5 iterations): %v\n", energyHistory[:5])
    fmt.Printf("Temperature History (first 5 iterations): %v\n", tempHistory[:5])
}