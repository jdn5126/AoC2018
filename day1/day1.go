package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Input file is passed as command line arg
	if len(os.Args) != 2 {
		fmt.Println("Usage: day1 <file_path>")
		return
	}

	// Read input file line by line using bufio
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to open", os.Args[1])
	}
	scanner := bufio.NewScanner(file)

	// Part 1: Calculate final frequency
	// Part 2: Print first repeated frequency
	var frequency int = 0
	var input int
	set := make(map[int]bool)
	set[0] = true
	var array []int
	var foundRepeat bool = false

	// Calculate frequency and store input values in array
	for scanner.Scan() {
		input, _ = strconv.Atoi(scanner.Text())
		array = append(array, input)
		frequency += input
		if _, ok := set[frequency]; ok && !foundRepeat {
			fmt.Println("Repeated frequency is:", frequency)
			foundRepeat = true
		}
		set[frequency] = true
	}
	file.Close()
	fmt.Println("Ending frequency is:", frequency)

	// If repeat value has not been found yet, restart iteration
Forever:
	for {
		for _, num := range array {
			frequency += num
			if _, ok := set[frequency]; ok {
				fmt.Println("Repeated frequency is:", frequency)
				break Forever
			}
			set[frequency] = true
		}
	}
}
