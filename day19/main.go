// day19.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseFile(filename string) (patterns []string, designs []string, err error) {
	file, errF := os.Open(filename)
	if errF != nil {
		err = errF
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if patterns == nil {
			// First line: split by commas
			patterns = strings.Split(line, ", ")
		} else if len(line) > 0 {
			// Subsequent lines: add to the other slice
			designs = append(designs, line)
		}
	}
	if errS := scanner.Err(); errS != nil {
		err = errS
		return
	}
	return
}

func solver(patterns *[]string, design string) int {
	n := len(design)
	dp := make([]int, n+1)
	dp[0] = 1                 // Base case: 1 way to form an empty string
	for i := 1; i <= n; i++ { // dp[i] = number of ways to create design[0:i]
		for _, pattern := range *patterns {
			p := len(pattern)
			if i >= p && design[i-p:i] == pattern {
				dp[i] += dp[i-p]
			}
		}
	}
	return dp[n] // The number of ways to form the entire design
}

func solvePart1(patterns []string, designs []string) int {
	startTime := time.Now()
	defer func() {
		fmt.Printf("Part 1 execution took %s\n", time.Since(startTime))
	}()
	designsCounts := make(map[string]int)
	for _, design := range designs {
		designsCounts[design] = solver(&patterns, design)
	}
	result := 0
	for _, count := range designsCounts {
		if count > 0 {
			result++
		}
	}
	return result
}

func solvePart2(patterns []string, designs []string) int {
	startTime := time.Now()
	defer func() {
		fmt.Printf("Part 2 execution took %s\n", time.Since(startTime))
	}()
	designsCounts := make(map[string]int)
	for _, design := range designs {
		designsCounts[design] = solver(&patterns, design)
	}
	result := 0
	for _, count := range designsCounts {
		result += count
	}
	return result
}

func main() {
	patterns, designs, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Part 1: ", solvePart1(patterns, designs))
	fmt.Println("Part 2: ", solvePart2(patterns, designs))
}
