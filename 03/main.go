package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
	"regexp"
)

var input string

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
		results[ind][0] = utils.ConvertToInt(match[1])
		results[ind][1] = utils.ConvertToInt(match[2])
	}
	return results
}

func findOccurrencesEnable(input string) [][2]int {
	re := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\))`)
	matches := re.FindAllStringSubmatch(input, -1)

	var results [][2]int
	results = make([][2]int, 0)
	do := true

	for _, match := range matches {
		if match[0] == "do()" {
			do = true
			continue
		} else if match[0] == "don't()" {
			do = false
			continue
		}
		if !do {
			continue
		}
		results = append(results, [2]int{utils.ConvertToInt(match[2]), utils.ConvertToInt(match[3])})
	}
	return results
}

type findOccurrencesFunc func(string) [][2]int

func multiplySum(occurrencesFunc findOccurrencesFunc) int {
	occurrences := occurrencesFunc(input)
	sum := 0
	for _, occurrence := range occurrences {
		sum += occurrence[0] * occurrence[1]
	}
	return sum
}

func main() {
	parseInputFile()

	println("Part 1: ", multiplySum(findOccurrences))
	println("Part 2: ", multiplySum(findOccurrencesEnable))
}
