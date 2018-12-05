package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Go does not provide a function to find the max or min of two integers...
func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

type Square struct {
	Id          int
	TopLeft     Coordinate
	BottomRight Coordinate
}

type Coordinate struct {
	X int
	Y int
}

// Check if overlap exists and update overlapMap if it does
func overlap(s1 Square, s2 Square, overlapMap map[Coordinate]bool) bool {
	// Given corners of each square, find overlapping square
	var topLeft Coordinate = Coordinate{
		max(s1.TopLeft.X, s2.TopLeft.X),
		max(s1.TopLeft.Y, s2.TopLeft.Y),
	}
	var bottomRight Coordinate = Coordinate{
		min(s1.BottomRight.X, s2.BottomRight.X),
		min(s1.BottomRight.Y, s2.BottomRight.Y),
	}
	// Check if overlap exists
	if topLeft.X > bottomRight.X || topLeft.Y > bottomRight.Y {
		return false
	}
	// Record overlapping square
	for i := topLeft.X; i <= bottomRight.X; i++ {
		for j := topLeft.Y; j <= bottomRight.Y; j++ {
			overlapMap[Coordinate{i, j}] = true
		}
	}
	return true
}

func main() {
	// Input file is passed as command line arg
	if len(os.Args) != 2 {
		fmt.Println("Usage: day3 <file_path>")
		return
	}

	// Read input file line by line using bufio
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to open", os.Args[1])
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Part 1: Find overlapping squares
	// Part 2: Find id of ONLY square that does not overlap
	var squares []Square
	// Track which coordinates overlap (use map for fast lookup)
	overlapMap := make(map[Coordinate]bool)
	// Track which squares do not overlap with any other squares
	safeMap := make(map[Square]bool)
	var overlaps bool

	for scanner.Scan() {
		// Split string and create square from string
		line := strings.Split(scanner.Text(), " @ ")
		num := strings.Split(line[0], "#")[1]
		lineSplit := strings.Split(line[1], ":")
		var coordinateString, sizeString string = lineSplit[0], lineSplit[1]
		coordinateSplit := strings.Split(coordinateString, ",")
		sizeSplit := strings.Split(sizeString, "x")
		id, _ := strconv.Atoi(num)
		x, _ := strconv.Atoi(coordinateSplit[1])
		y, _ := strconv.Atoi(coordinateSplit[0])
		width, _ := strconv.Atoi(strings.TrimSpace(sizeSplit[0]))
		height, _ := strconv.Atoi(sizeSplit[1])
		var square Square = Square{
			Id:          id,
			TopLeft:     Coordinate{x, y},
			BottomRight: Coordinate{x + height - 1, y + width - 1},
		}
		// Check if square overlaps with any other squares
		overlaps = false
		for _, s := range squares {
			if overlap(s, square, overlapMap) {
				overlaps = true
				delete(safeMap, s)
			}
		}
		if !overlaps {
			safeMap[square] = true
		}
		// Append to array
		squares = append(squares, square)
	}
	// Print inches of fabric claimed by two or more elves
	fmt.Println(len(overlapMap))
	// Print map of Squares that do not overlap with any other Squares
	for square, _ := range safeMap {
		fmt.Println(square.Id)
	}
}
