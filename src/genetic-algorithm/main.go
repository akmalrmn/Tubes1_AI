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

	// Print the column sums
	SumColumns(tables)

	// Print the row sums
	SumRows(tables)

	// Print the pole sums
	SumPoles(tables)

	// // Generate the top-down view tables (1B to 5B)
	// topViewTables := GetTopView(tables)

	// // Print the top-down view tables (1B to 5B)
	// fmt.Println("\nTop-Down View (Tables 1B to 5B):")
	// PrintTables(topViewTables, "B")

	// // Generate the side view tables (1C to 5C)
	// sideViewTables := GetSideView(tables)

	// // Print the side view tables (1C to 5C)
	// fmt.Println("\nSide View (Tables 1C to 5C):")
	// PrintTables(sideViewTables, "C")
}
