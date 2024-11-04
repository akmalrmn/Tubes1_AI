package simulatedannealing

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

type Individual struct {
	ID            int
	Tables        [][][]string
	ObjectiveFunc int
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
		}
	}

	return totalAbsSumColumns
}

func SumRows(tables [][][]string) int {
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
		}
	}

	return totalAbsSumRows
}

func SumPoles(tables [][][]string) int {
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
	}

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
		childTables := make([][][]string, 5)

		for j := 0; j < 3; j++ {
			childTables[j] = selected[i].Tables[j]
		}

		for j := 3; j < 5; j++ {
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
		tableIdx := rand.Intn(numTables)
		rowIdx := rand.Intn(rows)
		colIdx := rand.Intn(cols)
		point := [2]int{tableIdx, rowIdx*cols + colIdx}
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
			originalColIdx := point[1] % cols
			originalRowIdx := point[1] / cols

			var targetTableIdx, targetRowIdx, targetColIdx int
			for {
				targetTableIdx = rand.Intn(numTables)
				targetRowIdx = rand.Intn(rows)
				targetColIdx = rand.Intn(cols)

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

func Run() {
	var n, iterations int
	fmt.Println("Enter the number of individuals (population size): ")
	fmt.Scanln(&n)
	fmt.Println("Enter the number of iterations: ")
	fmt.Scanln(&iterations)

	startTime := time.Now()

	population := make([]Individual, n)
	for i := 0; i < n; i++ {
		tables := GenerateTable()
		objectiveFunctionValue := CalculateObjectiveFunction(tables)
		population[i] = Individual{
			ID:            i + 1,
			Tables:        tables,
			ObjectiveFunc: objectiveFunctionValue,
		}
	}

	var highestIndividual, lowestIndividual Individual
	var highestIter, lowestIter, highestValue, lowestValue int
	highestValue = -1
	lowestValue = int(^uint(0) >> 1)

	for it := 0; it < iterations; it++ {
		fmt.Printf("\nIteration %d:\n", it+1)

		for i := 0; i < n; i++ {
			fmt.Printf("\nIndividual %d:\n", population[i].ID)
			PrintTables(population[i].Tables)
			fmt.Printf("Objective Function Value for Individual %d: %d\n", population[i].ID, population[i].ObjectiveFunc)
		}

		objectiveValues := make([]int, n)
		for i := 0; i < n; i++ {
			objectiveValues[i] = population[i].ObjectiveFunc

			if objectiveValues[i] > highestValue {
				highestValue = objectiveValues[i]
				highestIndividual = population[i]
				highestIter = it + 1
			}
			if objectiveValues[i] < lowestValue {
				lowestValue = objectiveValues[i]
				lowestIndividual = population[i]
				lowestIter = it + 1
			}
		}

		fitnessValues := CalculateFitness(objectiveValues)
		fmt.Println()
		for i := 1; i <= n; i++ {
			fmt.Printf("Fitness Percentage for Individual %d: %.2f%%\n", i, fitnessValues[i-1])
		}

		selectedIndices := RouletteWheelSelection(fitnessValues, n)
		selectedParents := make([]Individual, n)
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
	}

	fmt.Printf("\nThe worst individual is %d in Iteration %d with objective function value %d.\n",
		highestIndividual.ID, highestIter, highestValue)
	fmt.Printf("The best individual is %d in Iteration %d with objective function value %d.\n",
		lowestIndividual.ID, lowestIter, lowestValue)

	elapsedTime := time.Since(startTime)
	fmt.Printf("\nDone executing %d population for %d iterations in %s\n", n, iterations, elapsedTime)

}
