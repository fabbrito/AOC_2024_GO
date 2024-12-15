// day07.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	allGoals := []int{}
	allValues := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		equation := strings.Split(line, ": ")
		result, err := strconv.Atoi(equation[0])
		if err != nil {
			return nil, nil, err
		}
		allGoals = append(allGoals, result)
		values := []int{}
		for _, valueString := range strings.Split(equation[1], " ") {
			value, err := strconv.Atoi(valueString)
			if err != nil {
				return nil, nil, err
			}
			values = append(values, value)
		}
		allValues = append(allValues, values)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}
	if len(allGoals) != len(allValues) {
		return nil, nil, fmt.Errorf("inconsistent lenght for results and allValues, check the input")
	}

	return allGoals, allValues, nil
}

func solvePart1(allGoals []int, allValues [][]int) int {
	total := 0
	for i, values := range allValues {
		curr := []int{values[0]}
		for _, value := range values[1:] {
			var next []int
			for _, acc := range curr {
				next = append(next, acc+value, acc*value)
			}
			curr = next
		}
		for _, acc := range curr {
			if acc == allGoals[i] {
				total += acc
				break
			}
		}
	}
	return total
}

func concat(a, b int) int {
	digits := 0
	temp := b
	for temp > 0 {
		digits++
		temp /= 10
	}
	for i := 0; i < digits; i++ {
		a *= 10
	}
	return a + b
}

func solvePart2(allGoals []int, allValues [][]int) int {
	total := 0
	for i, values := range allValues {
		curr := []int{values[0]}
		for _, value := range values[1:] {
			var next []int
			for _, acc := range curr {
				next = append(next, acc+value, acc*value, concat(acc, value))
			}
			curr = next
		}
		for _, acc := range curr {
			if acc == allGoals[i] {
				total += acc
				break
			}
		}
	}
	return total
}

func main() {
	allGoals, allValues, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(allGoals, allValues))
	fmt.Println("Part 2:", solvePart2(allGoals, allValues))
}
