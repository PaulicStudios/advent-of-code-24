package main

import (
	"bufio"
	"os"
)

func parseInputFile() []int {
	file, err := os.Open("09/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := make([]int, 0)
		for _, char := range line {
			data = append(data, int(char-'0'))
		}
		return data
	}

	return []int{}
}

func blockString(data []int) []int {
	block := make([]int, 0)
	//idnbr := make([]int, len(data))
	for id, char := range data {
		if id%2 == 0 {
			for iter := 0; iter < char; iter++ {
				block = append(block, id/2)
			}
		} else {
			for iter := 0; iter < char; iter++ {
				block = append(block, -1)
			}
		}
	}
	return block
}

func filesToLeft(data []int) []int {
	for indRev := len(data) - 1; indRev >= 0; indRev-- {
		if data[indRev] == -1 {
			continue
		}
		for ind := 0; ind < indRev; ind++ {
			if data[ind] == -1 {
				data[ind] = data[indRev]
				data[indRev] = -1
				break
			}
		}
	}

	return data
}

func checksum(data []int) int {
	sum := 0
	for ind, char := range data {
		if char == -1 {
			continue
		}
		sum += char * ind
	}
	return sum
}

func main() {
	data := parseInputFile()

	println("Part 1: ", checksum(filesToLeft(blockString(data))))
}
