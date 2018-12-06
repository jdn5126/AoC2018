package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"unicode"
)

// Return length of resulting polymer when ignore is excluded, where ignore is
// always passed as uppercase version of unicode.
func polymerLength(polymer []byte, ignore rune) int {
	var x int = 0
	// Find first index not equal to rune that we are ignoring
	for unicode.ToUpper(rune(polymer[x])) == ignore {
		x += 1
	}
	// Make sure we have not exceeded length of polymer
	if x >= len(polymer) {
		return 0
	}
	// Store reduced polymer
	var output []byte
	output = append(output, polymer[x])
	outputLen := 1
	for i := x + 1; i < len(polymer); i++ {
		currRune := rune(polymer[i])
		// Continue if this element is to be ignored
		if unicode.ToUpper(currRune) == ignore {
			continue
		}
		// Output can empty at any time
		var lastAdded rune
		if outputLen > 0 {
			lastAdded = rune(output[outputLen-1])
		} else {
			output = append(output, polymer[i])
			outputLen = 1
			continue
		}
		// Only remove if polarity is opposite
		if unicode.ToUpper(currRune) == unicode.ToUpper(lastAdded) && currRune != lastAdded {
			// Remove last value in output
			if outputLen == 1 {
				output = nil
			} else {
				output = output[:outputLen-1]
			}
			outputLen -= 1
		} else {
			output = append(output, polymer[i])
			outputLen += 1
		}
	}
	return outputLen
}

func main() {
	// Input file is passed as command line arg
	if len(os.Args) != 2 {
		fmt.Println("Usage: day5 <file_path>")
		return
	}

	// Read entire file into string using ioutil
	content, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Unable to open", os.Args[1])
	}
	// Remove newline character
	content = content[:len(content)-1]

	// Part1: Reduce array using delete-safe iteration (tradeoff storage for speed)
	fmt.Println("Part1:", polymerLength(content, '\x00'))

	// Part2: Find length of shortest polymer when ignoring single pair
	// Input is constrained to letters of the English alphabet, so build static rune alphabet
	// Need to be fast to beat Doug.
	var alphabet [26]rune
	for i := 0; i < 26; i++ {
		alphabet[i] = rune(65 + i)
	}

	shortestLength := int(^uint(0) >> 1)
	for _, val := range alphabet {
		length := polymerLength(content, val)
		if length < shortestLength {
			shortestLength = length
		}
	}
	fmt.Println("Part2:", shortestLength)
}
