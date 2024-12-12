package main

import (
	"bufio"
	"os"
	"unicode"
)

func parseInputFile() [][]rune {
	file, err := os.Open("12/test.txt")
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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, char := range line {
			row[i] = char
		}
		data = append(data, row)
	}

	return data
}

var landsMap = make([][][2]int, 0)

func floodFill(cur rune, data *[][]rune, x, y int, landsList *[][2]int) {
	if (*data)[x][y] != cur {
		return
	}

	(*data)[x][y] = unicode.ToLower(cur)

	(*landsList) = append((*landsList), [2]int{x, y})

	x = x + 1
	if x < len(*data) {
		floodFill(cur, data, x, y, landsList)
	}

	x = x - 1
	if x >= 0 {
		floodFill(cur, data, x, y, landsList)
	}

	y = y + 1
	if y < len((*data)[x]) {
		floodFill(cur, data, x, y, landsList)
	}

	y = y - 1
	if y >= 0 {
		floodFill(cur, data, x, y, landsList)
	}
}

func checkIfOutside(cur rune, data *[][]rune, x, y int, alreadyChecked *[][2]int) {
	if x < -1 || x > len(*data) || y < -1 || y > len((*data)[0]) {
		return
	}

	if x == -1 || x == len(*data) || y == -1 || y == len((*data)[0]) {
		// for _, alreadyChecked := range *alreadyChecked {
		// 	if alreadyChecked[0] == x && alreadyChecked[1] == y {
		// 		return
		// 	}
		// }
		(*alreadyChecked) = append((*alreadyChecked), [2]int{x, y})
		return
	}

	if (*data)[x][y] == unicode.ToLower(cur) {
		return
	}

	// for _, alreadyChecked := range *alreadyChecked {
	// 	if alreadyChecked[0] == x && alreadyChecked[1] == y {
	// 		return
	// 	}
	// }

	(*alreadyChecked) = append((*alreadyChecked), [2]int{x, y})
}

func countPerimeter(cur rune, data *[][]rune, lands *[][2]int) int {

	alreadyChecked := make([][2]int, 0)
	for _, land := range *lands {
		checkIfOutside(cur, data, land[0]+1, land[1], &alreadyChecked)
		checkIfOutside(cur, data, land[0]-1, land[1], &alreadyChecked)
		checkIfOutside(cur, data, land[0], land[1]+1, &alreadyChecked)
		checkIfOutside(cur, data, land[0], land[1]-1, &alreadyChecked)
		// Also check Diagonal
		// checkIfOutside(cur, data, land[0]+1, land[1]+1, &alreadyChecked)
		// checkIfOutside(cur, data, land[0]-1, land[1]-1, &alreadyChecked)
		// checkIfOutside(cur, data, land[0]+1, land[1]-1, &alreadyChecked)
		// checkIfOutside(cur, data, land[0]-1, land[1]+1, &alreadyChecked)
	}

	println("checking for ", string(cur))
	println("alreadyChecked", len(alreadyChecked))
	println("lands", len(*lands))

	return len(alreadyChecked)
}

func loopEachField(data *[][]rune) {
	for rowInd, row := range *data {
		for colInd, col := range row {
			if unicode.IsLower(col) {
				continue
			}

			landsList := make([][2]int, 0)
			floodFill(col, data, rowInd, colInd, &landsList)
			landsMap = append(landsMap, landsList)
		}
	}
}

func calcFenceCost(data *[][]rune) int {
	cost := 0
	for ind, land := range landsMap {
		newCost := countPerimeter((*data)[ind][ind], data, &land) * len(land)
		println(string((*data)[ind][ind]), newCost)
		cost += newCost
	}

	return cost
}

func main() {
	data := parseInputFile()

	loopEachField(&data)

	println("Part 1: ", calcFenceCost(&data))
}
