package main

import (
	"fmt"
	"math/rand"
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
			table[i][j] = fmt.Sprintf("%d", rand.Intn(25)+1)
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

// GetTopView returns the top-down view tables (1B to 5B) from the 3D cube
// It takes the rows from the original cube and rotates them to form the top view.
func GetTopView(tables [][][]string) [][][]string {
	topView := make([][][]string, rows)
	for i := 0; i < rows; i++ {
		topView[i] = make([][]string, rows)
		for j := 0; j < rows; j++ {
			topView[i][j] = make([]string, cols)
			for k := 0; k < cols; k++ {
				// Mapping rows from A tables into B tables
				topView[i][j][k] = tables[rows-1-j][i][k]
			}
		}
	}
	return topView
}

// GetSideView returns the side view tables (1C to 5C) from the 3D cube
// It takes the columns of the original cube and rotates them to form the side view.
func GetSideView(tables [][][]string) [][][]string {
	sideView := make([][][]string, rows)
	for i := 0; i < rows; i++ {
		sideView[i] = make([][]string, rows)
		for j := 0; j < rows; j++ {
			sideView[i][j] = make([]string, cols)
			for k := 0; k < cols; k++ {
				// Mapping columns from A tables into C tables
				sideView[i][j][k] = tables[k][j][i]
			}
		}
	}
	return sideView
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

	// Generate the top-down view tables (1B to 5B)
	topViewTables := GetTopView(tables)

	// Print the top-down view tables (1B to 5B)
	fmt.Println("\nTop-Down View (Tables 1B to 5B):")
	PrintTables(topViewTables, "B")

	// Generate the side view tables (1C to 5C)
	sideViewTables := GetSideView(tables)

	// Print the side view tables (1C to 5C)
	fmt.Println("\nSide View (Tables 1C to 5C):")
	PrintTables(sideViewTables, "C")
}
