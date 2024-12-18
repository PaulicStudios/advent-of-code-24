package main

import (
	"bufio"
	"math"
	"os"
	"sort"
	"time"
)

func parseInputFile() *[][]rune {
	file, err := os.Open("16/input.txt")
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

		data = append(data, []rune(line))
	}

	return &data
}

func getRune(data *[][]rune, find rune) (int, int) {
	for indRow, row := range *data {
		for indCol, col := range row {
			if col == find {
				return indRow, indCol
			}
		}
	}
	return -1, -1
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
	cost     int
	path     []State
}

func solve(data *[][]rune) (int, [][]State) {
	startRow, startCol := getRune(data, 'S')
	endRow, endCol := getRune(data, 'E')

	// Priority queue of states
	pq := make([]State, 0)
	// Initial state facing right
	pq = append(pq, State{startRow, startCol, Right, 0, []State{}})

	// visited[row][col][direction] = minimum cost
	visited := make([][][]int, len(*data))
	for i := range visited {
		visited[i] = make([][]int, len((*data)[0]))
		for j := range visited[i] {
			visited[i][j] = make([]int, 4)
			for k := range visited[i][j] {
				visited[i][j][k] = math.MaxInt
			}
		}
	}

	paths := make([][]State, 0)
	bestScore := math.MaxInt

	for len(pq) > 0 {
		// Get state with minimum cost
		current := pq[0]
		pq = pq[1:]

		if current.row == endRow && current.col == endCol {
			if current.cost < bestScore {
				bestScore = current.cost
				// Clear previous paths if we found a better score
				paths = [][]State{}
			}
			// Only add path if it matches the best score
			if current.cost == bestScore {
				paths = append(paths, append(current.path, current))
			}
		}

		// Try all possible moves
		moves := []Direction{Up, Down, Left, Right}
		for _, newDir := range moves {
			turnCost := 0
			if newDir != current.dir {
				turnCost = 1000
			}

			newRow, newCol := current.row, current.col
			switch newDir {
			case Up:
				newRow--
			case Down:
				newRow++
			case Left:
				newCol--
			case Right:
				newCol++
			}

			// Check if move is valid
			if newRow < 0 || newRow >= len(*data) ||
				newCol < 0 || newCol >= len((*data)[0]) ||
				(*data)[newRow][newCol] == '#' {
				continue
			}

			newCost := current.cost + turnCost + 1
			if newCost <= visited[newRow][newCol][newDir] {
				visited[newRow][newCol][newDir] = newCost

				newPath := make([]State, len(current.path))
				copy(newPath, current.path)
				newPath = append(newPath, current)

				// Insert into priority queue (maintain sorting by cost)
				insertIndex := sort.Search(len(pq), func(i int) bool {
					return pq[i].cost > newCost
				})
				pq = append(pq, State{})
				copy(pq[insertIndex+1:], pq[insertIndex:])
				pq[insertIndex] = State{newRow, newCol, newDir, newCost, newPath}
			}
		}
	}

	return bestScore, paths
}

func printMap(data *[][]rune, row, col int) {
	// print("\033[H\033[2J")
	for i := 0; i < len(*data); i++ {
		for j := 0; j < len((*data)[i]); j++ {
			if i == row && j == col {
				print("X")
			} else {
				print(string((*data)[i][j]))
			}
		}
		println()
	}
	println()
	time.Sleep(500 * time.Millisecond)
}

func main() {
	data := parseInputFile()

	bestScore, paths := solve(data)

	visited := make([][]int, len(*data))
	for i := range visited {
		visited[i] = make([]int, len((*data)[i]))
	}

	for _, path := range paths {
		for _, state := range path {
			visited[state.row][state.col] += 1
		}
	}

	atLeastOnce := 0
	for _, row := range visited {
		for _, col := range row {
			if col > 0 {
				atLeastOnce += 1
			}
		}
	}

	println("Part 1:", bestScore)
	println("Part 2:", atLeastOnce)
}
