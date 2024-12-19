package main

import (
	"adventofcode24/utils"
	"bufio"
	"math"
	"os"
	"strings"
)

type Arcade struct {
	buttonA [2]int
	buttonB [2]int
	prize   [2]int
}

func parseInputFile() *[]Arcade {
	file, err := os.Open("13/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data := make([]Arcade, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		arcade := Arcade{}

		split := strings.Split(line, ",")
		arcade.buttonA = [2]int{utils.ConvertToInt(strings.Split(split[0], "+")[1]), utils.ConvertToInt(strings.Split(split[1], "+")[1])}
		scanner.Scan()
		line = scanner.Text()
		split = strings.Split(line, ",")
		arcade.buttonB = [2]int{utils.ConvertToInt(strings.Split(split[0], "+")[1]), utils.ConvertToInt(strings.Split(split[1], "+")[1])}
		scanner.Scan()
		line = scanner.Text()
		split = strings.Split(line, ",")
		arcade.prize = [2]int{utils.ConvertToInt(strings.Split(split[0], "=")[1]), utils.ConvertToInt(strings.Split(split[1], "=")[1])}

		data = append(data, arcade)
	}

	return &data
}

func cheapestTokenArcade(arcade *Arcade, indAStart int, indBStart int) int {
	minCost := math.MaxInt

	for indA := 0; indA < 201; indA++ {
		for indB := 0; indB < 201; indB++ {
			x := (indA+indAStart)*arcade.buttonA[0] + (indB+indBStart)*arcade.buttonB[0]
			y := (indA+indAStart)*arcade.buttonA[1] + (indB+indBStart)*arcade.buttonB[1]
			if x == arcade.prize[0] && y == arcade.prize[1] {
				cost := (indA+indAStart)*3 + (indB + indBStart)
				if cost < minCost {
					minCost = cost
				}
				continue
			}

			if x > arcade.prize[0] || y > arcade.prize[1] {
				break
			}
		}
	}

	if minCost == math.MaxInt {
		return 0
	}

	return minCost
}

func cheapestTokenArcadePart2(arcade *Arcade, indA int, indB int, cheapest *int) int {
	cost := indA*3 + indB

	if cost >= *cheapest {
		return cost
	}

	x := indA*arcade.buttonA[0] + indB*arcade.buttonB[0]
	y := indA*arcade.buttonA[1] + indB*arcade.buttonB[1]

	if x == arcade.prize[0] && y == arcade.prize[1] {
		*cheapest = cost
		return cost
	}

	if x > arcade.prize[0] || y > arcade.prize[1] {
		return cost
	}

	cost1 := cheapestTokenArcadePart2(arcade, indA+1, indB, cheapest)
	cost2 := cheapestTokenArcadePart2(arcade, indA, indB+1, cheapest)

	if cost1 < cost2 {
		return cost1
	}

	return cost2
}

func cheapestTokenArcadePart2V2(arcade *Arcade) int {
	arcade.prize[0] += 10000000000000
	arcade.prize[1] += 10000000000000

	a := (arcade.prize[0]*arcade.buttonB[1] - arcade.prize[1]*arcade.buttonB[0]) / (arcade.buttonA[0]*arcade.buttonB[1] - arcade.buttonA[1]*arcade.buttonB[0])
	b := (arcade.prize[0]*arcade.buttonA[1] - arcade.prize[1]*arcade.buttonA[0]) / (arcade.buttonB[0]*arcade.buttonA[1] - arcade.buttonB[1]*arcade.buttonA[0])

	if a*arcade.buttonA[0]+b*arcade.buttonB[0] == arcade.prize[0] && a*arcade.buttonA[1]+b*arcade.buttonB[1] == arcade.prize[1] {
		return 3*a + b
	}
	return 0
}

func sumTokens(arcade *[]Arcade, part2 bool) int {
	sum := 0
	for _, a := range *arcade {
		if part2 {
			sum += cheapestTokenArcadePart2V2(&a)
		} else {
			sum += cheapestTokenArcade(&a, 0, 0)
		}
		print(".")
	}
	return sum
}

func main() {
	data := parseInputFile()

	println("Part 1: ", sumTokens(data, false))
	println("Part 2: ", sumTokens(data, true))
}
