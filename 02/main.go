package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

var input string

func convertToInt(s string) int {
	nbr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nbr
}

func parseInputFile() {
	file, err := os.Open("03/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	input = ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		input += line + "\n"
	}
}

func findOccurrences(input string) [][2]int {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := re.FindAllStringSubmatch(input, -1)

	var results [][2]int
	results = make([][2]int, 0)

	for ind, match := range matches {
		results = append(results, [2]int{})
		results[ind][0] = convertToInt(match[1])
		results[ind][1] = convertToInt(match[2])
	}
	return results
}

func multiplySum() int {
	occurrences := findOccurrences(input)
	sum := 0
	for _, occurrence := range occurrences {
		sum += occurrence[0] * occurrence[1]
	}
	return sum
}

func main() {
	parseInputFile()

	println("Part 1: ", multiplySum())
}
