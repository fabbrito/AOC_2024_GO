// day03.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func parseFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Append each line to the slice
	}

	return lines, scanner.Err()
}

func solvePart1(lines []string) int {
	pattern := `mul\((\d{1,3}),(\d{1,3})\)`
	regex := regexp.MustCompile(pattern)
	acc := 0
	for _, line := range lines {
		matches := regex.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			acc += num1 * num2
		}
	}
	return acc
}

func solvePart2(lines []string) int {
	doDontRegex := regexp.MustCompile(`do\(\)|don't\(\)`)
	mulRegex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	acc := 0
	lastCommand := "do()"
	for _, line := range lines {
		segments := doDontRegex.Split(line, -1)
		triggers := []string{lastCommand}
		triggersExtracted := doDontRegex.FindAllString(line, -1)
		triggers = append(triggers, triggersExtracted...)
		for i, segment := range segments {
			if triggers[i] == "do()" {
				matches := mulRegex.FindAllStringSubmatch(segment, -1)
				for _, match := range matches {
					num1, err1 := strconv.Atoi(match[1])
					num2, err2 := strconv.Atoi(match[2])
					if err1 == nil && err2 == nil {
						acc += num1 * num2
					}
				}
			}
		}
		lastCommand = triggers[len(triggers)-1]
	}
	return acc
}

func main() {
	lines, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(lines))

	// lines2, err := parseFile("input.txt")
	// if err != nil {
	// 	fmt.Printf("Error reading file: %v\n", err)
	// 	return
	// }
	fmt.Println("Part 2:", solvePart2(lines))
}
