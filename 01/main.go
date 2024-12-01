package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var list1 []int
var list2 []int

func convertToInt(s string) int {
	nbr, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return nbr
}

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

		list1 = append(list1, convertToInt(split[0]))
		list2 = append(list2, convertToInt(split[1]))
	}
}

func calcDistance() int {
	sort.Ints(list1)
	sort.Ints(list2)

	var distanceSum int
	for i := 0; i < len(list1); i++ {
		distance := list1[i] - list2[i]

		if distance < 0 {
			distance = -distance
		}
		distanceSum += distance
	}

	return distanceSum
}

func main() {
	parseInputFile()

	println(calcDistance())
}
