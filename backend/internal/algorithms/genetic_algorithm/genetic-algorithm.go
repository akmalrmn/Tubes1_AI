package genetic_algorithm

import (
    "fmt"
    "magic-cube-solver/internal/models"
    "math/rand"
    "time"
)

type Individual struct {
    ID            int
    Tables        [][][]string
    ObjectiveFunc int
}

func GenerateTable() [][][]string {
    return models.GenerateTable()
}

func PrintTables(tables [][][]string) {
    for i := 0; i < models.Rows; i++ {
        for t := 0; t < models.NumTables; t++ {
            for j := 0; j < models.Cols; j++ {
                fmt.Printf("%-4s", tables[t][i][j])
            }
            fmt.Printf("\t")
        }
        fmt.Println()
    }
}

func CalculateObjectiveFunction(tables [][][]string) int {
    return int(models.EvaluateIndividual(models.Cube{Tables: tables}))
}

func CalculateFitness(objectiveValues []int) []float64 {
    fitnessValues := make([]float64, len(objectiveValues))
    for i, obj := range objectiveValues {
        if obj > 0 {
            fitnessValues[i] = 1.0 / float64(obj)
        }
    }

    totalFitness := 0.0
    for _, fitness := range fitnessValues {
        totalFitness += fitness
    }

    for i := range fitnessValues {
        fitnessValues[i] = fitnessValues[i] / totalFitness * 100.0
    }

    return fitnessValues
}

func RouletteWheelSelection(fitnessValues []float64, numSelections int) []int {
    selected := make([]int, numSelections)
    rand.Seed(time.Now().UnixNano())

    cumulative := make([]float64, len(fitnessValues))
    cumulative[0] = fitnessValues[0]
    for i := 1; i < len(fitnessValues); i++ {
        cumulative[i] = cumulative[i-1] + fitnessValues[i]
    }

    for i := 0; i < numSelections; i++ {
        r := rand.Float64() * 100.0
        for j := range cumulative {
            if r <= cumulative[j] {
                selected[i] = j
                break
            }
        }
    }

    return selected
}

func Crossover(selected []Individual) []Individual {
    var children []Individual

    for i := 0; i < len(selected); i++ {
        childTables := make([][][]string, models.NumTables)

        for j := 0; j < 3; j++ {
            childTables[j] = selected[i].Tables[j]
        }

        for j := 3; j < models.NumTables; j++ {
            randIdx := i
            for randIdx == i {
                randIdx = rand.Intn(len(selected))
            }
            childTables[j] = selected[randIdx].Tables[j]
        }

        child := Individual{
            ID:            i + 1,
            Tables:        childTables,
            ObjectiveFunc: 0,
        }
        children = append(children, child)
    }

    return children
}

func Mutation(children []Individual) []Individual {
    rand.Seed(time.Now().UnixNano())
    mutationTimes := rand.Intn(6) + 10 // Random number between 10 and 15

    mutationPoints := make(map[[2]int]struct{})
    for len(mutationPoints) < mutationTimes {
        tableIdx := rand.Intn(models.NumTables)
        rowIdx := rand.Intn(models.Rows)
        colIdx := rand.Intn(models.Cols)
        point := [2]int{tableIdx, rowIdx*models.Cols + colIdx}
        mutationPoints[point] = struct{}{}
    }

    pointsSlice := make([][2]int, 0, len(mutationPoints))
    for point := range mutationPoints {
        pointsSlice = append(pointsSlice, point)
    }

    for i, child := range children {
        fmt.Printf("Mutating Child %d:\n", child.ID)

        for _, point := range pointsSlice {
            sourceTableIdx := point[0]
            originalColIdx := point[1] % models.Cols
            originalRowIdx := point[1] / models.Cols

            var targetTableIdx, targetRowIdx, targetColIdx int
            for {
                targetTableIdx = rand.Intn(models.NumTables)
                targetRowIdx = rand.Intn(models.Rows)
                targetColIdx = rand.Intn(models.Cols)

                if targetTableIdx != sourceTableIdx ||
                    (targetTableIdx == sourceTableIdx && (targetRowIdx != originalRowIdx || targetColIdx != originalColIdx)) {
                    break
                }
            }

            sourceValue := child.Tables[sourceTableIdx][originalRowIdx][originalColIdx]
            targetValue := child.Tables[targetTableIdx][targetRowIdx][targetColIdx]

            child.Tables[sourceTableIdx][originalRowIdx][originalColIdx] = targetValue
            child.Tables[targetTableIdx][targetRowIdx][targetColIdx] = sourceValue

            fmt.Printf("Swapped value %s at Table %d Position (%d, %d) with value %s at Table %d Position (%d, %d)\n",
                sourceValue, sourceTableIdx+1, originalRowIdx+1, originalColIdx+1,
                targetValue, targetTableIdx+1, targetRowIdx+1, targetColIdx+1)
        }

        child.ObjectiveFunc = CalculateObjectiveFunction(child.Tables)
        fmt.Printf("Objective Function Value for Mutated Child %d: %d\n\n", child.ID, child.ObjectiveFunc)

        fmt.Printf("Mutated Child %d\n", child.ID)
        for tableIdx, table := range child.Tables {
            fmt.Printf("Table %d:\n", tableIdx+1)
            for _, row := range table {
                fmt.Println(row)
            }
            fmt.Println()
        }

        children[i] = child
    }

    return children
}

