package simulated_annealing

import (
    "fmt"
    "magic-cube-solver/internal/models"
    "math"
    "math/rand"
    "time"
)

type State struct {
    Cube   models.Cube
    Energy float64
}

func SimulatedAnnealing(initialCube models.Cube) (State, State, []float64, []float64, string, int, float64, int) {
    startTime := time.Now()
    current := State{
        Cube:   initialCube,
        Energy: models.EvaluateIndividual(initialCube),
    }

    best := current
    initialEnergy := current.Energy
    initialTemp := 100.0
    coolingRate := 0.003
    maxIterations := 10000
    temp := initialTemp
    energyHistory := []float64{current.Energy}
    acceptanceProbHistory := []float64{}
    stuckCount := 0
    iteration := 0
    stabilizationCount := 0
    stabilizationThreshold := 500 // Number of iterations to check for stabilization
    temperatureThreshold := 1e-10 

    for iteration < maxIterations {
        if temp < temperatureThreshold {
            fmt.Println("Terminating search as temperature reached the threshold.")
            break
        }

        neighbor, tableIdx1, rowIdx1, colIdx1, tableIdx2, rowIdx2, colIdx2, oldValue1, oldValue2 := generateNeighbor(current)
        neighborEnergy := models.EvaluateIndividual(neighbor.Cube)
        acceptanceProb := 0.0

        if neighborEnergy < current.Energy {
            current = State{Cube: neighbor.Cube, Energy: neighborEnergy}
            if current.Energy < best.Energy {
                best = current
            }
        } else {
            stuckCount++

            if temp >= temperatureThreshold {
                acceptanceProb = acceptanceProbability(current.Energy, neighborEnergy, temp)
                if acceptanceProb > rand.Float64() {
                    current = State{Cube: neighbor.Cube, Energy: neighborEnergy}
                    if current.Energy < best.Energy {
                        best = current
                    }
                }
                if acceptanceProb != 1 {
                    acceptanceProbHistory = append(acceptanceProbHistory, acceptanceProb)
                }
            }
        }

        fmt.Printf("Iteration %d: Swapped value at Table %d, Row %d, Col %d from %s to %s with value at Table %d, Row %d, Col %d from %s to %s (Acceptance Probability: %f)\n", 
            iteration+1, tableIdx1+1, rowIdx1+1, colIdx1+1, oldValue1, neighbor.Cube.Tables[tableIdx1][rowIdx1][colIdx1], 
            tableIdx2+1, rowIdx2+1, colIdx2+1, oldValue2, neighbor.Cube.Tables[tableIdx2][rowIdx2][colIdx2], acceptanceProb)
        fmt.Printf("Iteration %d: Current State Energy: %f, Temperature: %f\n", iteration+1, current.Energy, temp)

        if current.Energy == 0 {
            fmt.Println("Magic cube has been built!")
            break
        }

        // Check for energy stabilization
        if iteration > 0 && energyHistory[iteration-1] == current.Energy {
            stabilizationCount++
            if stabilizationCount >= stabilizationThreshold {
                fmt.Println("Terminating early due to energy stabilization.")
                break
            }
        } else {
            stabilizationCount = 0
        }

        newTemp := temp * (1 - coolingRate)
        if newTemp < temperatureThreshold {
            temp = temperatureThreshold
            fmt.Println("Temperature reached threshold, will terminate in next iteration.")
        } else {
            temp = newTemp
        }
        
        energyHistory = append(energyHistory, current.Energy)
        iteration++
    }

    duration := time.Since(startTime)
    formattedDuration := fmt.Sprintf("%.3f", duration.Seconds())

    return current, best, energyHistory, acceptanceProbHistory, formattedDuration, stuckCount, initialEnergy, iteration
}

func generateNeighbor(state State) (State, int, int, int, int, int, int, string, string) {
    neighbor := state
    tableIdx1 := rand.Intn(models.NumTables)
    rowIdx1 := rand.Intn(models.Rows)
    colIdx1 := rand.Intn(models.Cols)
    tableIdx2 := tableIdx1
    rowIdx2 := rowIdx1
    colIdx2 := colIdx1

    // Ensure the second position is different from the first
    for tableIdx1 == tableIdx2 && rowIdx1 == rowIdx2 && colIdx1 == colIdx2 {
        tableIdx2 = rand.Intn(models.NumTables)
        rowIdx2 = rand.Intn(models.Rows)
        colIdx2 = rand.Intn(models.Cols)
    }

    oldValue1 := neighbor.Cube.Tables[tableIdx1][rowIdx1][colIdx1]
    oldValue2 := neighbor.Cube.Tables[tableIdx2][rowIdx2][colIdx2]

    // Swap values
    neighbor.Cube.Tables[tableIdx1][rowIdx1][colIdx1] = oldValue2
    neighbor.Cube.Tables[tableIdx2][rowIdx2][colIdx2] = oldValue1

    neighbor.Energy = models.EvaluateIndividual(neighbor.Cube)
    return neighbor, tableIdx1, rowIdx1, colIdx1, tableIdx2, rowIdx2, colIdx2, oldValue1, oldValue2
}

func acceptanceProbability(currentEnergy, newEnergy, temperature float64) float64 {
    return math.Exp((currentEnergy - newEnergy) / temperature)
}