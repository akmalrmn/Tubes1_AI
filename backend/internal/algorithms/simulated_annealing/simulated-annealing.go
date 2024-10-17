package simulated_annealing

import (
    "fmt"
    "magic-cube-solver/internal/models"
    "math"
    "math/rand"
    "strconv"
    "time"

    "gonum.org/v1/plot"
    "gonum.org/v1/plot/plotter"
    "gonum.org/v1/plot/vg"
)

type State struct {
    Cube   models.Cube
    Energy float64
}

func SimulatedAnnealing(initialTemp, coolingRate float64, maxIterations int) (State, State, []float64, []float64, time.Duration, int) {
    startTime := time.Now()
    current := State{
        Cube:   models.GenerateCube(),
        Energy: 0,
    }
    current.Energy = -models.EvaluateIndividual(flattenCube(current.Cube))

    best := current

    temp := initialTemp
    energyHistory := []float64{current.Energy}
    tempHistory := []float64{temp}
    stuckCount := 0

    for i := 0; i < maxIterations; i++ {
        neighbor := generateNeighbor(current)
        neighborEnergy := -models.EvaluateIndividual(flattenCube(neighbor.Cube))

        if acceptanceProbability(current.Energy, neighborEnergy, temp) > rand.Float64() {
            current = State{Cube: neighbor.Cube, Energy: neighborEnergy}

            if current.Energy < best.Energy {
                best = current
                stuckCount = 0
            } else {
                stuckCount++
            }
        } else {
            stuckCount++
        }

        temp *= 1 - coolingRate
        energyHistory = append(energyHistory, current.Energy)
        tempHistory = append(tempHistory, temp)
    }

    duration := time.Since(startTime)

    return current, best, energyHistory, tempHistory, duration, stuckCount
}

func generateNeighbor(state State) State {
    neighbor := state
    tableIdx := rand.Intn(models.NumTables)
    rowIdx := rand.Intn(models.Rows)
    colIdx := rand.Intn(models.Cols)
    newValue := strconv.Itoa(rand.Intn(5) + 1)
    neighbor.Cube.Tables[tableIdx][rowIdx][colIdx] = newValue
    return neighbor
}

func acceptanceProbability(currentEnergy, newEnergy, temperature float64) float64 {
    if newEnergy < currentEnergy {
        return 1.0
    }
    return math.Exp((currentEnergy - newEnergy) / temperature)
}

func flattenCube(cube models.Cube) []int {
    flattened := make([]int, 0, models.NumTables*models.Rows*models.Cols)
    for _, table := range cube.Tables {
        for _, row := range table {
            for _, val := range row {
                intVal, _ := strconv.Atoi(val)
                flattened = append(flattened, intVal)
            }
        }
    }
    return flattened
}

func plotHistory(data []float64, filename, xlabel, ylabel string) {
    p := plot.New()

    p.Title.Text = fmt.Sprintf("%s over Iterations", ylabel)
    p.X.Label.Text = xlabel
    p.Y.Label.Text = ylabel

    pts := make(plotter.XYs, len(data))
    for i := range pts {
        pts[i].X = float64(i)
        pts[i].Y = data[i]
    }

    line, err := plotter.NewLine(pts)
    if err != nil {
        panic(err)
    }
    p.Add(line)

    if err := p.Save(4*vg.Inch, 4*vg.Inch, filename); err != nil {
        panic(err)
    }
}