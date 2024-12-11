package main

import (
	"adventofcode24/utils"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseInputFile() []uint64 {
	file, err := os.Open("11/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	data := make([]uint64, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		for _, char := range split {
			data = append(data, uint64(utils.ConvertToInt(char)))
		}
	}

	return data
}

var cache = make(map[string]int)
var iterMax = 25

func blinkOncePart2(val, iter uint64) int {
	if int(iter) >= iterMax {
		return 1
	}
	iter++

	if val == 0 {
		val = 1
		return blinkOncePart2(val, iter)
	}
	str := strconv.FormatUint(val, 10)
	if val >= 10 && len(str)%2 == 0 {
		// splitMiddle := len(str) / 2
		// splitLeft := str[:splitMiddle]
		// splitRight := str[splitMiddle:]
		// (*data)[ind] = utils.ConvertToInt(splitLeft)

		num1, _ := strconv.ParseUint(str[:(len(str)/2)], 10, 64)
		num2, _ := strconv.ParseUint(str[(len(str)/2):], 10, 64)

		count := 0
		if _, ok := cache[strconv.FormatUint(num1, 10)+"_"+strconv.FormatUint(iter, 10)]; !ok {
			cache[strconv.FormatUint(num1, 10)+"_"+strconv.FormatUint(iter, 10)] = blinkOncePart2(num1, iter)
		}
		if _, ok := cache[strconv.FormatUint(num2, 10)+"_"+strconv.FormatUint(iter, 10)]; !ok {
			cache[strconv.FormatUint(num2, 10)+"_"+strconv.FormatUint(iter, 10)] = blinkOncePart2(num2, iter)
		}
		count += cache[strconv.FormatUint(num1, 10)+"_"+strconv.FormatUint(iter, 10)]
		count += cache[strconv.FormatUint(num2, 10)+"_"+strconv.FormatUint(iter, 10)]

		return count
	}

	return blinkOncePart2(val*2024, iter)
}

// func blinkPart1(data *[]int) {
// 	count := 0

// 	for _, val := range *data {
// 		count += blinkOncePart2(val)
// 		print(".")
// 	}
// 	println(count)
// 	// for _, val := range *data {
// 	// 	print(val)
// 	// }
// 	// println()
// }

func main() {
	data := parseInputFile()

	count := 0
	for _, val := range data {
		count += blinkOncePart2(val, 0)
		println(count)
	}

	println("Part 1: ", count)

	iterMax = 75
	count = 0
	cache = make(map[string]int)
	for _, val := range data {
		count += blinkOncePart2(val, 0)
		println(count)
	}

	println("Part 2: ", count)
}

// func main() {
// 	data := parseInputFile()

// 	count := 0
// 	var wg sync.WaitGroup // Add a WaitGroup to wait for goroutines to finish

// 	for _, val := range data {
// 		wg.Add(1) // Increment the WaitGroup counter
// 		go func(v uint64) {
// 			defer wg.Done() // Decrement the counter when the goroutine completes
// 			count += blinkOncePart2(v, 0)
// 		}(val)
// 	}

// 	wg.Wait() // Wait for all goroutines to finish
// 	println("Part 2: ", count)
// }
