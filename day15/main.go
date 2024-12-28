// day15.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) (gridData [][]rune, instructions []rune, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[0] == '#' {
			gridData = append(gridData, []rune(line))
		} else {
			instructions = append(instructions, []rune(line)...)
		}
	}

	if err = scanner.Err(); err != nil {
		return
	}

	return gridData, instructions, nil
}

type Point struct {
	r int
	c int
}

// Define the directions as a map of characters to Points
var DIRECTIONS = map[rune]Point{
	'^': {r: -1, c: 0},
	'>': {r: 0, c: 1},
	'v': {r: 1, c: 0},
	'<': {r: 0, c: -1},
}

type Grid [][]rune

func (grid *Grid) init(data [][]rune) (height, width int) {
	*grid = make(Grid, len(data))
	for i, row := range data {
		(*grid)[i] = make([]rune, len(row))
		copy((*grid)[i], row)
		if width == 0 {
			width = len(row)
		}
		height++
	}
	return
}

func (grid *Grid) swap(position, next Point) {
	(*grid)[position.r][position.c], (*grid)[next.r][next.c] = (*grid)[next.r][next.c], (*grid)[position.r][position.c]
}

func (grid *Grid) moveBox(position Point, direction Point) bool {
	next := Point{position.r + direction.r, position.c + direction.c}

	if (*grid)[next.r][next.c] == '.' {
		// If the next spot is empty, swap positions
		(*grid).swap(position, next)
		return true
	} else if (*grid)[next.r][next.c] == '#' {
		// If the next spot is a wall, stop all from moving
		return false
	} else {
		// Only move the current box if the next box can move
		if (*grid).moveBox(next, direction) {
			(*grid).swap(position, next)
			return true
		}
	}
	return false // This should never be reached
}

func (grid *Grid) print(robot Point) {
	(*grid)[robot.r][robot.c] = '@'
	for _, row := range *grid {
		fmt.Println(string(row))
	}
	(*grid)[robot.r][robot.c] = '.'
}

func solvePart1(data [][]rune, instructions []rune) int {
	var grid Grid
	height, width := grid.init(data)

	// Find the robot and clear its position
	robot := Point{0, 0}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '@' {
				robot = Point{r, c}
				grid[r][c] = '.'
			}
		}
	}

	// Process each instruction
	for _, instruction := range instructions {
		direction := DIRECTIONS[instruction]
		position := Point{robot.r + direction.r, robot.c + direction.c}

		// If there is a wall, don't move
		if grid[position.r][position.c] != '#' {
			// If there is an empty spot, move without moving boxes
			if grid[position.r][position.c] == '.' {
				robot = position
			}
			// If there is a box, try to move all the boxes, then move
			if grid[position.r][position.c] == 'O' && grid.moveBox(position, direction) {
				robot = position
			}
		}
		// grid.print(robot)
	}

	// GPS
	score := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == 'O' {
				score += r*100 + c
			}
		}
	}
	return score
}

func (grid *Grid) wideInit(data [][]rune) (height, width int) {
	(*grid) = make([][]rune, len(data))
	for i, row := range data {
		(*grid)[i] = make([]rune, 2*len(row))
		for j, char := range row {
			switch char {
			case '@':
				(*grid)[i][2*j] = '@'
				(*grid)[i][2*j+1] = '.'
			case 'O':
				(*grid)[i][2*j] = '['
				(*grid)[i][2*j+1] = ']'
			default:
				(*grid)[i][2*j] = char
				(*grid)[i][2*j+1] = char
			}
		}
	}
	height, width = len(*grid), len((*grid)[0])
	return
}

func (grid *Grid) moveWideBox(position Point, direction Point) bool {
	if direction.r == 0 { // horizontal
		return (*grid).moveBox(position, direction)
	}

	linkedBoxes := [][2]Point{}
	seenBoxes := make(map[[2]Point]bool)

	checkAndAppend := func(p Point) {
		if (*grid)[p.r][p.c] == '[' {
			box := [2]Point{p, {p.r, p.c + 1}}
			if _, ok := seenBoxes[box]; !ok {
				seenBoxes[box] = true
				linkedBoxes = append(linkedBoxes, box)
			}
		} else if (*grid)[p.r][p.c] == ']' {
			box := [2]Point{{p.r, p.c - 1}, p}
			if _, ok := seenBoxes[box]; !ok {
				seenBoxes[box] = true
				linkedBoxes = append(linkedBoxes, box)
			}
		}
	}

	checkAndAppend(position)
	count := 0
	for ; count < len(linkedBoxes); count++ {
		left, right := linkedBoxes[count][0], linkedBoxes[count][1]
		nextLeft := Point{left.r + direction.r, left.c + direction.c}
		nextRight := Point{right.r + direction.r, right.c + direction.c}
		if (*grid)[nextLeft.r][nextLeft.c] == '#' || (*grid)[nextRight.r][nextRight.c] == '#' {
			return false
		}
		checkAndAppend(nextLeft)
		checkAndAppend(nextRight)
	}

	for i := count - 1; i >= 0; i-- {
		left, right := linkedBoxes[i][0], linkedBoxes[i][1]
		nextLeft := Point{left.r + direction.r, left.c + direction.c}
		nextRight := Point{right.r + direction.r, right.c + direction.c}
		(*grid).swap(left, nextLeft)
		(*grid).swap(right, nextRight)
	}
	return true
}

func solvePart2(data [][]rune, instructions []rune) int {
	var grid Grid
	height, width := grid.wideInit(data)

	// Find the robot and clear its position
	robot := Point{0, 0}
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '@' {
				robot = Point{r, c}
				grid[r][c] = '.'
			}
		}
	}
	// Process each instruction
	for _, instruction := range instructions {
		direction := DIRECTIONS[instruction]
		position := Point{robot.r + direction.r, robot.c + direction.c}

		// If there is a wall, don't move
		if grid[position.r][position.c] != '#' {
			// If there is an empty spot, move without moving boxes
			if grid[position.r][position.c] == '.' {
				robot = position
			}
			// If there is a box, try to move all the boxes, then move
			if (grid[position.r][position.c] == '[' || grid[position.r][position.c] == ']') && grid.moveWideBox(position, direction) {
				robot = position
			}
		}
		// fmt.Printf("%c\r\n", instruction)
		// grid.print(robot)
	}

	// GPS
	score := 0
	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			if grid[r][c] == '[' {
				score += r*100 + c
			}
		}
	}
	return score
}

func main() {
	data, instructions, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error parsing file: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Part 1: ", solvePart1(data, instructions))
	fmt.Println("Part 2: ", solvePart2(data, instructions))
}
