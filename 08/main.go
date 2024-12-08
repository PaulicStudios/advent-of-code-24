package main

import (
	"bufio"
	"os"
)

func parseInputFile() []string {
	file, err := os.Open("08/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		data = append(data, line)
	}

	return data
}

func setNotes(antenna rune, indRow, indCol, indRowA, indColA int, data *[]string, antennaMap *[]string) {
	diffRow := indRow - indRowA
	diffCol := indCol - indColA

	lowRow := indRow - diffRow
	lowCol := indCol - diffCol
	highRow := indRow + diffRow
	highCol := indCol + diffCol

	if lowRow >= 0 && lowCol >= 0 && lowRow < len(*data) && lowCol < len((*data)[0]) {
		row := []rune((*data)[lowRow])
		if row[lowCol] == '.' {
			row[lowCol] = '#'
			(*data)[lowRow] = string(row)
		}
		rowAntenna := []rune((*antennaMap)[lowRow])
		if rowAntenna[lowCol] == '.' && row[lowCol] != antenna {
			rowAntenna[lowCol] = '#'
			(*antennaMap)[lowRow] = string(rowAntenna)
		}
		//setNotes(antenna, indRow, indCol, indRowA, indColA, data, antennaMap)
	}
	if highRow >= 0 && highCol >= 0 && highRow < len(*data) && highCol < len((*data)[0]) {
		row := []rune((*data)[highRow])
		if row[highCol] == '.' {
			row[highCol] = '#'
			(*data)[highRow] = string(row)
		}
		rowAntenna := []rune((*antennaMap)[highRow])
		if rowAntenna[highCol] == '.' && row[highCol] != antenna {
			rowAntenna[highCol] = '#'
			(*antennaMap)[highRow] = string(rowAntenna)
		}
	}
}

func searchNotes(antenna rune, indRowA int, indColA int, data *[]string, antennaMap *[]string) {
	for indRow, row := range *data {
		for indCol, col := range row {
			if col != antenna || (indRow == indRowA && indCol == indColA) {
				continue
			}
			setNotes(antenna, indRow, indCol, indRowA, indColA, data, antennaMap)
		}
	}
}

func searchAntenna(data *[]string) *[]string {
	antinodeMap := make([]string, len(*data))
	for ind, _ := range *data {
		for _, _ = range (*data)[0] {
			antinodeMap[ind] += "."
		}
	}
	for indRow, row := range *data {
		for indCol, col := range row {
			if col == '.' || col == '#' {
				continue
			}
			searchNotes(col, indRow, indCol, data, &antinodeMap)
		}
	}
	return &antinodeMap
}

func countNotes(data *[]string) int {
	count := 0
	for _, row := range *data {
		for _, col := range row {
			if col == '#' {
				count++
			}
		}
	}
	return count
}

func main() {
	data := parseInputFile()

	println("Part 1: ", countNotes(searchAntenna(&data)))
	println("Part 2: ", data[1])
}
