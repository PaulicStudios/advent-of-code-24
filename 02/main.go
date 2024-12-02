package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var list [][]int

func convertToInt(s string) int {
	nbr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nbr
}

func parseInputFile() {
	file, err := os.Open("02/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		list = append(list, make([]int, len(split)))

		for ind, s := range split {
			nbr := convertToInt(s)
			list[lineIndex][ind] = nbr
		}
		lineIndex++
	}
}

type Direction int

const (
	UNDEFINED Direction = iota
	INCREASE
	DECREASE
)

func numSaveReports() int {
	numValidReports := 0
	for _, col := range list {
		var dir = UNDEFINED
		for indRow, row := range col {
			if indRow == 0 {
				continue
			}

			diff := row - col[indRow-1]
			if diff == 0 {
				break
			}
			if dir == UNDEFINED {
				if diff > 0 {
					dir = INCREASE
				} else {
					dir = DECREASE
				}
			}
			if dir == INCREASE && (diff < 0 || diff > 3) {
				break
			}
			if dir == DECREASE && (diff > 0 || diff < -3) {
				break
			}
			if indRow == len(col)-1 {
				numValidReports++
			}
		}
	}
	return numValidReports
}

func main() {
	parseInputFile()

	println("Part 1: ", numSaveReports())
}
