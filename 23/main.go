package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Connection struct {
	pc1 string
	pc2 string
}

func parseInputFile() *[]Connection {
	file, err := os.Open("23/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	connections := make([]Connection, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, "-")
		connections = append(connections, Connection{
			pc1: split[0],
			pc2: split[1],
		})
	}

	return &connections
}

func getThreeCons(connections *[]Connection) *[][]string {
	threeCons := make([][]string, 0)
	for _, conn1 := range *connections {
		pc1Connections := make(map[string]bool)
		for _, conn2 := range *connections {
			if conn2.pc1 == conn1.pc1 {
				pc1Connections[conn2.pc2] = true
			}
			if conn2.pc2 == conn1.pc1 {
				pc1Connections[conn2.pc1] = true
			}
		}

		pc2Connections := make(map[string]bool)
		for _, conn2 := range *connections {
			if conn2.pc1 == conn1.pc2 {
				pc2Connections[conn2.pc2] = true
			}
			if conn2.pc2 == conn1.pc2 {
				pc2Connections[conn2.pc1] = true
			}
		}

		for pc := range pc1Connections {
			if pc2Connections[pc] {
				computers := []string{conn1.pc1, conn1.pc2, pc}
				slices.Sort(computers)

				found := false
				for _, existing := range threeCons {
					if slices.Equal(existing, computers) {
						found = true
						break
					}
				}
				if !found {
					threeCons = append(threeCons, computers)
				}
			}
		}
	}

	return &threeCons
}

func main() {
	connections := parseInputFile()
	threeCons := getThreeCons(connections)

	count := 0
	for _, computers := range *threeCons {
		hasT := false
		for _, pc := range computers {
			if strings.HasPrefix(pc, "t") {
				hasT = true
				break
			}
		}
		if hasT {
			count++
		}

		fmt.Println(strings.Join(computers, ","))
	}

	fmt.Println("Sets with 't':", count)
}
