// day17.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func parseFile(filename string) (registersData [3]int, program []int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regIndex := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && strings.HasPrefix(line, "Register") {
			reg, err := strconv.Atoi(strings.Split(line, ": ")[1])
			if err != nil {
				panic(err)
			}
			if regIndex < 3 {
				registersData[regIndex] = reg
				regIndex++
			}
		}
		if len(line) > 0 && strings.HasPrefix(line, "Program") {
			for _, strCode := range strings.Split(strings.Split(line, ": ")[1], ",") {
				code, err := strconv.Atoi(strCode)
				if err != nil {
					panic(err)
				}
				program = append(program, code)
			}
		}
	}

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	return
}

type Computer struct {
	ra, rb, rc int
	pc         int
	out        []int
	program    []int
}

func (c *Computer) combo(operand int) int {
	switch operand {
	case 4:
		return c.ra
	case 5:
		return c.rb
	case 6:
		return c.rc
	default:
		return operand
	}
}

func (c *Computer) exec(opcode, operand int) {
	c.pc += 2
	switch opcode {
	case 0: // adv
		c.ra = c.ra >> c.combo(operand)
	case 1: // bxl
		c.rb ^= operand
	case 2: // bst
		c.rb = c.combo(operand) % 8
	case 3: // jnz
		if c.ra != 0 {
			c.pc = operand
		}
	case 4: // bxc
		c.rb ^= c.rc
	case 5: // out
		c.out = append(c.out, c.combo(operand)%8)
	case 6: // bdv
		c.rb = c.ra >> c.combo(operand)
	case 7: // cdv
		c.rc = c.ra >> c.combo(operand)
	}
}

func (c *Computer) runProgram() {
	for c.pc <= len(c.program)-2 {
		c.exec(c.program[c.pc], c.program[c.pc+1])
	}
}

func (c *Computer) reset(ra, rb, rc int) {
	c.ra, c.rb, c.rc = ra, rb, rc
	c.pc = 0
	c.out = nil
}

func (c *Computer) output() (out string) {
	for i, num := range c.out {
		out += strconv.Itoa(num)
		if i < len(c.out)-1 {
			out += ","
		}
	}
	return out
}

func solvePart1(data [3]int, program []int) string {
	computer := Computer{ra: data[0], rb: data[1], rc: data[2], program: program}
	computer.runProgram()
	return computer.output()
}

func solvePart2(data [3]int, program []int) int {
	ra, rb, rc := 0, data[1], data[2]
	for i := len(program) - 1; i >= 0; i-- {
		ra <<= 3
		computer := Computer{ra: ra, rb: rb, rc: rc, program: program}
		computer.runProgram()
		for !reflect.DeepEqual(computer.out, program[i:]) {
			ra++
			computer.reset(ra, rb, rc)
			computer.runProgram()
		}
	}
	return ra
}

func main() {
	registersData, program := parseFile("input.txt")
	fmt.Println("Part 1: ", solvePart1(registersData, program))
	fmt.Println("Part 2: ", solvePart2(registersData, program))
}
