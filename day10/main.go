// day10.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) (grid [][]int, height, width int, err error) {
	width, height = 0, 0
	grid = [][]int{}
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if width == 0 {
			width = len(line) // Set the width based on the first line
		} else if len(line) != width {
			// Handle inconsistent widths
			err = fmt.Errorf("inconsistent line widths: expected %d, found %d", width, len(line))
			return
		}
		data := make([]int, width)
		for i, char := range line {
			if char == '.' {
				data[i] = -1
			} else {
				data[i] = int(char) - int('0')
			}
		}
		grid = append(grid, data)
		height++
	}

	err = scanner.Err()
	if err != nil {
		return
	}

	return
}

type Position struct{ r, c int }

func isInside(r, c, height, width int) bool {
	return (r >= 0) && (r < height) && (c >= 0) && (c < width)
}

func findTrailhead(data [][]int) []Position {
	trailheads := []Position{}
	for r, line := range data {
		for c, alt := range line {
			if alt == 0 {
				trailheads = append(trailheads, Position{r, c})
			}
		}
	}
	return trailheads
}

func computeNextUniquePosition(data [][]int, height, width int, pos Position, next *map[Position]struct{}) {
	for _, delta := range [...]Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		if rN, cN := pos.r+delta.r, pos.c+delta.c; isInside(rN, cN, height, width) {
			if data[rN][cN] == data[pos.r][pos.c]+1 {
				nextPos := Position{rN, cN}
				if _, ok := (*next)[nextPos]; !ok {
					(*next)[nextPos] = struct{}{}
				}
			}
		}
	}
}

func solvePart1(data [][]int, height, width int) int {
	result := 0

	trailheads := findTrailhead(data)

	for _, th := range trailheads {
		curr := make(map[Position]struct{})
		curr[th] = struct{}{}
		for alt := 1; alt <= 9; alt++ {
			next := make(map[Position]struct{})
			for pos := range curr {
				computeNextUniquePosition(data, height, width, pos, &next)
			}
			curr = next
		}
		result += len(curr)
	}
	return result
}

func computeNextPositionInTrail(data [][]int, height, width int, pos Position) []Position {
	nextPath := []Position{}
	for _, delta := range [...]Position{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		if rN, cN := pos.r+delta.r, pos.c+delta.c; isInside(rN, cN, height, width) {
			if data[rN][cN] == data[pos.r][pos.c]+1 {
				nextPath = append(nextPath, Position{rN, cN})
			}
		}
	}
	return nextPath
}

func solvePart2(data [][]int, height, width int) int {
	result := 0
	trailheads := findTrailhead(data)
	for _, th := range trailheads {
		curr := []Position{th}
		for alt := 1; alt <= 9; alt++ {
			next := []Position{}
			for _, pos := range curr {
				next = append(next, computeNextPositionInTrail(data, height, width, pos)...)
			}
			curr = next
		}
		result += len(curr)
	}
	return result
}

func main() {
	data, height, width, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("Part 1:", solvePart1(data, height, width))
	fmt.Println("Part 2:", solvePart2(data, height, width))
}