func RunGeneticAlgorithm(populationSize, generations int) (Individual, Individual, []float64, time.Duration, int) {
    startTime := time.Now()

    population := make([]Individual, populationSize)
    for i := 0; i < populationSize; i++ {
        tables := GenerateTable()
        objectiveFunctionValue := CalculateObjectiveFunction(tables)
        population[i] = Individual{
            ID:            i + 1,
            Tables:        tables,
            ObjectiveFunc: objectiveFunctionValue,
        }
    }

    var highestIndividual, lowestIndividual Individual
    var highestValue, lowestValue int
    highestValue = -1
    lowestValue = int(^uint(0) >> 1)

    objectiveHistory := make([]float64, generations)

    for it := 0; it < generations; it++ {
        fmt.Printf("\nIteration %d:\n", it+1)

        for i := 0; i < populationSize; i++ {
            fmt.Printf("\nIndividual %d:\n", population[i].ID)
            PrintTables(population[i].Tables)
            fmt.Printf("Objective Function Value for Individual %d: %d\n", population[i].ID, population[i].ObjectiveFunc)
        }

        objectiveValues := make([]int, populationSize)
        for i := 0; i < populationSize; i++ {
            objectiveValues[i] = population[i].ObjectiveFunc

            if objectiveValues[i] > highestValue {
                highestValue = objectiveValues[i]
                highestIndividual = population[i]
            }
            if objectiveValues[i] < lowestValue {
                lowestValue = objectiveValues[i]
                lowestIndividual = population[i]
            }
        }

        fitnessValues := CalculateFitness(objectiveValues)
        fmt.Println()
        for i := 1; i <= populationSize; i++ {
            fmt.Printf("Fitness Percentage for Individual %d: %.2f%%\n", i, fitnessValues[i-1])
        }

        selectedIndices := RouletteWheelSelection(fitnessValues, populationSize)
        selectedParents := make([]Individual, populationSize)
        fmt.Printf("\nSelected Individuals: ")
        for i, idx := range selectedIndices {
            selectedParents[i] = population[idx]
            if i != 0 {
                fmt.Print(", ")
            }
            fmt.Printf("%d", idx+1)
        }
        fmt.Println()

        for _, idx := range selectedIndices {
            fmt.Printf("Individual %d with objective function value %d selected.\n", idx+1, objectiveValues[idx])
        }

        fmt.Println()
        children := Crossover(selectedParents)
        population = Mutation(children)

        // Record the best objective function value for this generation
        objectiveHistory[it] = float64(lowestValue)
    }

    elapsedTime := time.Since(startTime)
    fmt.Printf("\nDone executing %d population for %d iterations in %s\n", populationSize, generations, elapsedTime)

    return highestIndividual, lowestIndividual, objectiveHistory, elapsedTime, generations
}