package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	rows      = 5
	cols      = 5
	numTables = 5
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GenerateTable() [][][]string {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, 125)
	for i := 0; i < 125; i++ {
		numbers[i] = i + 1
	}

	rand.Shuffle(len(numbers), func(i, j int) {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	})

	tables := make([][][]string, numTables)
	counter := 0
	for t := 0; t < numTables; t++ {
		table := make([][]string, rows)
		for i := range table {
			table[i] = make([]string, cols)
			for j := range table[i] {
				table[i][j] = fmt.Sprintf("%d", numbers[counter])
				counter++
			}
		}
		tables[t] = table
	}

	return tables
}

func PrintTables(tables [][][]string) {
	for i := 0; i < rows; i++ {
		for t := 0; t < numTables; t++ {
			for j := 0; j < cols; j++ {
				fmt.Printf("%-4s", tables[t][i][j])
			}
			fmt.Printf("\t")
		}
		fmt.Println()
	}
}

func GetTopView(tables [][][]string) [][][]string {
	topView := make([][][]string, rows)
	for i := 0; i < rows; i++ {
		topView[i] = make([][]string, rows)
		for j := 0; j < rows; j++ {
			topView[i][j] = make([]string, cols)
			for k := 0; k < cols; k++ {
				topView[i][j][k] = tables[rows-1-j][i][k]
			}
		}
	}
	return topView
}

func GetSideView(tables [][][]string) [][][]string {
	sideView := make([][][]string, rows)
	for i := 0; i < rows; i++ {
		sideView[i] = make([][]string, rows)
		for j := 0; j < rows; j++ {
			sideView[i][j] = make([]string, cols)
			for k := 0; k < cols; k++ {
				sideView[i][j][k] = tables[k][j][i]
			}
		}
	}
	return sideView
}

func SumColumns(tables [][][]string) int {
	totalAbsSumColumns := 0

	for tableIdx := 0; tableIdx < numTables; tableIdx++ {
		for col := 0; col < cols; col++ {
			sum := 0
			for row := 0; row < rows; row++ {
				val, _ := strconv.Atoi(tables[tableIdx][row][col])
				sum += val
			}
			abs_sum := abs(315 - sum)
			totalAbsSumColumns += abs_sum
			// fmt.Printf("Col %d (Table %d): %d %d\n", col+1+(tableIdx*cols), tableIdx+1, sum, abs_sum)
		}
	}

	// fmt.Printf("totalAbsSumColumns: %d\n", totalAbsSumColumns)
	return totalAbsSumColumns
}

func SumRows(tables [][][]string) int {
	// fmt.Println("\nRow Sums:")
	totalAbsSumRows := 0

	for tableIdx := 0; tableIdx < numTables; tableIdx++ {
		for row := 0; row < rows; row++ {
			sum := 0
			for col := 0; col < cols; col++ {
				val, _ := strconv.Atoi(tables[tableIdx][row][col])
				sum += val
			}
			abs_sum := abs(315 - sum)
			totalAbsSumRows += abs_sum
			// fmt.Printf("Row %d (Table %d): %d %d\n", row+1+(tableIdx*rows), tableIdx+1, sum, abs_sum)
		}
	}

	// fmt.Printf("totalAbsSumRows: %d\n", totalAbsSumRows)
	return totalAbsSumRows
}

func SumPoles(tables [][][]string) int {
	// fmt.Println("\nPole Sums:")
	totalAbsSumPoles := 0

	for pole := 0; pole < rows*cols; pole++ {
		sum := 0
		for tableIdx := 0; tableIdx < numTables; tableIdx++ {
			rowIdx := pole / cols
			colIdx := pole % cols
			val, _ := strconv.Atoi(tables[tableIdx][rowIdx][colIdx])
			sum += val
		}
		abs_sum := abs(315 - sum)
		totalAbsSumPoles += abs_sum
		// fmt.Printf("Pole %d: %d %d\n", pole+1, sum, abs_sum)
	}

	// fmt.Printf("totalAbsSumPoles: %d\n", totalAbsSumPoles)
	return totalAbsSumPoles
}

