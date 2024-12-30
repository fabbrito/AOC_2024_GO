// day18.go
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Point struct{ y, x int }

func parseFile(filename string) (data []Point, err error) {
	file, errFile := os.Open(filename)
	if errFile != nil {
		return nil, errFile
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			xyStr := strings.Split(line, ",")
			x, _ := strconv.Atoi(xyStr[0])
			y, _ := strconv.Atoi(xyStr[1])
			data = append(data, Point{x: x, y: y})
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return
}

func generateGrid(data []Point, height, width, step int) (grid [][]bool) {
	grid = make([][]bool, height)
	for y := 0; y < height; y++ {
		grid[y] = make([]bool, width)
	}
	step = min(step, len(data)-1)
	for _, pos := range data[:step] {
		grid[pos.y][pos.x] = true
	}
	return
}

func printGrid(grid [][]bool) {
	height, width := len(grid), len(grid[0])
	for i := 0; i < height; i++ {
		var line string
		for j := 0; j < width; j++ {
			if grid[i][j] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

// Define the directions as a map of Directions to Points
var DIRECTIONS = map[Direction]struct{ dy, dx int }{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

type Item struct {
	y, x int
	dir  Direction
	cost int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].cost < pq[j].cost // Min-heap based on cost
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // Avoid memory leak
	*pq = old[:n-1]
	return item
}

type State struct {
	y, x int
	dir  Direction
}

func solve(grid [][]bool, height, width, startX, startY, endX, endY int) int {
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{y: startY, x: startX, dir: DOWN, cost: 0})
	costs := make(map[State]int)
	costs[State{y: startY, x: startX, dir: DOWN}] = 0
	// Dijkstra's algorithm
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)
		// Goal check
		if current.x == endX && current.y == endY {
			return current.cost
		}
		// Explore all possible moves
		for nextDir, delta := range DIRECTIONS {
			ny, nx := current.y+delta.dy, current.x+delta.dx
			nextCost := current.cost + 1
			nextState := State{y: ny, x: nx, dir: nextDir}

			if ny >= 0 && nx >= 0 && ny < height && nx < width && !grid[ny][nx] {
				if oldCost, seen := costs[nextState]; !seen || (nextCost < oldCost) {
					costs[nextState] = nextCost
					heap.Push(pq, &Item{y: ny, x: nx, dir: nextDir, cost: nextCost})
				}
			}
		}
	}
	return -1
}

func solvePart1(data []Point, height, width, step int) int {
	startTime := time.Now()
	defer func(){
		fmt.Printf("Part 1 execution took %s\n", time.Since(startTime))
	}()

	grid := generateGrid(data, height, width, step)
	return solve(grid, height, width, 0, 0, width-1, height-1)
}

func solvePart2(data []Point, height, width, minStep int) string {
	startTime := time.Now()
	defer func(){
		fmt.Printf("Part 2 execution took %s\n", time.Since(startTime))
	}()

	step := len(data) - 1
	cost := -1
	for cost == -1 && step >= minStep {
		grid := generateGrid(data, height, width, step)
		cost = solve(grid, height, width, 0, 0, width-1, height-1)
		step--
	}
	step++
	return fmt.Sprintf("%d,%d", data[step].x, data[step].y)
}

func main() {
	height, width, step := 71, 71, 1024
	data, err := parseFile("input.txt")
	// height, width, step := 7, 7, 12
	// data, err := parseFile("sample0.txt")
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Part 1: ", solvePart1(data, height, width, step))
	fmt.Println("Part 2: ", solvePart2(data, height, width, step))
}
