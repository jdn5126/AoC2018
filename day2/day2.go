package main

import (
	"bufio"
	"fmt"
	"os"
)

// Return the intersection of two strings
func intersection(x string, y string) string {
	var intersection string
	for i, char := range x {
		if x[i] == y[i] {
			intersection += string(char)
		}
	}
	return intersection
}

// Return whether string has 2 and 3 instances of unique characters
func charCheck(x string) (bool, bool) {
	hasTwo, hasThree := false, false
	charSet := make(map[rune]int)
	// Get count for each char
	for _, char := range x {
		if _, ok := charSet[char]; ok {
			charSet[char] += 1
		} else {
			charSet[char] = 1
		}
	}
	// Determine if word has 2 and 3 instance characters
	for _, val := range charSet {
		if val == 2 {
			hasTwo = true
		} else if val == 3 {
			hasThree = true
		}
	}
	return hasTwo, hasThree
}

func main() {
	// Input file is passed as command line arg
	if len(os.Args) != 2 {
		fmt.Println("Usage: day2 <file_path>")
		return
	}

	// Read input file line by line using bufio
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to open", os.Args[1])
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Part 1: Count words with two and three letter repeats
	// Part 2: Find strings with Hamming distance of one
	var word string
	wordSet := make(map[string]bool)
	var twoCount int = 0
	var threeCount int = 0
	var output string = ""

	for scanner.Scan() {
		word = scanner.Text()
		// Determine if word meets character count conditions
		hasTwo, hasThree := charCheck(word)
		if hasTwo {
			twoCount += 1
		}
		if hasThree {
			threeCount += 1
		}

		// Determine if word has Hamming distance of one from other words
		for key, _ := range wordSet {
			intersection := intersection(word, key)
			if (len(word) - 1) == len(intersection) {
				output += intersection
			}
		}
		// Add word to word set
		wordSet[word] = true
	}
	fmt.Println("Checksum:", twoCount*threeCount)
	fmt.Println("Common letters:", output)
}
