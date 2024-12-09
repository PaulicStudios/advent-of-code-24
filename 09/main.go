package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
)

func parseInputFile() []int {
	file, err := os.Open("09/test.txt")
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
			data = append(data, utils.ConvertToInt(string(char)))
		}
		return data
	}

	return []int{}
}

func blockString(data []int) [][]int {
	block := make([][]int, 0)

	for id, char := range data {
		if char == 0 {
			continue
		}
		if id%2 == 0 {
			innerBlock := make([]int, 0)
			for iter := 0; iter < char; iter++ {
				innerBlock = append(innerBlock, id/2)
			}
			block = append(block, innerBlock)
		} else {
			innerBlock := make([]int, 0)
			for iter := 0; iter < char; iter++ {
				innerBlock = append(innerBlock, -1)
			}
			block = append(block, innerBlock)
		}
	}
	return block
}

func filesToLeft(dataNew [][]int) [][]int {
	data := make([]int, 0)
	for _, innerBlock := range dataNew {
		for _, char := range innerBlock {
			data = append(data, char)
		}
	}
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

	return [][]int{data}
}

func filesToLeftBlock(data [][]int) [][]int {
	for indRev := len(data) - 1; indRev >= 0; indRev-- {
		if data[indRev][0] == -1 {
			continue
		}
		spaceNeeded := len(data[indRev])
		for ind := 0; ind < indRev; ind++ {
			spaceAvailable := 0
			startInd := 0
			for innerInd, char := range data[ind] {
				if char == -1 {
					spaceAvailable++
					if spaceAvailable >= spaceNeeded {
						revStrInd := 0
						for i := startInd; i < startInd+spaceNeeded; i++ {
							data[ind][i] = data[indRev][revStrInd]
							data[indRev][revStrInd] = -1
							revStrInd++
						}
						startInd = -1
						break
					}
				} else {
					spaceAvailable = 0
					startInd = innerInd + 1
				}
			}
			if startInd == -1 {
				break
			}
		}
	}

	return data
}

func checksum(data [][]int) int {
	sum := 0
	ind := 0
	for _, char := range data {
		for _, innerChar := range char {
			if innerChar == -1 {
				ind++
				continue
			}
			sum += ind * innerChar
			ind++
		}
	}
	return sum
}

func main() {
	data := parseInputFile()

	println("Part 1: ", checksum(filesToLeft(blockString(data))))
	println("Part 2: ", checksum(filesToLeftBlock(blockString(data))))
}