func SumFaceDiagonal(tables [][][]string) int {
	// fmt.Println("\nFace Diagonal Sums:")
	totalAbsSumFaceDiagonal := 0

	// Face Diagonal Front
	// Diagonal from top-left to bottom-right
	sum1 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[0][i][i])
		sum1 += val
	}
	abs_sum1 := abs(315 - sum1)
	totalAbsSumFaceDiagonal += abs_sum1
	// fmt.Printf("Face Diagonal Front (Top-left to Bottom-right): %d %d\n", sum1, abs_sum1)

	// Diagonal from bottom-left to top-right
	sum2 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[0][rows-1-i][i])
		sum2 += val
	}
	abs_sum2 := abs(315 - sum2)
	totalAbsSumFaceDiagonal += abs_sum2
	// fmt.Printf("Face Diagonal Front (Bottom-left to Top-right): %d %d\n", sum2, abs_sum2)

	// Face Diagonal Back
	// Diagonal from top-left to bottom-right
	sum3 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[numTables-1][i][i])
		sum3 += val
	}
	abs_sum3 := abs(315 - sum3)
	totalAbsSumFaceDiagonal += abs_sum3
	// fmt.Printf("Face Diagonal Back (Top-left to Bottom-right): %d %d\n", sum3, abs_sum3)

	// Diagonal from bottom-left to top-right
	sum4 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[numTables-1][rows-1-i][i])
		sum4 += val
	}
	abs_sum4 := abs(315 - sum4)
	totalAbsSumFaceDiagonal += abs_sum4
	// fmt.Printf("Face Diagonal Back (Bottom-left to Top-right): %d %d\n", sum4, abs_sum4)

	// Face Diagonal Top
	// From front to back
	sum5 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][0][i])
		sum5 += val
	}
	abs_sum5 := abs(315 - sum5)
	totalAbsSumFaceDiagonal += abs_sum5
	// fmt.Printf("Face Diagonal Top (Front to Back): %d %d\n", sum5, abs_sum5)

	// From back to front
	sum6 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][0][cols-1-i])
		sum6 += val
	}
	abs_sum6 := abs(315 - sum6)
	totalAbsSumFaceDiagonal += abs_sum6
	// fmt.Printf("Face Diagonal Top (Back to Front): %d %d\n", sum6, abs_sum6)

	// Face Diagonal Down
	// From front to back
	sum7 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1][0+i])
		sum7 += val
	}
	abs_sum7 := abs(315 - sum7)
	totalAbsSumFaceDiagonal += abs_sum7
	// fmt.Printf("Face Diagonal Down (Front to Back): %d %d\n", sum7, abs_sum7)

	// From back to front
	sum8 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1][cols-1-i])
		sum8 += val
	}
	abs_sum8 := abs(315 - sum8)
	totalAbsSumFaceDiagonal += abs_sum8
	// fmt.Printf("Face Diagonal Down (Back to Front): %d %d\n", sum8, abs_sum8)

	// Face Diagonal Left
	// From front to back
	sum9 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][i][0])
		sum9 += val
	}
	abs_sum9 := abs(315 - sum9)
	totalAbsSumFaceDiagonal += abs_sum9
	// fmt.Printf("Face Diagonal Left (Front to Back): %d %d\n", sum9, abs_sum9)

	// From back to front
	sum10 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1-i][0])
		sum10 += val
	}
	abs_sum10 := abs(315 - sum10)
	totalAbsSumFaceDiagonal += abs_sum10
	// fmt.Printf("Face Diagonal Left (Back to Front): %d %d\n", sum10, abs_sum10)

	// Face Diagonal Right
	// From front to back
	sum11 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][i][cols-1])
		sum11 += val
	}
	abs_sum11 := abs(315 - sum11)
	totalAbsSumFaceDiagonal += abs_sum11
	// fmt.Printf("Face Diagonal Right (Front to Back): %d %d\n", sum11, abs_sum11)

	// From back to front
	sum12 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1-i][cols-1])
		sum12 += val
	}
	abs_sum12 := abs(315 - sum12)
	totalAbsSumFaceDiagonal += abs_sum12
	// fmt.Printf("Face Diagonal Right (Back to Front): %d %d\n", sum12, abs_sum12)

	// Return total abs_sum
	// fmt.Printf("totalAbsSumFaceDiagonal: %d\n", totalAbsSumFaceDiagonal)
	return totalAbsSumFaceDiagonal
}

