package main

import (
	"bufio"
	"os"
)

func parseInputFile() *[][]string {
	file, err := os.Open("25/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	patterns := make([][]string, 0)
	currentPattern := make([]string, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(currentPattern) > 0 {
				patterns = append(patterns, currentPattern)
				currentPattern = make([]string, 0)
			}
			continue
		}
		currentPattern = append(currentPattern, line)
	}
	if len(currentPattern) > 0 {
		patterns = append(patterns, currentPattern)
	}

	return &patterns
}

func canCombine(pat1, pat2 []string) bool {
	for y := 0; y < len(pat1); y++ {
		for x := 0; x < len(pat1[y]); x++ {
			if pat1[y][x] == '#' && pat2[y][x] == '#' {
				return false
			}
		}
	}
	return true
}

func tryCombinations(patterns *[][]string) int {
	count := 0
	for i := 0; i < len(*patterns); i++ {
		for j := i + 1; j < len(*patterns); j++ {
			if canCombine((*patterns)[i], (*patterns)[j]) {
				count++
			}
		}
	}
	return count
}

func main() {
	patterns := parseInputFile()

	println(tryCombinations(patterns))
}
