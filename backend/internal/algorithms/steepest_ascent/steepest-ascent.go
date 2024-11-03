package steepest_ascent

import (
    "fmt"
    "magic-cube-solver/internal/models"
    "time"
)

type State struct {
    Cube           models.Cube
    ObjectiveValue float64
}

func SteepestAscent(maxIterations int) (State, State, []float64, time.Duration, int) {
    startTime := time.Now()

    initial := State{
        Cube:           models.GenerateCube(),
        ObjectiveValue: 0,
    }
    initial.ObjectiveValue = models.EvaluateIndividual(initial.Cube)

    current := initial
    objectiveHistory := []float64{current.ObjectiveValue}

    for i := 0; i < maxIterations; i++ {
        bestNeighbor, improved, swapMessages := findBestNeighbor(current)

        if improved {
            current = bestNeighbor
            objectiveHistory = append(objectiveHistory, current.ObjectiveValue)
            fmt.Printf("\nIteration %d: Objective Value = %.0f\n", i+1, current.ObjectiveValue)
            for _, message := range swapMessages {
                fmt.Println("   " + message)
            }
        } else {
            break
        }
    }

    duration := time.Since(startTime)

    return initial, current, objectiveHistory, duration, len(objectiveHistory)
}

func findBestNeighbor(current State) (State, bool, []string) {
    bestNeighbor := current
    improved := false
    lastSwaps := make(map[string]string)

    for tableIdx1 := 0; tableIdx1 < models.NumTables; tableIdx1++ {
        for rowIdx1 := 0; rowIdx1 < models.Rows; rowIdx1++ {
            for colIdx1 := 0; colIdx1 < models.Cols; colIdx1++ {
                for tableIdx2 := 0; tableIdx2 < models.NumTables; tableIdx2++ {
                    for rowIdx2 := 0; rowIdx2 < models.Rows; rowIdx2++ {
                        for colIdx2 := 0; colIdx2 < models.Cols; colIdx2++ {
                            if tableIdx1 == tableIdx2 && rowIdx1 == rowIdx2 && colIdx1 == colIdx2 {
                                continue
                            }

                            neighborCube := copyCube(current.Cube)
                            val1 := neighborCube.Tables[tableIdx1][rowIdx1][colIdx1]
                            val2 := neighborCube.Tables[tableIdx2][rowIdx2][colIdx2]
                            neighborCube.Tables[tableIdx1][rowIdx1][colIdx1], neighborCube.Tables[tableIdx2][rowIdx2][colIdx2] = val2, val1

                            neighborObjectiveValue := models.EvaluateIndividual(neighborCube)

                            if neighborObjectiveValue < bestNeighbor.ObjectiveValue {
                                bestNeighbor = State{Cube: neighborCube, ObjectiveValue: neighborObjectiveValue}
                                improved = true

                                swapMessage := fmt.Sprintf("Swapped value %s at Table %d Position (%d, %d) with value %s at Table %d Position (%d, %d)",
                                    val1, tableIdx1+1, rowIdx1+1, colIdx1+1, val2, tableIdx2+1, rowIdx2+1, colIdx2+1)
                                lastSwaps[fmt.Sprintf("%d,%d,%d", tableIdx1, rowIdx1, colIdx1)] = swapMessage
                            }
                        }
                    }
                }
            }
        }
    }

    swapMessages := make([]string, 0, len(lastSwaps))
    for _, message := range lastSwaps {
        swapMessages = append(swapMessages, message)
    }

    return bestNeighbor, improved, swapMessages
}

func copyCube(cube models.Cube) models.Cube {
    newCube := models.Cube{
        Tables: make([][][]string, len(cube.Tables)),
    }

    for i := range cube.Tables {
        newCube.Tables[i] = make([][]string, len(cube.Tables[i]))
        for j := range cube.Tables[i] {
            newCube.Tables[i][j] = make([]string, len(cube.Tables[i][j]))
            copy(newCube.Tables[i][j], cube.Tables[i][j])
        }
    }

    return newCube
}