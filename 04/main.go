package main

import (
	"bufio"
	"os"
)

var input string

func parseInputFile() {
	file, err := os.Open("04/input.txt")
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

func main() {
	parseInputFile()

	println(input)
}
