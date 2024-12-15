// day06.go
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

type Position struct{ r, c int }

var moves = []Position{
	{-1, 0}, // UP
	{0, 1},  // RIGHT
	{1, 0},  // DOWN
	{0, -1}, // LEFT
}

type GuardState struct {
	pos Position
	dir int
}

func (guard GuardState) peek() (r, c int) {
	return guard.pos.r + moves[guard.dir].r, guard.pos.c + moves[guard.dir].c
}

func findGuard(data [][]rune, width, height int) GuardState {
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if data[r][c] == '^' {
				return GuardState{Position{r, c}, 0}
			}
		}
	}
	return GuardState{}
}

func solveSteps(data [][]rune, width, height int, guard GuardState) (int, map[Position]struct{}) {
	steps := make(map[Position]struct{})
	steps[guard.pos] = struct{}{}

	for guard.pos.r > 0 && guard.pos.r < height-1 && guard.pos.c > 0 && guard.pos.c < width-1 {
		rNext, cNext := guard.peek() // peek
		if data[rNext][cNext] == '#' {
			guard.dir = (guard.dir + 1) % len(moves) // turn
			rNext, cNext = guard.peek()              // peek
		}
		guard.pos.r, guard.pos.c = rNext, cNext // Step forward
		_, ok := steps[guard.pos]
		if !ok {
			steps[guard.pos] = struct{}{}
		}
	}
	return len(steps), steps
}

func solvePart1(data [][]rune, width, height int) int {
	guard := findGuard(data, width, height)
	count, _ := solveSteps(data, width, height, guard)
	return count
}

func isLoop(data [][]rune, width, height int, guard GuardState) bool {
	visited := make(map[GuardState]struct{})
	visited[guard] = struct{}{}

	for guard.pos.r > 0 && guard.pos.r < height-1 && guard.pos.c > 0 && guard.pos.c < width-1 {
		rNext, cNext := guard.peek() // peek
		for data[rNext][cNext] == '#' {
			guard.dir = (guard.dir + 1) % len(moves) // turn
			rNext, cNext = guard.peek()              // peek

		}
		guard.pos.r, guard.pos.c = rNext, cNext // Step forward
		_, ok := visited[guard]
		if !ok {
			visited[guard] = struct{}{}
		} else {
			return true
		}
	}
	return false
}

func solvePart2(data [][]rune, width, height int) int {
	guard := findGuard(data, width, height)
	_, steps := solveSteps(data, width, height, guard)

	bad := 0
	for pos := range steps {
		if pos == guard.pos {
			continue
		}
		data[pos.r][pos.c] = '#'
		loop := isLoop(data, width, height, guard)
		data[pos.r][pos.c] = '.'
		if loop {
			bad++
		}
	}
	return bad
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
