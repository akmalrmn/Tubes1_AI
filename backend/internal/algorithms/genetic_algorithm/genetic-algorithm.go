package genetic_algorithm

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

// Cube represents a 3D cube
type Cube struct {
    Tables [][][]string
}

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

// GenerateCube generates a 3D cube with random values
func GenerateCube() Cube {
    rand.Seed(time.Now().UnixNano())
    tables := make([][][]string, numTables)
    for i := 0; i < numTables; i++ {
        tables[i] = GenerateTable()
    }
    return Cube{Tables: tables}
}

// SumColumns calculates the sum of every column from table 1A to 5A
func SumColumns(tables [][][]string) []int {
    sums := make([]int, cols*numTables)
    for tableIdx := 0; tableIdx < numTables; tableIdx++ {
        for col := 0; col < cols; col++ {
            sum := 0
            for row := 0; row < rows; row++ {
                val, _ := strconv.Atoi(tables[tableIdx][row][col])
                sum += val
            }
            sums[tableIdx*cols+col] = sum
        }
    }
    return sums
}

// SumRows calculates the sum of every row from table 1A to 5A
func SumRows(tables [][][]string) []int {
    sums := make([]int, rows*numTables)
    for tableIdx := 0; tableIdx < numTables; tableIdx++ {
        for row := 0; row < rows; row++ {
            sum := 0
            for col := 0; col < cols; col++ {
                val, _ := strconv.Atoi(tables[tableIdx][row][col])
                sum += val
            }
            sums[tableIdx*rows+row] = sum
        }
    }
    return sums
}

// SumPoles calculates the sum of each pole (1-25)
func SumPoles(tables [][][]string) []int {
    sums := make([]int, rows*cols)
    for pole := 0; pole < rows*cols; pole++ {
        sum := 0
        for tableIdx := 0; tableIdx < numTables; tableIdx++ {
            // Calculate row and column indices for the pole
            rowIdx := pole / cols
            colIdx := pole % cols
            val, _ := strconv.Atoi(tables[tableIdx][rowIdx][colIdx])
            sum += val
        }
        sums[pole] = sum
    }
    return sums
}

// SumFaceDiagonal calculates the sum of the face diagonals
func SumFaceDiagonal(tables [][][]string) []int {
    sums := make([]int, 12)

    // Face Diagonal Front
    // Diagonal from top-left to bottom-right
    sum1 := 0
    for i := 0; i < rows; i++ {
        val, _ := strconv.Atoi(tables[0][i][i])
        sum1 += val
    }
    sums[0] = sum1

    // Diagonal from bottom-left to top-right
    sum2 := 0
    for i := 0; i < rows; i++ {
        val, _ := strconv.Atoi(tables[0][rows-1-i][i])
        sum2 += val
    }
    sums[1] = sum2

    // Face Diagonal Back
    // Diagonal from top-left to bottom-right
    sum3 := 0
    for i := 0; i < rows; i++ {
        val, _ := strconv.Atoi(tables[numTables-1][i][i])
        sum3 += val
    }
    sums[2] = sum3

    // Diagonal from bottom-left to top-right
    sum4 := 0
    for i := 0; i < rows; i++ {
        val, _ := strconv.Atoi(tables[numTables-1][rows-1-i][i])
        sum4 += val
    }
    sums[3] = sum4

    // Face Diagonal Top
    // From front to back
    sum5 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][0][i])
        sum5 += val
    }
    sums[4] = sum5

    // From back to front
    sum6 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][0][cols-1-i])
        sum6 += val
    }
    sums[5] = sum6

    // Face Diagonal Down
    // From front to back
    sum7 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][rows-1][0+i])
        sum7 += val
    }
    sums[6] = sum7

    // From back to front
    sum8 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][rows-1][cols-1-i])
        sum8 += val
    }
    sums[7] = sum8

    // Face Diagonal Left
    // From front to back
    sum9 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][0])
        sum9 += val
    }
    sums[8] = sum9

    // From back to front
    sum10 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][rows-1-i][0])
        sum10 += val
    }
    sums[9] = sum10

    // Face Diagonal Right
    // From front to back
    sum11 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][cols-1])
        sum11 += val
    }
    sums[10] = sum11

    // From back to front
    sum12 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][rows-1-i][cols-1])
        sum12 += val
    }
    sums[11] = sum12

    return sums
}

// SumSpaceDiagonal calculates the sum of the space diagonals
func SumSpaceDiagonal(tables [][][]string) []int {
    sums := make([]int, 4)

    // Space Diagonal 1
    sum1 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][i]) // row (i) col (i)
        sum1 += val
    }
    sums[0] = sum1

    // Space Diagonal 2
    sum2 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][cols-1-i]) // row (i) col (cols-1-i)
        sum2 += val
    }
    sums[1] = sum2

    // Space Diagonal 3
    sum3 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][rows-1-i][i]) // row (rows-1-i) col (i)
        sum3 += val
    }
    sums[2] = sum3

    // Space Diagonal 4
    sum4 := 0
    for i := 0; i < numTables; i++ {
        val, _ := strconv.Atoi(tables[i][rows-1-i][cols-1-i]) // row (rows-1-i) col (cols-1)
        sum4 += val
    }
    sums[3] = sum4

    return sums
}