// day12.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) (grid [][]rune, height, width int, err error) {
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
		gridLine := make([]rune, 0, width)
		for _, char := range line {
			gridLine = append(gridLine, char)
		}
		grid = append(grid, gridLine)
		height++
	}

	err = scanner.Err()
	if err != nil {
		return
	}

	return
}

type Coord2D struct{ x, y int }

func (c Coord2D) inBounds(height, width int) bool {
	return c.x >= 0 && c.x < height && c.y >= 0 && c.y < width
}

var deltas = [4]Coord2D{
	{-1, 0}, // NORTH
	{0, 1},  // EAST
	{1, 0},  // SOUTH
	{0, -1}, // WEST
}

func bfs(grid [][]rune, height, width int, visited [][]bool, startX, startY int, region map[Coord2D]bool) (perimeter int) {
	queue := []Coord2D{{startX, startY}}
	for len(queue) > 0 {
		// Dequeue the front element
		curr := queue[0]
		queue = queue[1:]
		if !visited[curr.x][curr.y] {
			visited[curr.x][curr.y] = true
			region[curr] = true
			for _, delta := range deltas {
				next := Coord2D{curr.x + delta.x, curr.y + delta.y}
				if !next.inBounds(height, width) || grid[curr.x][curr.y] != grid[next.x][next.y] {
					perimeter++
				} else if !visited[next.x][next.y] {
					queue = append(queue, next)
				}
			}
		}
	}
	return perimeter
}

func solvePart1(grid [][]rune, height, width int) int {
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}
	result := 0
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if !visited[x][y] {
				region := make(map[Coord2D]bool)
				perimeter := bfs(grid, height, width, visited, x, y, region)
				area := len(region)
				result += area * perimeter
			}
		}
	}
	return result
}

func countEdges(region map[Coord2D]bool) (edges int) {
	for coord := range region {
		for i := range deltas {
			next := Coord2D{coord.x + deltas[i].x, coord.y + deltas[i].y}
			if !region[next] {
				iNeg90 := (len(deltas) + i - 1) % len(deltas)
				nextNeg90 := Coord2D{coord.x + deltas[iNeg90].x, coord.y + deltas[iNeg90].y}
				nextNeg45 := Coord2D{coord.x + deltas[i].x + deltas[iNeg90].x, coord.y + deltas[i].y + deltas[iNeg90].y}
				if !region[nextNeg90] || region[nextNeg45] {
					edges++
				}
			}
		}
	}
	return edges
}

func solvePart2(grid [][]rune, height, width int) int {
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}
	result := 0
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if !visited[x][y] {
				region := make(map[Coord2D]bool)
				bfs(grid, height, width, visited, x, y, region)
				area := len(region)
				edges := countEdges(region)
				result += area * edges
			}
		}
	}
	return result
}

func main() {
	grid, height, width, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(grid, height, width))
	fmt.Println("Part 2:", solvePart2(grid, height, width))
}
