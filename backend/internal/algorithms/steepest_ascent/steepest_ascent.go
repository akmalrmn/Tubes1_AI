// package main
package steepest_ascent

import (
	"fmt"
	"magic-cube-solver/internal/models"
	"strconv"
	"time"
)

type State struct {
	Cube           models.Cube
	ObjectiveValue float64
}

func SteepestAscent(maxIterations int, targetSum int) (State, State, time.Duration) {
	startTime := time.Now()

	initial := State{
		Cube:           models.GenerateCube(),
		ObjectiveValue: 0,
	}
	initial.ObjectiveValue = -models.EvaluateIndividual(flattenCube(initial.Cube))

	current := initial

	for i := 0; i < maxIterations; i++ {
		bestNeighbor, improved := findBestNeighbor(current)

		if !improved {
			break
		}

		current = bestNeighbor

		fmt.Printf("\nIteration %d: Objective Value = %.0f\n", i+1, current.ObjectiveValue)
	}

	duration := time.Since(startTime)

	return initial, current, duration
}

func findBestNeighbor(current State) (State, bool) {
	bestNeighbor := current
	improved := false

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

							neighborObjectiveValue := -models.EvaluateIndividual(flattenCube(neighborCube))

							if neighborObjectiveValue < bestNeighbor.ObjectiveValue {
								bestNeighbor = State{Cube: neighborCube, ObjectiveValue: neighborObjectiveValue}
								improved = true

								fmt.Printf("   Swapping positions:\n")
								fmt.Printf("  	  Position 1: Table %d, Row %d, Column %d (Value %s)\n", tableIdx1, rowIdx1, colIdx1, val1)
								fmt.Printf("  	  Position 2: Table %d, Row %d, Column %d (Value %s)\n", tableIdx2, rowIdx2, colIdx2, val2)
								fmt.Printf("  	  New Objective Value: %.0f\n", neighborObjectiveValue)
							}
						}
					}
				}
			}
		}
	}

	return bestNeighbor, improved
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

// func main() {
// 	maxIterations := 1000
// 	targetSum := 25
// 	initialState, finalState, duration := SteepestAscent(maxIterations, targetSum)

// 	fmt.Println("\nInitial Cube State:")
// 	for _, table := range initialState.Cube.Tables {
// 		for _, row := range table {
// 			fmt.Println(row)
// 		}
// 		fmt.Println()
// 	}

// 	fmt.Println("Final Cube State:")
// 	for _, table := range finalState.Cube.Tables {
// 		for _, row := range table {
// 			fmt.Println(row)
// 		}
// 		fmt.Println()
// 	}

// 	fmt.Println("Final Sums:")

// 	columnSums := models.SumColumns(finalState.Cube.Tables)
// 	rowSums := models.SumRows(finalState.Cube.Tables)
// 	poleSums := models.SumPoles(finalState.Cube.Tables)
// 	faceDiagonalSums := models.SumFaceDiagonal(finalState.Cube.Tables)
// 	spaceDiagonalSums := models.SumSpaceDiagonal(finalState.Cube.Tables)

// 	fmt.Printf("Column Sums: %v\n", columnSums)
// 	fmt.Printf("Row Sums: %v\n", rowSums)
// 	fmt.Printf("Pole Sums: %v\n", poleSums)
// 	fmt.Printf("Face Diagonal Sums: %v\n", faceDiagonalSums)
// 	fmt.Printf("Space Diagonal Sums: %v\n", spaceDiagonalSums)

// 	allSumsMatch := true
// 	for _, sum := range append(append(append(append(columnSums, rowSums...), poleSums...), faceDiagonalSums...), spaceDiagonalSums...) {
// 		if sum != targetSum {
// 			allSumsMatch = false
// 			break
// 		}
// 	}

// 	if allSumsMatch {
// 		fmt.Println("\nThe final cube configuration is correct and meets the magic number requirements.")
// 	} else {
// 		fmt.Println("\nThe final cube configuration does not fully meet the magic number requirements.\nReached maximum iterations or a local optimum.")
// 	}

// 	fmt.Printf("\nFinal Objective Function Value: %.0f\n", finalState.ObjectiveValue)
// 	fmt.Printf("Search Duration: %v\n", duration)
// }
