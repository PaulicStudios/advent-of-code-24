package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
)

func parseInputFile() [][]int {
	file, err := os.Open("10/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data := make([][]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]int, 0)
		for _, char := range line {
			if char == '.' {
				row = append(row, -1)
				continue
			}
			row = append(row, utils.ConvertToInt(string(char)))
		}
		data = append(data, row)
	}

	return data
}

var part1 = true

func checkPos(row, col, curHeight int, data *[][]int, visited *[][2]int) bool {
	if row < 0 || row >= len(*data) {
		return false
	}
	if col < 0 || col >= len((*data)[0]) {
		return false
	}
	if (*data)[row][col] != curHeight+1 {
		return false
	}
	if part1 {
		for _, pos := range *visited {
			if pos[0] == row && pos[1] == col {
				return false
			}
		}
	}
	return true
}

var validRoutes = 0

func step(curRow, curCol int, data *[][]int, visited *[][2]int) {
	curHeight := (*data)[curRow][curCol]

	if curHeight == 9 {
		validRoutes++
		return
	}

	if checkPos(curRow-1, curCol, curHeight, data, visited) {
		*visited = append(*visited, [2]int{curRow - 1, curCol})
		step(curRow-1, curCol, data, visited)
	}
	if checkPos(curRow+1, curCol, curHeight, data, visited) {
		*visited = append(*visited, [2]int{curRow + 1, curCol})
		step(curRow+1, curCol, data, visited)
	}
	if checkPos(curRow, curCol-1, curHeight, data, visited) {
		*visited = append(*visited, [2]int{curRow, curCol - 1})
		step(curRow, curCol-1, data, visited)
	}
	if checkPos(curRow, curCol+1, curHeight, data, visited) {
		*visited = append(*visited, [2]int{curRow, curCol + 1})
		step(curRow, curCol+1, data, visited)
	}
}

func findStarts(data *[][]int) {
	for row := range *data {
		for col := range (*data)[row] {
			if (*data)[row][col] == 0 {
				visited := [][2]int{{row, col}}
				step(row, col, data, &visited)
			}
		}
	}
}

func main() {
	data := parseInputFile()
	findStarts(&data)
	println("Part 1: ", validRoutes)
	part1 = false
	validRoutes = 0
	findStarts(&data)
	println("Part 2: ", validRoutes)
}
