package main

import (
	"bufio"
	"os"
)

func parseInputFile() (*[][]rune, *string) {
	file, err := os.Open("15/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data := make([][]rune, 0)
	actions := ""
	parseMap := true
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			parseMap = false
			continue
		}

		if parseMap {
			data = append(data, []rune(line))
		} else {
			actions += line
		}
	}

	return &data, &actions
}

func getRobotPos(data *[][]rune) (int, int) {
	for indRow, row := range *data {
		for indCol, col := range row {
			if col == '@' {
				return indRow, indCol
			}
		}
	}
	return -1, -1
}

func move(data *[][]rune, row, col, dirRow, dirCol int) (int, int) {
	if (*data)[row+dirRow][col+dirCol] == '#' {
		return row, col
	}

	if (*data)[row+dirRow][col+dirCol] == '.' {
		(*data)[row][col] = '.'
		(*data)[row+dirRow][col+dirCol] = '@'
		return row + dirRow, col + dirCol
	}

	if (*data)[row+dirRow][col+dirCol] != 'O' {
		return row, col
	}

	newRow, newCol := row+dirRow, col+dirCol

	for {
		newRow, newCol = newRow+dirRow, newCol+dirCol
		if (*data)[newRow][newCol] == '#' {
			return row, col
		}

		if (*data)[newRow][newCol] == '.' {
			break
		}
	}

	(*data)[newRow][newCol] = 'O'
	(*data)[row][col] = '.'
	(*data)[row+dirRow][col+dirCol] = '@'

	return row + dirRow, col + dirCol
}

func executeActions(data *[][]rune, actions *string) {
	robotX, robotY := getRobotPos(data)

	for _, action := range *actions {
		switch action {
		case '>':
			robotX, robotY = move(data, robotX, robotY, 0, 1)
		case '<':
			robotX, robotY = move(data, robotX, robotY, 0, -1)
		case '^':
			robotX, robotY = move(data, robotX, robotY, -1, 0)
		case 'v':
			robotX, robotY = move(data, robotX, robotY, 1, 0)
		}
	}
}

func printMap(data *[][]rune) {
	for _, row := range *data {
		println(string(row))
	}
}

func calcSumBoxes(data *[][]rune) int {
	sum := 0
	for indRow, row := range *data {
		for indCol, col := range row {
			if col == 'O' {
				sum += indRow*100 + indCol
			}
		}
	}
	return sum
}

func main() {
	data, actions := parseInputFile()

	executeActions(data, actions)
	printMap(data)
	println("Part 1:", calcSumBoxes(data))
}
