package main

import (
	"bufio"
	"os"
)

var grid []string

func parseInputFile() {
	file, err := os.Open("04/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}
}

func countXMASPart1() int {
	rows := len(grid)
	cols := len(grid[0])
	word := "XMAS"
	wordLen := len(word)
	count := 0

	directions := [][2]int{
		{0, 1},   // right
		{1, 0},   // down
		{1, 1},   // down-right
		{1, -1},  // down-left
		{0, -1},  // left
		{-1, 0},  // up
		{-1, -1}, // up-left
		{-1, 1},  // up-right
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			for _, dir := range directions {
				found := true
				for k := 0; k < wordLen; k++ {
					nr := r + k*dir[0]
					nc := c + k*dir[1]
					if nr < 0 || nr >= rows || nc < 0 || nc >= cols || grid[nr][nc] != word[k] {
						found = false
						break
					}
				}
				if found {
					count++
				}
			}
		}
	}
	return count
}

func countXMASPart2(grid []string) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] == 'A' {
				str1 := ""
				str1 += string(grid[r-1][c-1])
				str1 += string(grid[r+1][c+1])

				str2 := ""
				str2 += string(grid[r-1][c+1])
				str2 += string(grid[r+1][c-1])

				if str1 == "MS" || str1 == "SM" {
					if str2 == "SM" || str2 == "MS" {
						count++
					}
				}
			}
		}
	}
	return count
}

func main() {
	parseInputFile()

	println("Part 1: ", countXMASPart1())
	println("Part 2: ", countXMASPart2(grid))
}
