package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Robot struct {
	pos [2]int
	vel [2]int
}

func parseInputFile() *[]Robot {
	file, err := os.Open("14/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data := make([]Robot, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		robot := Robot{}

		split := strings.Split(line, " ")
		robot.pos = [2]int{utils.ConvertToInt(strings.Split(strings.Trim(split[0], "p="), ",")[0]), utils.ConvertToInt(strings.Split(strings.Trim(split[0], "p="), ",")[1])}
		robot.vel = [2]int{utils.ConvertToInt(strings.Split(strings.Trim(split[1], "v="), ",")[0]), utils.ConvertToInt(strings.Split(strings.Trim(split[1], "v="), ",")[1])}

		data = append(data, robot)
	}

	return &data
}

func moveRobots(data *[]Robot, maxX int, maxY int) {
	for ind, robot := range *data {
		(*data)[ind].pos[0] += robot.vel[0]
		if (*data)[ind].pos[0] < 0 {
			(*data)[ind].pos[0] += maxX
		}
		if (*data)[ind].pos[0] >= maxX {
			(*data)[ind].pos[0] -= maxX
		}

		(*data)[ind].pos[1] += robot.vel[1]
		if (*data)[ind].pos[1] < 0 {
			(*data)[ind].pos[1] += maxY
		}
		if (*data)[ind].pos[1] >= maxY {
			(*data)[ind].pos[1] -= maxY
		}
	}
}

func printRobots(data *[]Robot, maxX int, maxY int) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			if slices.IndexFunc(*data, func(r Robot) bool {
				return r.pos[0] == x && r.pos[1] == y
			}) != -1 {
				print("*")
			} else {
				print(" ")
			}
		}
		println()
	}
}

func countRobotsQuadrant(data *[]Robot, minX, maxX, minY, maxY int) (int, bool) {
	count := 0
	unique := true
	for indRobot, robot := range *data {
		if robot.pos[0] >= minX && robot.pos[0] <= maxX && robot.pos[1] >= minY && robot.pos[1] <= maxY {
			count++
			for indRobot2, robot2 := range *data {
				if indRobot2 != indRobot && robot.pos[0] == robot2.pos[0] && robot.pos[1] == robot2.pos[1] {
					unique = false
					break
				}
			}
		}
	}

	return count, unique
}

func calcSafetyScore(data *[]Robot, maxX int, maxY int) (int, bool) {
	first, uniqueFirst := countRobotsQuadrant(data, 0, maxX/2-1, 0, maxY/2-1)
	second, uniqueSecond := countRobotsQuadrant(data, maxX-maxX/2, maxX, 0, maxY/2-1)
	third, uniqueThird := countRobotsQuadrant(data, 0, maxX/2-1, maxY-maxY/2, maxY)
	fourth, uniqueFourth := countRobotsQuadrant(data, maxX-maxX/2, maxX, maxY-maxY/2, maxY)

	if uniqueFirst && uniqueSecond && uniqueThird && uniqueFourth {
		return first * second * third * fourth, true
	}

	return first * second * third * fourth, false
}

func main() {
	data := parseInputFile()

	for i := 0; i < 9000; i++ {
		moveRobots(data, 101, 103)
		score, unique := calcSafetyScore(data, 101, 103)
		if unique {
			printRobots(data, 101, 103)
			println("---- " + strconv.Itoa(i+1) + " ----")
		}
		if i == 99 {
			println("Part 1: ", score)
		}
	}

	// printRobots(data, 101, 103)
	// println(calcSafetyScore(data, 101, 103))
}
