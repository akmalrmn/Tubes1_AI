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
	numTables = 5 // Number of tables (layers of the cube)
)

// GenerateTable generates a 5x5 table with random values
func GenerateTable() [][]string {
	table := make([][]string, rows)
	for i := range table {
		table[i] = make([]string, cols)
		for j := range table[i] {
			// Generate random number for each cell
			table[i][j] = fmt.Sprintf("%d", rand.Intn(5)+1)
		}
	}
	return table
}

// prints multiple tables (e.g., 1A to 5A)
func PrintTables(tables [][][]string, view string) {
	// Print table headers
	for i := 0; i < numTables; i++ {
		fmt.Printf("Table %d%s\t\t", i+1, view)
	}
	fmt.Println()

	// Print rows of each table side by side
	for i := 0; i < rows; i++ {
		for t := 0; t < numTables; t++ {
			// Print one row of the current table
			for j := 0; j < cols; j++ {
				fmt.Printf("%-4s", tables[t][i][j]) // Adjust width with %-4s for tighter spacing
			}
			// Print space between tables
			fmt.Printf("\t")
		}
		fmt.Println()
	}
}

// SumColumns prints the sum of every column from table 1A to 5A
func SumColumns(tables [][][]string) {
	fmt.Println("\nColumn Sums:")
	for tableIdx := 0; tableIdx < numTables; tableIdx++ {
		for col := 0; col < cols; col++ {
			sum := 0
			for row := 0; row < rows; row++ {
				val, _ := strconv.Atoi(tables[tableIdx][row][col])
				sum += val
			}
			fmt.Printf("Col %d (Table %d): %d\n", col+1+(tableIdx*cols), tableIdx+1, sum)
		}
	}
}

// SumRows prints the sum of every row from table 1A to 5A
func SumRows(tables [][][]string) {
	fmt.Println("\nRow Sums:")
	for tableIdx := 0; tableIdx < numTables; tableIdx++ {
		for row := 0; row < rows; row++ {
			sum := 0
			for col := 0; col < cols; col++ {
				val, _ := strconv.Atoi(tables[tableIdx][row][col])
				sum += val
			}
			fmt.Printf("Row %d (Table %d): %d\n", row+1+(tableIdx*rows), tableIdx+1, sum)
		}
	}
}

// SumPoles prints the sum of each pole (1-25)
func SumPoles(tables [][][]string) {
	fmt.Println("\nPole Sums:")
	for pole := 0; pole < rows*cols; pole++ {
		sum := 0
		for tableIdx := 0; tableIdx < numTables; tableIdx++ {
			// Calculate row and column indices for the pole
			rowIdx := pole / cols
			colIdx := pole % cols
			val, _ := strconv.Atoi(tables[tableIdx][rowIdx][colIdx])
			sum += val
		}
		fmt.Printf("Pole %d: %d\n", pole+1, sum)
	}
}

// SumFaceDiagonal prints the sum of the face diagonals
func SumFaceDiagonal(tables [][][]string) {
	fmt.Println("\nFace Diagonal Sums:")

	// Face Diagonal Front
	// Diagonal from top-left to bottom-right
	sum1 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[0][i][i])
		sum1 += val
	}
	fmt.Printf("Face Diagonal Front (Top-left to Bottom-right): %d\n", sum1)

	// Diagonal from bottom-left to top-right
	sum2 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[0][rows-1-i][i])
		sum2 += val
	}
	fmt.Printf("Face Diagonal Front (Bottom-left to Top-right): %d\n", sum2)

	// Face Diagonal Back
	// Diagonal from top-left to bottom-right
	sum3 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[numTables-1][i][i])
		sum3 += val
	}
	fmt.Printf("Face Diagonal Back (Top-left to Bottom-right): %d\n", sum3)

	// Diagonal from bottom-left to top-right
	sum4 := 0
	for i := 0; i < rows; i++ {
		val, _ := strconv.Atoi(tables[numTables-1][rows-1-i][i])
		sum4 += val
	}
	fmt.Printf("Face Diagonal Back (Bottom-left to Top-right): %d\n", sum4)

	// Face Diagonal Top
	// From front to back
	sum5 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][0][i])
		sum5 += val
	}
	fmt.Printf("Face Diagonal Top (Front to Back): %d\n", sum5)

	// From back to front
	sum6 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][0][cols-1-i])
		sum6 += val
	}
	fmt.Printf("Face Diagonal Top (Back to Front): %d\n", sum6)

	// Face Diagonal Down
	// From front to back
	sum7 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1][0+i])
		sum7 += val
	}
	fmt.Printf("Face Diagonal Down (Front to Back): %d\n", sum7)

	// From back to front
	sum8 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1][cols-1-i])
		sum8 += val
	}
	fmt.Printf("Face Diagonal Down (Back to Front): %d\n", sum8)

	// Face Diagonal Left
	// From front to back
	sum9 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][i][0])
		sum9 += val
	}
	fmt.Printf("Face Diagonal Left (Front to Back): %d\n", sum9)

	// From back to front
	sum10 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1-i][0])
		sum10 += val
	}
	fmt.Printf("Face Diagonal Left (Back to Front): %d\n", sum10)

	// Face Diagonal Right
	// From front to back
	sum11 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][i][cols-1])
		sum11 += val
	}
	fmt.Printf("Face Diagonal Right (Front to Back): %d\n", sum11)

	// From back to front
	sum12 := 0
	for i := 0; i < numTables; i++ {
		val, _ := strconv.Atoi(tables[i][rows-1-i][cols-1])
		sum12 += val
	}
	fmt.Printf("Face Diagonal Right (Back to Front): %d\n", sum12)
}

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate multiple random 5x5 tables (representing the front view: 1A to 5A)
	tables := make([][][]string, numTables)
	for i := 0; i < numTables; i++ {
		tables[i] = GenerateTable()
	}

	// Print the front view tables (1A to 5A)
	fmt.Println("Front View (Tables 1A to 5A):")
	PrintTables(tables, "A")

	// Sum columns, rows, and poles
	SumColumns(tables)
	SumRows(tables)
	SumPoles(tables)

	// Sum face diagonals
	SumFaceDiagonal(tables)
}
