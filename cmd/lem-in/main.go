package main

import (
	"fmt"
	"os"

	"mimo/internal/algorithm"
	"mimo/internal/output"
	"mimo/internal/parser"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ./lem-in input.txt")
		os.Exit(1)
	}

	filePath := os.Args[1]
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filePath, err)
		os.Exit(1)
	}
	defer file.Close()

	// 1. Parse the input
	farm, err := parser.Parse(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing input:", err)
		os.Exit(1)
	}

	// 2. Find all paths
	allPaths := algo.FindAllPaths(farm)
	if len(allPaths) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No paths found from start to end")
		os.Exit(1)
	}

	// 3. Find the best group of paths
	bestGroup := algo.FindBestGroup(farm.AntCount, allPaths)
	if len(bestGroup) == 0 {
		fmt.Fprintln(os.Stderr, "Error: No valid group of paths found")
		os.Exit(1)
	}

	// 4. Output original input (mandatory for lem-in project)
	for _, line := range farm.Input {
		fmt.Println(line)
	}
	fmt.Println()

	// 5. Simulate ants moving
	output.SimulateAntsSmart(farm, bestGroup)
}
