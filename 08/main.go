package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInputFile() []string {
	file, err := os.Open("08/input.txt")
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
	data := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()

		data = append(data, line)
	}

	return data
}

var part2 = false

func setNotesDiff(antenna rune, indRow, indCol int, data *[]string, antennaMap *[]string) bool {
	if indRow < 0 || indCol < 0 || indRow >= len(*data) || indCol >= len((*data)[0]) {
		return false
	}

	rowAntenna := []rune((*antennaMap)[indRow])
	if rowAntenna[indCol] != '#' {
		if !part2 && (*data)[indRow][indCol] == uint8(antenna) {
			return false
		}
		rowAntenna[indCol] = '#'
		(*antennaMap)[indRow] = string(rowAntenna)
	}

	return part2
}

func setNotes(antenna rune, indRow, indCol, indRowA, indColA int, data *[]string, antennaMap *[]string) {
	diffRow := indRow - indRowA
	diffCol := indCol - indColA

	lowRow := indRow - diffRow
	lowCol := indCol - diffCol
	highRow := indRow + diffRow
	highCol := indCol + diffCol

	for setNotesDiff(antenna, lowRow, lowCol, data, antennaMap) {
		lowRow -= diffRow
		lowCol -= diffCol
	}
	for setNotesDiff(antenna, highRow, highCol, data, antennaMap) {
		highRow += diffRow
		highCol += diffCol
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
	for ind := range *data {
		for range (*data)[0] {
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
		count += strings.Count(row, "#")
	}
	return count
}

func main() {
	data := parseInputFile()

	println("Part 1: ", countNotes(searchAntenna(&data)))
	part2 = true
	println("Part 2: ", countNotes(searchAntenna(&data)))
}
