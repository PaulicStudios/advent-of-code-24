package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
	"strings"
)

func parseInputFile() ([]int, [][]int) {
	file, err := os.Open("07/test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := make([]int, 0)
	nbrs := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, ": ")

		sum = append(sum, utils.ConvertToInt(split[0]))

		splitnbrs := strings.Split(split[1], " ")
		nbrss := make([]int, len(splitnbrs))
		for ind, s := range splitnbrs {
			nbr := utils.ConvertToInt(s)
			nbrss[ind] = nbr
		}
		nbrs = append(nbrs, nbrss)
	}

	return sum, nbrs
}

func generateCombinations(chars []rune, length int, current []rune, index int, results *[]string) {
	if index == length {
		*results = append(*results, string(current))
		return
	}

	for _, char := range chars {
		current[index] = char
		generateCombinations(chars, length, current, index+1, results)
	}
}

func checkValidCalc(sum int, nbrsTemplate []int, combinations *[]string) bool {
	for _, combination := range *combinations {
		nbrs := make([]int, len(nbrsTemplate))
		copy(nbrs, nbrsTemplate)
		for ind, nbr := range nbrs {
			if ind == 0 {
				continue
			}
			if combination[ind] == '+' {
				nbrs[ind] = nbrs[ind-1] + nbr
			} else if combination[ind] == '*' {
				nbrs[ind] = nbrs[ind-1] * nbr
			}
		}
		if nbrs[len(nbrs)-1] == sum {
			return true
		}
	}

	return false
}

func calcSumValidNbrs(sum []int, nbrs [][]int) int {
	validSum := 0

	for ind, s := range sum {
		operators := []rune{'+', '*'}
		length := len(nbrs[ind])
		results := make([]string, 0)
		current := make([]rune, length)

		generateCombinations(operators, length, current, 0, &results)

		if checkValidCalc(s, nbrs[ind], &results) {
			validSum += s
		}
	}

	return validSum
}

func main() {
	sum, nbrs := parseInputFile()

	println("Part 1: ", calcSumValidNbrs(sum, nbrs))

}
