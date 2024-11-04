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
        Cube:   copyCube(initialCube),  
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

        // Generate a neighbor state
        neighbor := generateNeighbor(current)
        neighborEnergy := models.EvaluateIndividual(neighbor.Cube)
        acceptanceProb := 0.0

        // Check if the neighbor has a lower energy (better solution)
        if neighborEnergy < current.Energy {
            current = State{Cube: copyCube(neighbor.Cube), Energy: neighborEnergy}
            if current.Energy < best.Energy {
                best = current
            }
        } else {
            stuckCount++
            
            if temp >= temperatureThreshold {
                acceptanceProb = acceptanceProbability(current.Energy, neighborEnergy, temp)
                if acceptanceProb > rand.Float64() {
                    current = State{Cube: copyCube(neighbor.Cube), Energy: neighborEnergy}
                    if current.Energy < best.Energy {
                        best = current
                    }
                }
                acceptanceProbHistory = append(acceptanceProbHistory, acceptanceProb)
            }
        }

        if current.Energy == 0 {
            fmt.Println("Magic cube has been built!")
            break
        }

        if iteration > 0 && energyHistory[iteration-1] == current.Energy {
            stabilizationCount++
            if stabilizationCount >= stabilizationThreshold {
                fmt.Println("Terminating early due to energy stabilization.")
                break
            }
        } else {
            stabilizationCount = 0
        }

        // Update temperature and energy history
        temp *= (1 - coolingRate)
        energyHistory = append(energyHistory, current.Energy)
        iteration++
    }

    duration := time.Since(startTime)
    formattedDuration := fmt.Sprintf("%.3f", duration.Seconds())

    return State{Cube: copyCube(initialCube), Energy: initialEnergy}, best, energyHistory, acceptanceProbHistory, formattedDuration, stuckCount, initialEnergy, iteration
}

// Generate a new neighbor by swapping two random elements in the cube
func generateNeighbor(state State) State {
    neighbor := State{
        Cube:   copyCube(state.Cube),
        Energy: state.Energy,
    }
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

    // Recalculate energy for the new configuration
    neighbor.Energy = models.EvaluateIndividual(neighbor.Cube)
    return neighbor
}

// Helper function to create a deep copy of the cube
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

// Calculate the acceptance probability for simulated annealing
func acceptanceProbability(currentEnergy, newEnergy, temperature float64) float64 {
    return math.Exp((currentEnergy - newEnergy) / temperature)
}
