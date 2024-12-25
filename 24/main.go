package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
	"strings"
)

type Connection struct {
	from1 string
	op    string
	from2 string
	to    string
	done  bool
}

func parseInputFile() (*map[string]bool, *[]Connection) {
	file, err := os.Open("24/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	variables := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		variables[line[0:3]] = line[len(line)-1] == '1'
	}

	connections := make([]Connection, 0)
	for scanner.Scan() {
		line := scanner.Text()

		split := strings.Split(line, " ")
		connections = append(connections, Connection{
			from1: split[0],
			op:    split[1],
			from2: split[2],
			to:    split[4],
			done:  false,
		})
	}

	return &variables, &connections
}

func processConnection(variables *map[string]bool, connection *Connection) bool {
	if _, ok := (*variables)[connection.from1]; !ok {
		return false
	}
	if _, ok := (*variables)[connection.from2]; !ok {
		return false
	}

	switch connection.op {
	case "AND":
		(*variables)[connection.to] = (*variables)[connection.from1] && (*variables)[connection.from2]
	case "OR":
		(*variables)[connection.to] = (*variables)[connection.from1] || (*variables)[connection.from2]
	case "XOR":
		(*variables)[connection.to] = (*variables)[connection.from1] != (*variables)[connection.from2]
	}

	return true
}

func zBitsToNumber(bits *[]bool) int {
	number := 0
	for i := 0; i < len(*bits); i++ {
		if (*bits)[i] {
			number += 1 << (len(*bits) - 1 - i)
		}
	}
	return number
}

func main() {
	variables, connections := parseInputFile()

	changed := true
	for changed {
		changed = false
		for ind, connection := range *connections {
			if connection.done {
				continue
			}

			if processConnection(variables, &connection) {
				(*connections)[ind].done = true
				changed = true
			}
		}
	}

	zcount := 0
	for name := range *variables {
		if name[0] == 'z' {
			zcount++
		}
	}

	bits := make([]bool, zcount)
	for name, value := range *variables {
		if name[0] == 'z' {
			index := utils.ConvertToInt(name[1:])
			bits[zcount-index-1] = value
		}
	}

	println(zBitsToNumber(&bits))
}
