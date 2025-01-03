package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
	"sort"
	"strings"
)

var list1 []int
var list2 []int

func parseInputFile() {
	file, err := os.Open("01/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "   ")

		list1 = append(list1, utils.ConvertToInt(split[0]))
		list2 = append(list2, utils.ConvertToInt(split[1]))
	}
}

func calcDistance() int {
	var distanceSum int
	for i := 0; i < len(list1); i++ {
		distance := list1[i] - list2[i]

		distanceSum += utils.Abs(distance)
	}

	return distanceSum
}

func timesNumberAppears(nbr int) int {
	var times int
	for i := 0; i < len(list2); i++ {
		if list2[i] > nbr {
			break
		}

		if list2[i] == nbr {
			times++
		}
	}

	return times
}

func calcSimilarity() int {
	var similaritySum int
	for i := 0; i < len(list1); i++ {
		similaritySum += list1[i] * timesNumberAppears(list1[i])
	}

	return similaritySum
}

func main() {
	parseInputFile()
	sort.Ints(list1)
	sort.Ints(list2)

	println("Part 1: ", calcDistance())
	println("Part 2: ", calcSimilarity())
}
