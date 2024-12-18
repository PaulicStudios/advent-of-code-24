package main

import (
	"adventofcode24/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func parseInputFile() (int, int, int, []int) {
	file, err := os.Open("17/input.txt")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	a, b, c, program := 0, 0, 0, make([]int, 0)

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	a = utils.ConvertToInt(strings.TrimLeft(scanner.Text(), "Register A: "))
	scanner.Scan()
	b = utils.ConvertToInt(strings.TrimLeft(scanner.Text(), "Register B: "))
	scanner.Scan()
	c = utils.ConvertToInt(strings.TrimLeft(scanner.Text(), "Register C: "))
	scanner.Scan()
	scanner.Scan()
	programStrings := strings.Split(strings.TrimLeft(scanner.Text(), "Program: "), ",")
	for _, p := range programStrings {
		program = append(program, utils.ConvertToInt(p))
	}

	return a, b, c, program
}

var instructionPointer = 0
var notIncrease = false
var out = make([]int, 0)

func getCombo(comboOp, a, b, c int) int {
	if comboOp <= 3 {
		return comboOp
	} else if comboOp == 4 {
		return a
	} else if comboOp == 5 {
		return b
	} else if comboOp == 6 {
		return c
	}
	panic("Invalid comboOp")
}

func opcode0(op int, a, b, c *int) {
	comboOp := getCombo(op, *a, *b, *c)
	result := *a / int(math.Pow(2, float64(comboOp)))
	*a = result
}

func opcode1(op int, _, b, _ *int) {
	*b = *b ^ op
}

func opcode2(op int, a, b, c *int) {
	*b = getCombo(op, *a, *b, *c) % 8
}

func opcode3(op int, a, b, c *int) {
	if *a == 0 || instructionPointer == op {
		return
	}
	notIncrease = true
	instructionPointer = op
}

func opcode4(op int, a, b, c *int) {
	*b = *b ^ *c
}

func opcode5(op int, a, b, c *int) {
	outCombo := getCombo(op, *a, *b, *c) % 8
	out = append(out, outCombo)
	println(outCombo)
}

func opcode6(op int, a, b, c *int) {
	comboOp := getCombo(op, *a, *b, *c)
	result := *a / int(math.Pow(2, float64(comboOp)))
	*b = result
}

func opcode7(op int, a, b, c *int) {
	comboOp := getCombo(op, *a, *b, *c)
	result := *a / int(math.Pow(2, float64(comboOp)))
	*c = result
}

func runProgram(a, b, c *int, program *[]int) {
	for instructionPointer < len(*program) {
		opcode := (*program)[instructionPointer]
		operand := (*program)[instructionPointer+1]
		switch opcode {
		case 0:
			opcode0(operand, a, b, c)
		case 1:
			opcode1(operand, a, b, c)
		case 2:
			opcode2(operand, a, b, c)
		case 3:
			opcode3(operand, a, b, c)
		case 4:
			opcode4(operand, a, b, c)
		case 5:
			opcode5(operand, a, b, c)
		case 6:
			opcode6(operand, a, b, c)
		case 7:
			opcode7(operand, a, b, c)
		}

		if notIncrease {
			notIncrease = false
			continue
		}
		instructionPointer += 2
	}
}

func main() {
	a, b, c, program := parseInputFile()

	runProgram(&a, &b, &c, &program)

	for _, o := range out {
		fmt.Print(o, ",")
	}
}
