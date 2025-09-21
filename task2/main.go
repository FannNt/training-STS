package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"unicode"
	"unicode/utf8"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <filename>")
	}

	filename := os.Args[1]
	
	// Read the file
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	// Count alphabet characters, numbers, and unreadable characters
	charCount := make(map[rune]int)
	numberCount := make(map[rune]int)
	unreadableCount := 0
	totalProcessed := 0
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()
		
		// Process each byte in the line to handle invalid UTF-8
		for len(line) > 0 {
			r, size := utf8.DecodeRuneInString(line)
			totalProcessed++
			
			if r == utf8.RuneError && size == 1 {
				// Invalid UTF-8 sequence
				unreadableCount++
				fmt.Printf("Warning: Found unreadable character at position %d, continuing...\n", totalProcessed)
			} else if unicode.IsLetter(r) {
				// Valid letter character
				charCount[unicode.ToLower(r)]++
			} else if unicode.IsDigit(r) {
				// Valid number character
				numberCount[r]++
			}
			// Skip other valid characters (spaces, punctuation, etc.)
			
			line = line[size:]
		}
	}

	// Handle scanner errors but continue processing
	if err := scanner.Err(); err != nil {
		fmt.Printf("Warning: Error reading file: %v, but continuing with processed data...\n", err)
	}

	// Sort characters alphabetically for consistent output
	var chars []rune
	for char := range charCount {
		chars = append(chars, char)
	}
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})

	// Sort numbers numerically for consistent output
	var numbers []rune
	for num := range numberCount {
		numbers = append(numbers, num)
	}
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})

	// Print results
	fmt.Println("\n=== Character Count Results ===")
	if len(chars) > 0 {
		for _, char := range chars {
			fmt.Printf("%c = %d\n", char, charCount[char])
		}
	} else {
		fmt.Println("No alphabet characters found in the file.")
	}

	fmt.Println("\n=== Number Count Results ===")
	if len(numbers) > 0 {
		for _, num := range numbers {
			fmt.Printf("%c = %d\n", num, numberCount[num])
		}
	} else {
		fmt.Println("No numbers found in the file.")
	}
	
	// Print summary
	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total alphabet characters: %d\n", getTotalCount(charCount))
	fmt.Printf("Total numbers: %d\n", getTotalCount(numberCount))
	fmt.Printf("Unreadable characters: %d\n", unreadableCount)
	fmt.Printf("Total characters processed: %d\n", totalProcessed)
}

func getTotalCount(charCount map[rune]int) int {
	total := 0
	for _, count := range charCount {
		total += count
	}
	return total
}