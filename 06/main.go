package main

import (
	"bufio"
	"os"
)

func parseInputFile() [][]int32 {
	file, err := os.Open("06/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mapp := make([][]int32, 0)
	for scanner.Scan() {
		line := scanner.Text()

		intLine := make([]int32, len(line))
		for ind, char := range line {
			intLine[ind] = char
		}
		mapp = append(mapp, intLine)
	}

	return mapp
}

func executeStep(row, col int, mapp *[][]int32) (int, int) {
	switch (*mapp)[row][col] {
	case '^':
		if row == 0 {
			return -1, -1
		}
		return row - 1, col
	case '>':
		if col == len((*mapp)[row])-1 {
			return -1, -1
		}
		return row, col + 1
	case 'v':
		if row == len(*mapp)-1 {
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

func startPos(mapp *[][]int32) (int, int) {
	for indRow, row := range *mapp {
		for indCol, col := range row {
			if col == '^' {
				return indRow, indCol
			}
		}
	}
	return -1, -1
}

func walkLoop(mapp *[][]int32) bool {
	rowInd, colInd := startPos(mapp)
	if rowInd == -1 || colInd == -1 {
		return false
	}
	initRow, initCol := rowInd, colInd
	initFirstRow, initFirstCol := -1, -1

	count := 0
	countAgain := 0

	for {
		newRow, newCol := executeStep(rowInd, colInd, mapp)
		if newRow == -1 || newCol == -1 {
			break
		}
		if initFirstRow == -1 && initFirstCol == -1 {
			initFirstRow, initFirstCol = newRow, newCol
		}

		if newRow == initFirstRow && newCol == initFirstCol && rowInd == initRow && colInd == initCol {
			count++
			if count == 2 {
				return true
			}
		}

		obstruction := (*mapp)[newRow][newCol]
		if obstruction == '#' {
			(*mapp)[rowInd][colInd] = turnRight((*mapp)[rowInd][colInd])
		} else {
			(*mapp)[newRow][newCol] = (*mapp)[rowInd][colInd]
			(*mapp)[rowInd][colInd] = 'X'
			rowInd, colInd = newRow, newCol
		}
		countAgain++
		if countAgain > 10000 {
			return true
		}
	}
	return false
}

func countVisited(mapp *[][]int32) int {
	count := 1
	for _, row := range *mapp {
		for _, col := range row {
			if col == 'X' {
				count++
			}
		}
	}
	return count
}

func infiniteLoop(mapp [][]int32) int {
	infiniteLoopMapCount := 0

	copyMapp := make([][]int32, len(mapp))
	for ind, _ := range mapp {
		copyMapp[ind] = make([]int32, len(mapp[ind]))
	}

	for row := 0; row < len(mapp); row++ {
		for col := 0; col < len(mapp[row]); col++ {
			for ind, _ := range mapp {
				copy(copyMapp[ind], mapp[ind])
			}

			copyMapp[row][col] = '#'
			if walkLoop(&copyMapp) {
				println("Infinite loop at row: ", row, " col: ", col)
				infiniteLoopMapCount++
			}
		}
	}
	return infiniteLoopMapCount
}

func main() {
	mapp := parseInputFile()

	println("Part 2: ", infiniteLoop(mapp))

	//walkLoop(&mapp)
	//println("Part 1: ", countVisited(&mapp))
}
