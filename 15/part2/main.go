package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseInputFile() (*[][]rune, *string) {
	file, err := os.Open("15/part2/input.txt")
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

		newLine := ""
		for _, char := range line {
			switch char {
			case '#':
				newLine += "##"
			case 'O':
				newLine += "[]"
			case '.':
				newLine += ".."
			case '@':
				newLine += "@."
			}
		}

		if parseMap {
			data = append(data, []rune(newLine))
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

func checkMoveVertical(data *[][]rune, row, col, dirRow int) bool {
	if (*data)[row][col] == '#' {
		return false
	}

	if (*data)[row][col] == '.' {
		return true
	}

	free := checkMoveVertical(data, row+dirRow, col, dirRow)
	if !free {
		return false
	}

	if (*data)[row][col] == '[' {
		free = checkMoveVertical(data, row+dirRow, col+1, dirRow)
	} else if (*data)[row][col] == ']' {
		free = checkMoveVertical(data, row+dirRow, col-1, dirRow)
	} else {
		panic("Invalid character: " + string((*data)[row][col]))
	}

	return free
}

func moveVertical(data *[][]rune, row, col, dirRow int) {
	if (*data)[row][col] == '.' {
		return
	}

	moveVertical(data, row+dirRow, col, dirRow)

	if (*data)[row][col] == '[' {
		moveVertical(data, row+dirRow, col+1, dirRow)
	} else if (*data)[row][col] == ']' {
		moveVertical(data, row+dirRow, col-1, dirRow)
	} else {
		panic("Invalid character: " + string((*data)[row][col]))
	}

	if (*data)[row][col] == '[' {
		(*data)[row][col+1] = '.'
		(*data)[row+dirRow][col] = '['
		(*data)[row+dirRow][col+1] = ']'
	} else if (*data)[row][col] == ']' {
		(*data)[row][col-1] = '.'
		(*data)[row+dirRow][col] = ']'
		(*data)[row+dirRow][col-1] = '['
	}
	(*data)[row][col] = '.'
}

func moveHorizontal(data *[][]rune, row, col, dirCol int) bool {
	if (*data)[row][col] == '#' {
		return false
	}

	if (*data)[row][col] == '.' {
		return true
	}

	free := moveHorizontal(data, row, col+dirCol, dirCol)
	if !free {
		return false
	}

	(*data)[row][col+dirCol] = (*data)[row][col]
	(*data)[row][col] = '.'

	return true
}

func executeActions(data *[][]rune, actions *string) {
	robotRow, robotCol := getRobotPos(data)

	for _, action := range *actions {
		free := true
		dirRow, dirCol := 0, 0
		switch action {
		case '>':
			dirRow, dirCol = 0, 1
			free = moveHorizontal(data, robotRow, robotCol+dirCol, dirCol)
		case '<':
			dirRow, dirCol = 0, -1
			free = moveHorizontal(data, robotRow, robotCol+dirCol, dirCol)
		case '^':
			dirRow, dirCol = -1, 0
			free = checkMoveVertical(data, robotRow+dirRow, robotCol, dirRow)
			if free {
				moveVertical(data, robotRow+dirRow, robotCol, dirRow)
			}
		case 'v':
			dirRow, dirCol = 1, 0
			free = checkMoveVertical(data, robotRow+dirRow, robotCol, dirRow)
			if free {
				moveVertical(data, robotRow+dirRow, robotCol, dirRow)
			}
		}

		if free {
			(*data)[robotRow][robotCol] = '.'
			robotRow, robotCol = robotRow+dirRow, robotCol+dirCol
			(*data)[robotRow][robotCol] = '@'
		}
		// printMap(data)
		// bufio.NewReader(os.Stdin).ReadString('\n')
	}
}

func printMap(data *[][]rune) {
	// print("\033[H\033[2J")
	for _, row := range *data {
		fmt.Println(string(row))
	}
	fmt.Println()
	// time.Sleep(100 * time.Millisecond)
}

func calcSumBoxes(data *[][]rune) int {
	sum := 0
	for indRow, row := range *data {
		for indCol, col := range row {
			if col == '[' {
				sum += indRow*100 + indCol
			}
		}
	}
	return sum
}

func main() {
	data, actions := parseInputFile()

	printMap(data)

	executeActions(data, actions)
	println("Part 2:", calcSumBoxes(data))
}
