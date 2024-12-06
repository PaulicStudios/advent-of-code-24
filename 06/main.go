package main

import (
	"bufio"
	"os"
)

var mapp [][]int32

func parseInputFile() {
	file, err := os.Open("06/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		intLine := make([]int32, len(line))
		for ind, char := range line {
			intLine[ind] = char
		}
		mapp = append(mapp, intLine)
	}
}

func executeStep(row, col int) (int, int) {
	switch mapp[row][col] {
	case '^':
		if row == 0 {
			return -1, -1
		}
		return row - 1, col
	case '>':
		if col == len(mapp[row])-1 {
			return -1, -1
		}
		return row, col + 1
	case 'v':
		if row == len(mapp)-1 {
			return -1, -1
		}
		return row + 1, col
	case '<':
		if col == 0 {
			return -1, -1
		}
		return row, col - 1
	}
	return -1, -1
}

func turnRight(direction int32) int32 {
	switch direction {
	case '^':
		return '>'
	case '>':
		return 'v'
	case 'v':
		return '<'
	case '<':
		return '^'
	}
	return 0
}

func startPos() (int, int) {
	for indRow, row := range mapp {
		for indCol, col := range row {
			if col == '^' {
				return indRow, indCol
			}
		}
	}
	return -1, -1
}

func walkLoop() {
	rowInd, colInd := startPos()
	for {
		newRow, newCol := executeStep(rowInd, colInd)
		if newRow == -1 || newCol == -1 {
			break
		}

		obstruction := mapp[newRow][newCol]
		if obstruction == '#' {
			mapp[rowInd][colInd] = turnRight(mapp[rowInd][colInd])
		} else {
			mapp[newRow][newCol] = mapp[rowInd][colInd]
			mapp[rowInd][colInd] = 'X'
			rowInd, colInd = newRow, newCol
		}
	}
}

func countVisited() int {
	count := 1
	for _, row := range mapp {
		for _, col := range row {
			if col == 'X' {
				count++
			}
		}
	}
	return count
}

func main() {
	parseInputFile()

	walkLoop()
	println("Part 1: ", countVisited())
}
