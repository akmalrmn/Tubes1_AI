package models

import (
    "fmt"
    "math/rand"
    "strconv"
    "time"
)

const (
    Rows      = 5
    Cols      = 5
    NumTables = 5
)

type Cube struct {
    Tables [][][]string
}

func GenerateTable() [][][]string {
    rand.Seed(time.Now().UnixNano())

    numbers := make([]int, Rows*Cols*NumTables)
    for i := 0; i < Rows*Cols*NumTables; i++ {
        numbers[i] = i + 1
    }

    rand.Shuffle(len(numbers), func(i, j int) {
        numbers[i], numbers[j] = numbers[j], numbers[i]
    })

    fmt.Println("Generated numbers:", numbers)

    tables := make([][][]string, NumTables)
    counter := 0
    for t := 0; t < NumTables; t++ {
        table := make([][]string, Rows)
        for i := range table {
            table[i] = make([]string, Cols)
            for j := range table[i] {
                table[i][j] = fmt.Sprintf("%d", numbers[counter])
                counter++
            }
        }
        tables[t] = table
    }

    return tables
}

func GenerateCube() Cube {
    return Cube{Tables: GenerateTable()}
}

func SumColumns(tables [][][]string) []int {
    sums := make([]int, Cols*NumTables)
    for tableIdx := 0; tableIdx < NumTables; tableIdx++ {
        for col := 0; col < Cols; col++ {
            sum := 0
            for row := 0; row < Rows; row++ {
                val, _ := strconv.Atoi(tables[tableIdx][row][col])
                sum += val
            }
            sums[tableIdx*Cols+col] = sum
        }
    }
    return sums
}

func SumRows(tables [][][]string) []int {
    sums := make([]int, Rows*NumTables)
    for tableIdx := 0; tableIdx < NumTables; tableIdx++ {
        for row := 0; row < Rows; row++ {
            sum := 0
            for col := 0; col < Cols; col++ {
                val, _ := strconv.Atoi(tables[tableIdx][row][col])
                sum += val
            }
            sums[tableIdx*Rows+row] = sum
        }
    }
    return sums
}

func SumPoles(tables [][][]string) []int {
    sums := make([]int, Rows*Cols)
    for pole := 0; pole < Rows*Cols; pole++ {
        sum := 0
        for tableIdx := 0; tableIdx < NumTables; tableIdx++ {
            rowIdx := pole / Cols
            colIdx := pole % Cols
            val, _ := strconv.Atoi(tables[tableIdx][rowIdx][colIdx])
            sum += val
        }
        sums[pole] = sum
    }
    return sums
}

func EvaluateIndividual(cube Cube) float64 {
    columnSums := SumColumns(cube.Tables)
    rowSums := SumRows(cube.Tables)
    poleSums := SumPoles(cube.Tables)
    faceDiagonalSums := SumFaceDiagonal(cube.Tables)
    spaceDiagonalSums := SumSpaceDiagonal(cube.Tables)

    targetSum :=  315
    fitness := 0.0

    for _, sum := range columnSums {
        fitness -= float64(abs(sum - targetSum))
    }
    for _, sum := range rowSums {
        fitness -= float64(abs(sum - targetSum))
    }
    for _, sum := range poleSums {
        fitness -= float64(abs(sum - targetSum))
    }
    for _, sum := range faceDiagonalSums {
        fitness -= float64(abs(sum - targetSum))
    }
    for _, sum := range spaceDiagonalSums {
        fitness -= float64(abs(sum - targetSum))
    }

    return -fitness
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func SumFaceDiagonal(tables [][][]string) []int {
    sums := make([]int, 12)

    sum1 := 0
    for i := 0; i < Rows; i++ {
        val, _ := strconv.Atoi(tables[0][i][i])
        sum1 += val
    }
    sums[0] = sum1

    sum2 := 0
    for i := 0; i < Rows; i++ {
        val, _ := strconv.Atoi(tables[0][Rows-1-i][i])
        sum2 += val
    }
    sums[1] = sum2

    sum3 := 0
    for i := 0; i < Rows; i++ {
        val, _ := strconv.Atoi(tables[NumTables-1][i][i])
        sum3 += val
    }
    sums[2] = sum3

    sum4 := 0
    for i := 0; i < Rows; i++ {
        val, _ := strconv.Atoi(tables[NumTables-1][Rows-1-i][i])
        sum4 += val
    }
    sums[3] = sum4

    sum5 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][0][i])
        sum5 += val
    }
    sums[4] = sum5

    sum6 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][0][Cols-1-i])
        sum6 += val
    }
    sums[5] = sum6

    sum7 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][Rows-1][0+i])
        sum7 += val
    }
    sums[6] = sum7

    sum8 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][Rows-1][Cols-1-i])
        sum8 += val
    }
    sums[7] = sum8

    sum9 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][0])
        sum9 += val
    }
    sums[8] = sum9

    sum10 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][Rows-1-i][0])
        sum10 += val
    }
    sums[9] = sum10

    sum11 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][Cols-1])
        sum11 += val
    }
    sums[10] = sum11

    sum12 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][Rows-1-i][Cols-1])
        sum12 += val
    }
    sums[11] = sum12

    return sums
}

func SumSpaceDiagonal(tables [][][]string) []int {
    sums := make([]int, 4)

    sum1 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][i])
        sum1 += val
    }
    sums[0] = sum1

    sum2 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][i][Cols-1-i])
        sum2 += val
    }
    sums[1] = sum2

    sum3 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][Rows-1-i][i])
        sum3 += val
    }
    sums[2] = sum3

    sum4 := 0
    for i := 0; i < NumTables; i++ {
        val, _ := strconv.Atoi(tables[i][Rows-1-i][Cols-1-i])
        sum4 += val
    }
    sums[3] = sum4

    return sums
}