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

//func checkOrder(startInd int, list []int) bool {
//	for _, ord := range order {
//		for i := startInd; i < len(list); i++ {
//			if list[i] > ord[1] {
//				list[i]
//			}
//		}
//	}
//}

func checkNbrOrder(first, after int, list []int) bool {
	foundLast := false

	for _, nbr := range list {
		if nbr == first {
			return !foundLast
		}
		if nbr == after {
			foundLast = true
		}
	}

	return true
}

func validList(list []int) bool {
	for _, ord := range order {
		if !checkNbrOrder(ord[0], ord[1], list) {
			return false
		}
	}
	return true
}

func getMiddleNumber(list []int) int {
	for i, nbr := range list {
		if i == len(list)/2 {
			return nbr
		}
	}
	return 0
}

func countMiddleNumbers() int {
	sum := 0

	for _, list := range lists {
		if validList(list) {
			sum += getMiddleNumber(list)
		}
	}

	return sum
}

func main() {
	parseInputFile()

	//for _, ord := range order {
	//	println(ord[0], ord[1])
	//}
	//for _, list := range lists {
	//	for _, nbr := range list {
	//		print(nbr, " ")
	//	}
	//}

	println("Part 1: ", countMiddleNumbers())
}
