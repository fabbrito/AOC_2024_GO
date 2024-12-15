// day08.go
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

	return grid, height, width, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Position struct{ r, c int }

func (p Position) isInside(height, width int) bool {
	return p.r >= 0 && p.r < height && p.c >= 0 && p.c < width
}

func nthAntinode(a1, a2 Position, n int) Position {
	return Position{(n+1)*a1.r - n*a2.r, (n+1)*a1.c - n*a2.c}
}

func solvePart1(data [][]rune, height, width int) int {
	antennaMap := make(map[rune][]Position)
	for r, dataLine := range data {
		for c, char := range dataLine {
			if char != '.' {
				if _, ok := antennaMap[char]; !ok {
					antennaMap[char] = []Position{}
				}
				antennaMap[char] = append(antennaMap[char], Position{r, c})
			}
		}
	}
	antinodes := make(map[Position]struct{})
	for _, antennas := range antennaMap {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				an0 := nthAntinode(antennas[i], antennas[j], 1)
				if an0.isInside(height, width) {
					antinodes[an0] = struct{}{}
				}
				an1 := nthAntinode(antennas[j], antennas[i], 1)
				if an1.isInside(height, width) {
					antinodes[an1] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}

func solvePart2(data [][]rune, height, width int) int {
	antennaMap := make(map[rune][]Position)
	for r, dataLine := range data {
		for c, char := range dataLine {
			if char != '.' {
				if _, ok := antennaMap[char]; !ok {
					antennaMap[char] = []Position{}
				}
				antennaMap[char] = append(antennaMap[char], Position{r, c})
			}
		}
	}
	antinodes := make(map[Position]struct{})
	for _, antennas := range antennaMap {
		for i := 0; i < len(antennas)-1; i++ {
			for j := i + 1; j < len(antennas); j++ {
				curr := make(map[Position]struct{})
				antinode := nthAntinode(antennas[i], antennas[j], 0)
				nth := 0
				for antinode.isInside(height, width) {
					curr[antinode] = struct{}{}
					antinodes[antinode] = struct{}{}
					nth++
					antinode = nthAntinode(antennas[i], antennas[j], nth)
				}

				antinode = nthAntinode(antennas[j], antennas[i], 0)
				nth = 0
				for antinode.isInside(height, width) {
					curr[antinode] = struct{}{}
					antinodes[antinode] = struct{}{}
					nth++
					antinode = nthAntinode(antennas[j], antennas[i], nth)
				}
			}
		}
	}
	return len(antinodes)
}

func main() {
	data, height, width, err := parseFile("test.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(data, height, width))
	fmt.Println("Part 2:", solvePart2(data, height, width))
}
