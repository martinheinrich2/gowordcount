package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// CountLines counts the number of lines in the file
func CountLines(filename string) int {
	if filepath.Ext(filename) != ".txt" {
		fmt.Println("File must have a .txt extension, but has", filepath.Ext(filename))
		return 0
	}
	fmt.Println("Counting lines in file:", filename)
	// Open file
	file, err := os.Open(filename)
	// Check for error opening file
	if err != nil {
		fmt.Println(err)
	}
	// Close file at end of program
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing file", err)
		}
	}(file)

	// Create a new scanner
	scanner := bufio.NewScanner(file)
	// Set the split function(ScanLines) for the scanner and count the lines
	scanner.Split(bufio.ScanLines)
	count := 0
	for scanner.Scan() {
		count += 1
	}
	// Check for error during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scanning file", err)
	}
	return count
}
