package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func openFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}


func main() {
	file, err := openFile("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	forward, aim, depth := 0, 0, 0
	direction := "none"
	distance := 0

	// This entire loop is incredibly poorly optimized, but it works.
	//
	// If it works, it ain't stupid
	// - Murphy... Probably
	//
	for _, s := range file {
		a := strings.Split(s, " ")

		for i, b := range a {
			if i == 0 {
				direction = b
			} else {
				distance, _ = strconv.Atoi(b) // Convert the second value to an int
				switch direction {
				case "forward":
					if aim == 0 {
						forward += distance
					} else if aim > 0 {
						forward += distance
						depth =  depth + (distance * aim)
					}
				case "up":
					aim -= distance
				case "down":
					aim += distance
				}
			}
		}
	fmt.Printf("\n(Input: %s %v) Current coordinates: Forward=%v, Depth=%v, Aim=%v", direction, distance, forward, depth, aim)

	}

	finalCoordinates := forward * depth
	fmt.Printf("\n\nFinal sum: %v", finalCoordinates)

}
