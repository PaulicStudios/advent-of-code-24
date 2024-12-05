package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
	"strings"
)

var order [][2]int
var lists [][]int

func parseInputFile() {
	file, err := os.Open("05/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		split := strings.Split(line, "|")
		order = append(order, [2]int{utils.ConvertToInt(split[0]), utils.ConvertToInt(split[1])})
	}
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ",")

		list := make([]int, len(split))
		for ind, s := range split {
			nbr := utils.ConvertToInt(s)
			list[ind] = nbr
		}
		lists = append(lists, list)
	}
}

func checkNbrOrder(first, after int, list []int) (bool, int) {
	foundLast := false

	for ind, nbr := range list {
		if nbr == first {
			return !foundLast, ind
		}
		if nbr == after {
			foundLast = true
		}
	}

	return true, -1
}

func validList(list []int) bool {
	for _, ord := range order {
		valid, _ := checkNbrOrder(ord[0], ord[1], list)
		if !valid {
			return false
		}
	}
	return true
}

func getMiddleNumber(list []int) int {
	return list[len(list)/2]
}

func sortOrdering() {
	for _, list := range lists {
		for _, ord := range order {
			valid, wrongInd := checkNbrOrder(ord[0], ord[1], list)
			if valid {
				continue
			}

			listcopy := make([]int, len(list))
			copy(listcopy, list)
			for i := 0; i < len(listcopy); i++ {
				if listcopy[i] == ord[1] {
					listcopy[i] = ord[0]
				}
			}
			listcopy[wrongInd] = ord[1]
			copy(list, listcopy)
		}
	}
}

func getWrongOrdered() [][]int {
	wrongOrdered := make([][]int, 0)
	for _, list := range lists {
		sorted := true
		for _, ord := range order {
			valid, _ := checkNbrOrder(ord[0], ord[1], list)
			if !valid {
				sorted = false
			}
		}
		if !sorted {
			wrongOrdered = append(wrongOrdered, list)
		}
	}
	return wrongOrdered
}

func countMiddleNumbers(newList [][]int) int {
	sum := 0

	for _, list := range newList {
		if validList(list) {
			sum += getMiddleNumber(list)
		}
	}

	return sum
}

func main() {
	parseInputFile()

	println("Part 1: ", countMiddleNumbers(lists))

	wrongOrdered := getWrongOrdered()
	lists = wrongOrdered

	for {
		if len(getWrongOrdered()) == 0 {
			break
		}
		sortOrdering()
	}

	println("Part 2: ", countMiddleNumbers(lists))
}
