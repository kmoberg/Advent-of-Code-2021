package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// readInts takes in a STRING and converts it to a slice of ints. What the hell Golang.
func readInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	// Initialize a slice of integers
	var result []int

	// Use scanner to go loop over the input string
	for scanner.Scan() {
		// Stringconv.Atoi converts the strings to ints.
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}

		// Append the result of this loop to the slice
		result = append(result, x)
	}

	// Return me baby!
	return result, scanner.Err()
}

// openFile allows for the read of a file and save it as a slice of strings
func openFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// Don't forget to close the file
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// *shocked Pikachu* We need a slice of strings to insert the file into!
	// The dumb part? bufio.NewScanner can't fucking scan it into integers,
	// it can only use strings, booleans and a few other things, but not numbers-
	var lines []string

	// Why Golang. Why.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Return me. I guess.
	return lines, scanner.Err()
}

func main() {
	// Start by opening the input file
	file, err := openFile("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	// Now comes the stupid shit. Convert the slice of strings to a one-line string.
	superString := strings.Join(file, " ")

	// Take that superString and read it back as a slice of ints. Yes, it's dumb.
	lines, err := readInts(strings.NewReader(superString))

	fmt.Println(lines)  // Useful for debugging

	// Initialize a bunch of variables
	prevNum, numIncrease, numDecrease := 0, 0, 0

	// Now for some simple checks, start by looping through the slice of ints
	for i,line := range lines {

		// The first number obviously can't be larger than the previous, so do a dirty
		// hack just ignore it
		if i > 0 {

			// Is the current number higher than the previous?
			if line > prevNum {
				numIncrease++
				fmt.Printf("[%v] %v (increase) \n", i, line)
			} else if line < prevNum {
				numDecrease++
				fmt.Printf("[%v] %v (decrease) \n", i, line)
			} else {
				fmt.Printf("[%v] %v (no change) \n", i, line)
			}
		} else { // Handle the first number and just mark it as no change
			fmt.Printf("[%v] %v (no change) \n", i, line)
		}
		// Set the previous number to this loop, then start all over
		prevNum = line

	}

	fmt.Println("\nFinal tally:")
	fmt.Printf("No. Increases: %v \n", numIncrease)
	fmt.Printf("No. Decreases: %v", numDecrease)

}
