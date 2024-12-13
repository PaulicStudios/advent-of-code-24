package main

import (
	"bufio"
	"os"
	"unicode"
)

func parseInputFile() [][]rune {
	file, err := os.Open("12/input.txt")
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
	if x < 0 || y < 0 || x >= len(*data) || y >= len((*data)[0]) {
		return
	}
	if (*data)[x][y] != cur {
		return
	}

	(*data)[x][y] = unicode.ToLower(cur)

	(*landsList) = append((*landsList), [2]int{x, y})

	floodFill(cur, data, x+1, y, landsList)
	floodFill(cur, data, x-1, y, landsList)
	floodFill(cur, data, x, y+1, landsList)
	floodFill(cur, data, x, y-1, landsList)
}

func checkIfOutside(cur rune, data *[][]rune, x, y int) bool {
	if x < -1 || x > len(*data) || y < -1 || y > len((*data)[0]) {
		return false
	}

	if x == -1 || x == len(*data) || y == -1 || y == len((*data)[0]) {
		return true
	}
	return (*data)[x][y] != cur
}

func countPerimeter(cur rune, data *[][]rune, lands *[][2]int) int {
	fences := 0
	for _, land := range *lands {
		if checkIfOutside(cur, data, land[0]+1, land[1]) {
			fences++
		}
		if checkIfOutside(cur, data, land[0]-1, land[1]) {
			fences++
		}
		if checkIfOutside(cur, data, land[0], land[1]+1) {
			fences++
		}
		if checkIfOutside(cur, data, land[0], land[1]-1) {
			fences++
		}
	}

	println("checking for ", string(cur))
	println("alreadyChecked", fences)
	println("lands", len(*lands))

	return fences
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
	for _, land := range landsMap {
		newCost := countPerimeter((*data)[land[0][0]][land[0][1]], data, &land) * len(land)
		println(string((*data)[land[0][0]][land[0][1]]), newCost)
		cost += newCost
	}

	return cost
}

func main() {
	data := parseInputFile()

	loopEachField(&data)

	println("Part 1: ", calcFenceCost(&data))
}
