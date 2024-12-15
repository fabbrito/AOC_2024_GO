// day04.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) ([][]rune, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, -1, -1, err
	}
	defer file.Close()

	var grid [][]rune
	width, height := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if width == 0 {
			width = len(line) // Set the width based on the first line
		} else if len(line) != width {
			// Handle inconsistent widths
			return nil, -1, -1, fmt.Errorf("inconsistent line widths: expected %d, found %d", width, len(line))
		}
		grid = append(grid, []rune(line))
		height++
	}

	if err := scanner.Err(); err != nil {
		return nil, -1, -1, err
	}

	return grid, width, height, nil
}

type Grid struct {
	data   [][]rune
	width  int
	height int
}

func (g Grid) IsValidPos(r, c int) bool {
	return (r >= 0) && (c >= 0) && (r < g.height) && (c < g.width)
}

func (g Grid) IncrementRow(r, c int) int {
	return r + (c+1)/g.width
}

func (g Grid) IncrementCol(c int) int {
	return (c + 1) % g.width
}

type Direction int

const (
	NORTH Direction = iota
	NORTHEAST
	EAST
	SOUTHEAST
	SOUTH
	SOUTHWEST
	WEST
	NORTHWEST
)

func (d Direction) Move(r, c int) (int, int) {
	switch d {
	case NORTH:
		return r - 1, c
	case NORTHEAST:
		return r - 1, c + 1
	case EAST:
		return r, c + 1
	case SOUTHEAST:
		return r + 1, c + 1
	case SOUTH:
		return r + 1, c
	case SOUTHWEST:
		return r + 1, c - 1
	case WEST:
		return r, c - 1
	case NORTHWEST:
		return r - 1, c - 1
	default:
		return r, c
	}
}

func countXmas(grid Grid, rStart, cStart, ix int, xmas []rune, dir Direction) int {
	if ix >= len(xmas) {
		return 1
	}
	r, c := dir.Move(rStart, cStart)
	if grid.IsValidPos(r, c) && (grid.data[r][c] == xmas[ix]) {
		return countXmas(grid, r, c, ix+1, xmas, dir)
	}
	return 0
}

func solvePart1(data [][]rune, width, height int) int {
	grid := Grid{data: data, width: width, height: height}
	res := 0
	xmas := []rune("MAS")
	allDirections := [...]Direction{NORTH, NORTHEAST, EAST, SOUTHEAST, SOUTH, SOUTHWEST, WEST, NORTHWEST}
	for r := 0; r < grid.height; r++ {
		for c := 0; c < grid.width; c++ {
			if grid.data[r][c] == 'X' {
				for _, dir := range allDirections {
					res += countXmas(grid, r, c, 0, xmas, dir)
				}

			}
		}
	}
	return res
}

func solvePart2(data [][]rune, width, height int) int {
	patterns := [][]rune{
		{
			'M', '.', 'M',
			'.', 'A', '.',
			'S', '.', 'S',
		},
		{
			'M', '.', 'S',
			'.', 'A', '.',
			'M', '.', 'S',
		},
		{
			'S', '.', 'M',
			'.', 'A', '.',
			'S', '.', 'M',
		},
		{
			'S', '.', 'S',
			'.', 'A', '.',
			'M', '.', 'M',
		},
	}
	pSize := 3
	res := 0
	for r := 0; r <= height-pSize; r++ {
		for c := 0; c <= width-pSize; c++ {
			for _, pattern := range patterns {
				match := true
				for p := 0; p < len(pattern); p++ {
					if pattern[p] == '.' {
						continue
					}
					if pattern[p] != data[r+p/pSize][c+(p%pSize)] {
						match = false
						break
					}
				}
				if match {
					res++
				}
			}
		}
	}
	return res
}

func main() {
	data, width, height, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(data, width, height))
	fmt.Println("Part 2:", solvePart2(data, width, height))
}
