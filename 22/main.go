package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
)

func parseInputFile() *[]int {
	file, err := os.Open("22/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	nbrs := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		nbrs = append(nbrs, utils.ConvertToInt(line))
	}

	return &nbrs
}

func nextSecretNumber(secretNbr int) int {
	secretNbr ^= secretNbr * 64
	secretNbr %= 16777216

	secretNbr ^= secretNbr / 32
	secretNbr %= 16777216

	secretNbr ^= secretNbr * 2048
	secretNbr %= 16777216

	return secretNbr
}

func main() {
	nbrs := parseInputFile()

	for i := 0; i < 2000; i++ {
		for ind, nbr := range *nbrs {
			(*nbrs)[ind] = nextSecretNumber(nbr)
		}
	}

	sum := 0
	for _, nbr := range *nbrs {
		sum += nbr
	}

	println(sum)
}
