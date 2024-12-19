package main

import (
	"bufio"
	"os"
	"strings"
)

func parseInputFile() (*[]string, *[]string) {
	file, err := os.Open("19/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	available := make([]string, 0)
	required := make([]string, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	available = strings.Split(scanner.Text(), ", ")

	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()

		required = append(required, line)
	}

	return &available, &required
}

// func checkifCanBeMade(available *[]string, required string) bool {
// 	for indReq := 0; indReq < len(required); indReq++ {
// 		found := true
// 		skipChars := 0
// 		for indAv := 0; indAv < len(*available); indAv++ {
// 			curSkipChars := 0
// 			found = true
// 			for indAvChar := 0; indAvChar < len((*available)[indAv]); indAvChar++ {
// 				if indReq+indAvChar >= len(required) {
// 					found = false
// 					break
// 				}
// 				if required[indReq+indAvChar] != (*available)[indAv][indAvChar] {
// 					found = false
// 					break
// 				}
// 				curSkipChars++
// 			}
// 			if found {
// 				if curSkipChars > skipChars {
// 					skipChars = curSkipChars
// 				}
// 			}
// 		}
// 		if skipChars == 0 {
// 			return false
// 		} else {
// 			indReq += skipChars
// 		}
// 	}
// 	return true
// }

func checkifCanBeMade(available *[]string, required string) bool {
	dp := make([]bool, len(required)+1)
	dp[0] = true

	for i := 0; i <= len(required); i++ {
		if !dp[i] {
			continue
		}
		for _, avail := range *available {
			if i+len(avail) <= len(required) {
				valid := true
				for j := 0; j < len(avail); j++ {
					if avail[j] != required[i+j] {
						valid = false
						break
					}
				}
				if valid {
					dp[i+len(avail)] = true
				}
			}
		}
	}

	return dp[len(required)]
}

func howManyCombinations(available *[]string, required string, cache *map[string]int) int {
	if count, exists := (*cache)[required]; exists {
		return count
	}

	count := 0
	for _, avail := range *available {
		if len(avail) <= len(required) {
			valid := true
			for i := 0; i < len(avail); i++ {
				if avail[i] != required[i] {
					valid = false
					break
				}
			}
			if valid {
				if len(required[len(avail):]) == 0 {
					count++
				} else {
					count += howManyCombinations(available, required[len(avail):], cache)
				}
			}
		}
	}

	(*cache)[required] = count
	return count
}

func canMakeCount(available *[]string, required *[]string) int {
	count := 0
	for _, req := range *required {
		if checkifCanBeMade(available, req) {
			count++
		}
	}
	return count
}

func combinationsCount(available *[]string, required *[]string) int {
	count := 0
	cache := make(map[string]int)
	for _, req := range *required {
		count += howManyCombinations(available, req, &cache)
	}
	return count
}

func main() {
	available, required := parseInputFile()

	println("Part 1:", canMakeCount(available, required))
	println("Part 2:", combinationsCount(available, required))
}
