// day16.go
package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
)

func parseFile(filename string) (mazeData [][]rune, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] == '#' {
			mazeData = append(mazeData, []rune(line))
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return mazeData, nil
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

// Define the directions as a map of Directions to Points
var DIRECTIONS = map[Direction]struct{ dx, dy int }{
	NORTH: {-1, 0},
	EAST:  {0, 1},
	SOUTH: {1, 0},
	WEST:  {0, -1},
}

func generateMaze(data [][]rune) (maze [][]bool, startX, startY, endX, endY int) {
	height, width := len(data), len(data[0])
	maze = make([][]bool, height)
	for x := 0; x < height; x++ {
		maze[x] = make([]bool, width)
		for y := 0; y < width; y++ {
			if data[x][y] == '#' {
				maze[x][y] = true
			} else if data[x][y] == 'S' {
				startX, startY = x, y
			} else if data[x][y] == 'E' {
				endX, endY = x, y
			}
		}
	}
	return
}

func printMaze(maze [][]bool) {
	height, width := len(maze), len(maze[0])
	for i := 0; i < height; i++ {
		line := ""
		for j := 0; j < width; j++ {
			if maze[i][j] {
				line += "#"
			} else {
				line += "."
			}
		}
		fmt.Println(line)
	}
}

func printMazePath(data [][]rune, path [][2]int) {
	height, width := len(data), len(data[0])
	maze := make([][]rune, height)
	for x := 0; x < height; x++ {
		maze[x] = make([]rune, width)
		copy(maze[x], data[x])
	}
	for _, pos := range path {
		maze[pos[0]][pos[1]] = 'O'
	}
	for _, row := range maze {
		fmt.Println(string(row))
	}
}

type Item struct {
	x, y         int
	dir          Direction
	cost         int
	predecessors []*Item
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
	x, y int
	dir  Direction
}

func solvePart1(data [][]rune) int {
	maze, startX, startY, endX, endY := generateMaze(data)
	height, width := len(maze), len(maze[0])

	// Priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{startX, startY, EAST, 0, nil})

	costs := make(map[State]int)
	costs[State{startX, startY, 0}] = 0

	// Dijkstra's algorithm
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)

		// Goal check
		if current.x == endX && current.y == endY {
			return current.cost
		}
		// Explore all possible moves
		for nextDir, delta := range DIRECTIONS {
			nx, ny := current.x+delta.dx, current.y+delta.dy
			turnCost := 0
			if nextDir != current.dir {
				turnCost = 1000 // Cost of turning
			}
			stepCost := 1 // Cost of moving forward
			totalCost := current.cost + turnCost + stepCost

			nextState := State{nx, ny, nextDir}
			if nx >= 0 && ny >= 0 && nx < height && ny < width && !maze[nx][ny] {
				if oldCost, seen := costs[nextState]; !seen || (totalCost < oldCost) {
					costs[nextState] = totalCost
					heap.Push(pq, &Item{nx, ny, nextDir, totalCost, nil})
				}
			}
		}
	}
	return -1
}

func reconstructPaths(current *Item, path [][2]int, allPaths *[][][2]int) {
	// Add the current position to the path
	path = append([][2]int{{current.x, current.y}}, path...)
	// If there are no predecessors, we've reached the starting state
	if len(current.predecessors) == 0 {
		*allPaths = append(*allPaths, append([][2]int{}, path...)) // Add a copy of the path
		return
	}
	// Recurse for each predecessor
	for _, predecessor := range current.predecessors {
		reconstructPaths(predecessor, path, allPaths)
	}
}

func solvePart2(data [][]rune) int {
	maze, startX, startY, endX, endY := generateMaze(data)
	height, width := len(maze), len(maze[0])

	// Priority queue
	pq := &PriorityQueue{}
	heap.Init(pq)
	start := &Item{startX, startY, EAST, 0, nil}
	heap.Push(pq, start)

	costs := make(map[State]int)
	states := make(map[State]*Item)

	states[State{startX, startY, 0}] = start
	costs[State{startX, startY, 0}] = 0
	minGoalCost := -1 // Track the minimum cost to reach the goal
	var goalStates []*Item

	// Dijkstra's algorithm
	for pq.Len() > 0 {
		current := heap.Pop(pq).(*Item)

		// Goal check
		if current.x == endX && current.y == endY {
			if minGoalCost == -1 || current.cost < minGoalCost {
				minGoalCost = current.cost
				goalStates = []*Item{current} // Reset with the new minimum cost
			} else if current.cost == minGoalCost {
				goalStates = append(goalStates, current)
			}
			continue
		}
		// Explore all possible moves
		for nextDir, delta := range DIRECTIONS {
			nx, ny := current.x+delta.dx, current.y+delta.dy
			turnCost := 0
			if nextDir != current.dir {
				turnCost = 1000 // Cost of turning
			}
			stepCost := 1 // Cost of moving forward
			totalCost := current.cost + turnCost + stepCost

			nextState := State{nx, ny, nextDir}
			if nx >= 0 && ny >= 0 && nx < height && ny < width && !maze[nx][ny] {
				if oldCost, seen := costs[nextState]; !seen || (totalCost < oldCost) {
					costs[nextState] = totalCost
					next := &Item{nx, ny, nextDir, totalCost, []*Item{current}}
					states[nextState] = next
					heap.Push(pq, next)
				} else if totalCost == oldCost {
					states[nextState].predecessors = append(states[nextState].predecessors, current)
				}
			}
		}
	}
	// Reconstruct the path
	mapPath := make(map[[2]int]bool)
	// If no goal states are found
	if len(goalStates) == 0 {
		return 0
	}
	// // copy the maze to lay the paths on top
	// mazeData := make([][]rune, height)
	// for x := 0; x < height; x++ {
	// 	mazeData[x] = make([]rune, width)
	// 	copy(mazeData[x], data[x])
	// }

	// Reconstruct all paths
	var allPaths [][][2]int
	for _, goal := range goalStates {
		reconstructPaths(goal, [][2]int{}, &allPaths)
	}
	for _, path := range allPaths {
		// // override paths on top of the maze
		// for _, pos := range path {
		// 	mazeData[pos[0]][pos[1]] = 'O'
		// }
		for _, pos := range path {
			if _, seen := mapPath[pos]; !seen {
				mapPath[pos] = true
			}
		}
	}
	// // print the maze
	// for _, row := range mazeData {
	// 	fmt.Println(string(row))
	// }
	return len(mapPath)
}

func main() {
	data, err := parseFile("input.txt")

	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Part 1: ", solvePart1(data))
	fmt.Println("Part 2: ", solvePart2(data))
}
