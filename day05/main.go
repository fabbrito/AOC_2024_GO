// day05.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([][]int, [][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Parse first section - all lines until empty line
	rules := make([][]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			break // Found the section separator
		}

		numbers := strings.Split(line, "|")
		if len(numbers) != 2 {
			return nil, nil, fmt.Errorf("invalid format in first section, expected two numbers separated by |")
		}

		row := make([]int, 2)
		for i, numStr := range numbers {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing number in first section: %v", err)
			}
			row[i] = num
		}
		rules = append(rules, row)
	}

	if len(rules) == 0 {
		return nil, nil, fmt.Errorf("first section is empty")
	}

	// Parse second section
	pages := make([][]int, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		numbers := strings.Split(line, ",")
		row := make([]int, len(numbers))
		for i, numStr := range numbers {
			num, err := strconv.Atoi(strings.TrimSpace(numStr))
			if err != nil {
				return nil, nil, fmt.Errorf("error parsing number in second section: %v", err)
			}
			row[i] = num
		}
		pages = append(pages, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading file: %v", err)
	}

	if len(pages) == 0 {
		return nil, nil, fmt.Errorf("second section is empty")
	}

	return rules, pages, nil
}

func validateOrder(rules [][]int, sequence []int) bool {
	// Create position map
	pos := make(map[int]int)
	for i, v := range sequence {
		pos[v] = i
	}

	// Check each rule
	for _, rule := range rules {
		before, after := rule[0], rule[1]
		// If both elements exist in sequence
		if pBefore, okBefore := pos[before]; okBefore {
			if pAfter, okAfter := pos[after]; okAfter {
				if pBefore >= pAfter {
					return false
				}
			}
		}
	}
	return true
}

func fixOrder(rules [][]int, sequence []int) []int {
	// Create position map
	pos := make(map[int]int)
	for i, v := range sequence {
		pos[v] = i
	}

	// Check each rule
	for _, rule := range rules {
		before, after := rule[0], rule[1]
		// If both elements exist in sequence
		if pBefore, okBefore := pos[before]; okBefore {
			if pAfter, okAfter := pos[after]; okAfter {
				if pBefore >= pAfter {
					pos[before], pos[after] = pAfter, pBefore
				}
			}
		}
	}
	res := make([]int, len(sequence))
	for number, pos := range pos {
		res[pos] = number
	}
	return res
}

func solvePart1(rules, pages [][]int) int {
	res := 0
	for _, page := range pages {
		if validateOrder(rules, page) {
			res += page[len(page)/2]
		}
	}
	return res
}

func solvePart2(rules, pages [][]int) int {
	res := 0
	for _, page := range pages {
		if !validateOrder(rules, page) {
			count := 0
			for !validateOrder(rules, page) && (count < 10) {
				page = fixOrder(rules, page)
				count++
			}
			res += page[len(page)/2]
		}
	}
	return res
}

func main() {
	rules, pages, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("Part 1:", solvePart1(rules, pages))
	fmt.Println("Part 2:", solvePart2(rules, pages))
}
