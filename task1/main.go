package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Check command line arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		fmt.Println("Example: go run main.go input.txt output.txt")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Read numbers from input file
	numbers, err := readNumbersFromFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading input file: %v", err)
	}

	// Sort numbers in ascending order
	sort.Ints(numbers)

	// Write sorted numbers to output file
	err = writeNumbersToFile(outputFile, numbers)
	if err != nil {
		log.Fatalf("Error writing output file: %v", err)
	}

	fmt.Printf("Successfully sorted %d numbers from %s to %s\n", len(numbers), inputFile, outputFile)
}

// readNumbersFromFile reads integers from a file, one per line
func readNumbersFromFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines
		if line == "" {
			continue
		}

		// Convert string to integer
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, fmt.Errorf("invalid number '%s' on line %d: %w", line, lineNumber, err)
		}

		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	return numbers, nil
}

// writeNumbersToFile writes integers to a file, one per line
func writeNumbersToFile(filename string, numbers []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filename, err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, num := range numbers {
		_, err := fmt.Fprintf(writer, "%d\n", num)
		if err != nil {
			return fmt.Errorf("failed to write number %d: %w", num, err)
		}
	}

	return nil
}