func SumSpaceDiagonal(tables [][][]string) int {
	// fmt.Println("\nSpace Diagonal Sums:")
	totalAbsSumSpaceDiagonal := 0

	// Space Diagonal 1
	sum1 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][i][i]) // row (i) col (i)
		sum1 += val
	}
	abs_sum1 := abs(315 - sum1)
	totalAbsSumSpaceDiagonal += abs_sum1
	// fmt.Printf("Space Diagonal 1: %d %d\n", sum1, abs_sum1)

	// Space Diagonal 2
	sum2 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][i][cols-1-i]) // row (i) col (cols-1-i)
		sum2 += val
	}
	abs_sum2 := abs(315 - sum2)
	totalAbsSumSpaceDiagonal += abs_sum2
	// fmt.Printf("Space Diagonal 2: %d %d\n", sum2, abs_sum2)

	// Space Diagonal 3
	sum3 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1-i][i]) // row (rows-1-i) col (i)
		sum3 += val
	}
	abs_sum3 := abs(315 - sum3)
	totalAbsSumSpaceDiagonal += abs_sum3
	// fmt.Printf("Space Diagonal 3: %d %d\n", sum3, abs_sum3)

	// Space Diagonal 4
	sum4 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1-i][cols-1-i]) // row (rows-1-i) col (cols-1)
		sum4 += val
	}
	abs_sum4 := abs(315 - sum4)
	totalAbsSumSpaceDiagonal += abs_sum4
	// fmt.Printf("Space Diagonal 4: %d %d\n", sum4, abs_sum4)

	// Print and return the total abs_sum
	// fmt.Printf("totalAbsSumSpaceDiagonal: %d\n", totalAbsSumSpaceDiagonal)
	return totalAbsSumSpaceDiagonal
}

func CalculateObjectiveFunction(tables [][][]string) int {
	totalAbsSumColumns := SumColumns(tables)
	totalAbsSumRows := SumRows(tables)
	totalAbsSumPoles := SumPoles(tables)
	totalAbsSumFaceDiagonal := SumFaceDiagonal(tables)
	totalAbsSumSpaceDiagonal := SumSpaceDiagonal(tables)

	totalAbsSum := totalAbsSumColumns + totalAbsSumRows + totalAbsSumPoles +
		totalAbsSumFaceDiagonal + totalAbsSumSpaceDiagonal

	// fmt.Printf("Total Objective Function: %d\n", totalAbsSum)
	return totalAbsSum
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

func RouletteWheelSelection(fitnessValues []float64, numParents int) []int {
	selected := make([]int, numParents)
	rand.Seed(time.Now().UnixNano())

	cumulative := make([]float64, len(fitnessValues))
	cumulative[0] = fitnessValues[0]
	for i := 1; i < len(fitnessValues); i++ {
		cumulative[i] = cumulative[i-1] + fitnessValues[i]
	}

	for i := 0; i < numParents; i++ {
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

func main() {
	var n int
	fmt.Println("Enter the number of individuals (population size): ")
	fmt.Scanln(&n)

	objectiveValues := make([]int, n)
	fitnessValues := make([]float64, n)

	for i := 1; i <= n; i++ {
		fmt.Printf("\nIndividual %d:\n", i)
		tables := GenerateTable()
		PrintTables(tables)
		objectiveFunctionValue := CalculateObjectiveFunction(tables)
		objectiveValues[i-1] = objectiveFunctionValue
		fmt.Printf("Objective Function Value for Individual %d: %d\n", i, objectiveFunctionValue)
	}

	fitnessValues = CalculateFitness(objectiveValues)
	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("Fitness Percentage for Individual %d: %.2f%%\n", i, fitnessValues[i-1])
	}

	var numParents int
	if n < 4 {
		numParents = 2
	} else {
		numParents = (n + 1) / 2
	}

	selectedParents := RouletteWheelSelection(fitnessValues, numParents)

	fmt.Printf("\nSelected Individuals: ")
	for i, idx := range selectedParents {
		if i != 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%d", idx+1)
	}
	fmt.Println()

	for _, idx := range selectedParents {
		fmt.Printf("Individual %d with objective function value %d selected.\n", idx+1, objectiveValues[idx])
	}
}
