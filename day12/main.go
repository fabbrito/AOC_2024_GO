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

var directions = [4][2]int{
	{-1, 0}, // Up
	{0, 1},  // Right
	{1, 0},  // Down
	{0, -1}, // Left
}

func bfs(grid [][]rune, height, width int, visited [][]bool, startX, startY int, char rune) (int, int) {
	queue := [][2]int{{startX, startY}}
	visited[startX][startY] = true
	area, perimeter := 0, 0
	for len(queue) > 0 {
		// Dequeue the front element
		x, y := queue[0][0], queue[0][1]
		queue = queue[1:]
		area++ // Increment area for each cell in the region
		// Check all 4 neighbors
		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			// If neighbor is out of bounds or a different character, it's part of the perimeter
			if nx < 0 || nx >= height || ny < 0 || ny >= width || grid[nx][ny] != char {
				perimeter++
			} else if !visited[nx][ny] { // If valid and not visited, enqueue it
				visited[nx][ny] = true
				queue = append(queue, [...]int{nx, ny})
			}
		}
	}
	return area, perimeter
}

func solvePart1(grid [][]rune, height, width int) int {
	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}
	result := 0
	// Process each cell and calculate regions
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if !visited[x][y] {
				char := grid[x][y]
				area, perimeter := bfs(grid, height, width, visited, x, y, char)
				result += area * perimeter
				// fmt.Printf("Region with char '%c': Area = %d, Perimeter = %d\n", char, area, perimeter)
			}
		}
	}

	return result
}

func solvePart2(data [][]rune, height, width int) int {
	result := 0

	return result
}

func main() {
	grid, height, width, err := parseFile("sample0.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// for _, row := range grid {
	// 	fmt.Println(string(row))
	// }

	fmt.Println("Part 1:", solvePart1(grid, height, width))
	fmt.Println("Part 2:", solvePart2(grid, height, width))
}
