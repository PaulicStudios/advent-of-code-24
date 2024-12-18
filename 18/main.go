package main

import (
	"adventofcode24/utils"
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func parseInputFile() *[][2]int {
	file, err := os.Open("18/input.txt")
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

func getWeight(visited *[][]int, row, col int) int {
	if row < 0 || col < 0 || row >= len(*visited) || col >= len((*visited)[0]) {
		return math.MaxInt
	}

	return (*visited)[row][col] + distanceGoal(visited, row, col)
}

func distanceGoal(mapp *[][]int, row, col int) int {
	return len(*mapp) - row + len((*mapp)[0]) - col
}

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type State struct {
	row, col int
	dir      Direction
	moves    int
}

func solve(mapp *[][]rune) int {
	ol := make([]State, 0)
	ol = append(ol, State{0, 0, Down, 0})

	cl := make([][][]int, len(*mapp))
	for i := range cl {
		cl[i] = make([][]int, len((*mapp)[i]))
		for j := range cl[i] {
			cl[i][j] = make([]int, 4)
			for k := range cl[i][j] {
				cl[i][j][k] = math.MaxInt
			}
		}
	}

	for len(ol) > 0 {
		current := ol[0]
		ol = ol[1:]

		if current.row == len(*mapp)-1 && current.col == len((*mapp)[0])-1 {
			return current.moves
		}

		dirs := []Direction{Up, Down, Left, Right}
		for _, dir := range dirs {
			newRow, newCol := current.row, current.col
			switch dir {
			case Up:
				newRow--
			case Down:
				newRow++
			case Left:
				newCol--
			case Right:
				newCol++
			}

			if newRow < 0 || newCol < 0 || newRow >= len(*mapp) || newCol >= len((*mapp)[0]) || (*mapp)[newRow][newCol] == '#' {
				continue
			}

			newMoves := current.moves + 1
			if newMoves < cl[newRow][newCol][dir] {
				cl[newRow][newCol][dir] = newMoves
				ol = append(ol, State{newRow, newCol, dir, newMoves})
				sort.Slice(ol, func(i, j int) bool {
					return ol[i].moves < ol[j].moves
				})
			}
		}
	}

	return math.MaxInt
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

func printMapWithVisited(visited *[][]int) {
	for i := range *visited {
		for j := range (*visited)[i] {
			print(strconv.Itoa((*visited)[i][j]))
		}
		println()
	}
	time.Sleep(50 * time.Millisecond)
}

func main() {
	data := parseInputFile()

	corruptedMap := createCoruptedMap(data, 1024, 71)
	println("Part 1: ", solve(corruptedMap))

	for i := 1024; i < len(*data); i++ {
		corruptedMap := createCoruptedMap(data, i, 71)

		// printMap(corruptedMap, 0, 0)

		// minSteps := math.MaxInt - 1
		// visited := make([][]bool, len(*corruptedMap))
		// for i := range visited {
		// 	visited[i] = make([]bool, len((*corruptedMap)[i]))
		// }
		// lastMinSteps := math.MaxInt
		// for minSteps < lastMinSteps {
		// 	lastMinSteps = minSteps
		// 	steps, _ := minimumSteps(corruptedMap, 0, 0, 0, visited, &minSteps)
		// 	println(steps)
		// }
		moves := solve(corruptedMap)
		println(moves)
		if moves == math.MaxInt {
			println("Breaking out of loop at ", i)
			i -= 1
			println("Part 2: ", (*data)[i][0], (*data)[i][1])
			break
		}
	}
}
