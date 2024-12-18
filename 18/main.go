package main

import (
	"adventofcode24/utils"
	"bufio"
	"math"
	"os"
	"strings"
	"time"
)

func parseInputFile() *[][2]int {
	file, err := os.Open("18/test.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data := make([][2]int, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		data = append(data, [2]int{utils.ConvertToInt(split[0]), utils.ConvertToInt(split[1])})
	}

	return &data
}

func createCoruptedMap(data *[][2]int, corruptedBytes, size int) *[][]rune {
	mapp := make([][]rune, size)
	for i := range mapp {
		mapp[i] = make([]rune, size)
		for j := range mapp[i] {
			mapp[i][j] = '.'
		}
	}

	for i := 0; i < corruptedBytes; i++ {
		mapp[(*data)[i][1]][(*data)[i][0]] = '#'
	}

	return &mapp
}

func minimumSteps(mapp *[][]rune, row, col, steps int, visited [][]bool, scoreMap *[][]int, minSteps *int) (bool, int) {
	if row < 0 || col < 0 || row >= len(*mapp) || col >= len((*mapp)[0]) {
		return false, 0
	}

	if (*mapp)[row][col] == '#' {
		return false, 0
	}

	if visited[row][col] {
		return false, 0
	}

	visited[row][col] = true

	printMap(mapp, row, col)

	if row == len(*mapp)-1 && col == len((*mapp)[0])-1 {
		if *minSteps > steps {
			*minSteps = steps
		}
		printMapWithVisited(mapp, visited)
		return true, steps
	}

	foundAny := false
	stepsAny := math.MaxInt

	found, step := minimumSteps(mapp, row+1, col, steps+1, visited, scoreMap, minSteps)
	if found {
		(*scoreMap)[row][col] = step
		foundAny = true
		stepsAny = step
	}
	found, step = minimumSteps(mapp, row-1, col, steps+1, visited, scoreMap, minSteps)
	if found {
		(*scoreMap)[row][col] = step
		foundAny = true
		if step < stepsAny {
			stepsAny = step
		}
	}
	found, step = minimumSteps(mapp, row, col+1, steps+1, visited, scoreMap, minSteps)
	if found {
		(*scoreMap)[row][col] = step
		foundAny = true
		if step < stepsAny {
			stepsAny = step
		}
	}
	found, step = minimumSteps(mapp, row, col-1, steps+1, visited, scoreMap, minSteps)
	if found {
		(*scoreMap)[row][col] = step
		foundAny = true
		if step < stepsAny {
			stepsAny = step
		}
	}

	return foundAny, stepsAny
}

func findMinSteps(scoreMap *[][]int, row, col int, visited [][]bool) int {
	if row < 0 || col < 0 || row >= len(*scoreMap) || col >= len((*scoreMap)[0]) {
		return math.MaxInt
	}
	if (*scoreMap)[row][col] == math.MaxInt {
		return math.MaxInt
	}
	if visited[row][col] {
		return math.MaxInt
	}

	visited[row][col] = true

	stepsUp := findMinSteps(scoreMap, row-1, col, visited)
	stepsDown := findMinSteps(scoreMap, row+1, col, visited)
	stepsLeft := findMinSteps(scoreMap, row, col-1, visited)
	stepsRight := findMinSteps(scoreMap, row, col+1, visited)

	return int(math.Min(math.Min(float64(stepsUp), float64(stepsDown)), math.Min(float64(stepsLeft), float64(stepsRight))))
}

func printMap(mapp *[][]rune, row, col int) {
	for i := range *mapp {
		for j := range (*mapp)[i] {
			if i == row && j == col {
				print("X")
			} else {
				print(string((*mapp)[i][j]))
			}
		}
		println()
	}
	time.Sleep(50 * time.Millisecond)
}

func printMapWithVisited(mapp *[][]rune, visited [][]bool) {
	for i := range *mapp {
		for j := range (*mapp)[i] {
			if visited[i][j] {
				print("X")
			} else {
				print(string((*mapp)[i][j]))
			}
		}
		println()
	}
}

func main() {
	data := parseInputFile()

	corruptedMap := createCoruptedMap(data, 12, 7)

	printMap(corruptedMap, 0, 0)

	minSteps := math.MaxInt
	visited := make([][]bool, len(*corruptedMap))
	for i := range visited {
		visited[i] = make([]bool, len((*corruptedMap)[i]))
	}
	scoreMap := make([][]int, len(*corruptedMap))
	for i := range scoreMap {
		scoreMap[i] = make([]int, len((*corruptedMap)[i]))
		for j := range scoreMap[i] {
			scoreMap[i][j] = math.MaxInt
		}
	}
	minimumSteps(corruptedMap, 0, 0, 0, visited, &scoreMap, &minSteps)

	println(findMinSteps(&scoreMap, 0, 0, visited))
}
