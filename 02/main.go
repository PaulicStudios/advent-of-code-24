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

func checkRow(row []int) bool {
	dir := UNDEFINED
	for indRow, col := range row {
		if indRow == 0 {
			continue
		}

		diff := col - row[indRow-1]
		if diff == 0 {
			return false
		}
		if dir == UNDEFINED {
			if diff > 0 {
				dir = INCREASE
			} else {
				dir = DECREASE
			}
		}
		if dir == INCREASE && (diff < 0 || diff > 3) {
			return false
		}
		if dir == DECREASE && (diff > 0 || diff < -3) {
			return false
		}
	}
	return true
}

func numSaveReports(removeOneLevel bool) int {
	numValidReports := 0
	for _, row := range list {
		if checkRow(row) {
			numValidReports++
			continue
		}
		if removeOneLevel {
			for ind := range row {
				rowCopy := make([]int, len(row))
				copy(rowCopy, row)
				rowCopy = append(rowCopy[:ind], rowCopy[ind+1:]...)
				if checkRow(rowCopy) {
					numValidReports++
					break
				}
			}
		}
	}
	return numValidReports
}

func main() {
	parseInputFile()

	println("Part 1: ", numSaveReports(false))
	println("Part 2: ", numSaveReports(true))
}
