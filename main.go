package main

import (
	geneticalgorithm "Tubes1_AI/algorithms/genetic-algorithm"
	simulatedannealing "Tubes1_AI/algorithms/simulated-annealing"
	steepestascent "Tubes1_AI/algorithms/steepest-ascent"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Select an algorithm to run:")
	fmt.Println("1. Steepest Ascent")
	fmt.Println("2. Simulated Annealing")
	fmt.Println("3. Genetic Algorithm")

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter choice (1, 2, or 3): ")
	input, _ := reader.ReadString('\n')
	choice, err := strconv.Atoi(input[:len(input)-1])

	if err != nil {
		fmt.Println("Invalid input. Please enter a number (1, 2, or 3).")
		return
	}

	switch choice {
	case 1:
		fmt.Println("Running Steepest Ascent Algorithm...")
		steepestascent.Run()
	case 2:
		fmt.Println("Running Simulated Annealing Algorithm...")
		simulatedannealing.Run()
	case 3:
		fmt.Println("Running Genetic Algorithm...")
		geneticalgorithm.Run()
	default:
		fmt.Println("Invalid choice. Please select 1, 2, or 3.")
	}
}
