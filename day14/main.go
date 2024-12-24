// day14.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const (
	height = 103
	width  = 101
)

type Coord2D struct{ x, y int }

type Robot struct{ id, x, y, vx, vy int }

func (r *Robot) move() {
	r.x = (height + r.x + r.vx) % height
	r.y = (width + r.y + r.vy) % width
}

func parseFile(filename string) (robots []Robot, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	robotPattern := regexp.MustCompile(`p=(\d+),(\d+) v=(-{0,1}\d+),(-{0,1}\d+)`)

	scanner := bufio.NewScanner(file)
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		if matches := robotPattern.FindStringSubmatch(line); matches != nil {
			y, _ := strconv.Atoi(matches[1])
			x, _ := strconv.Atoi(matches[2])
			vy, _ := strconv.Atoi(matches[3])
			vx, _ := strconv.Atoi(matches[4])
			robots = append(robots, Robot{id, x, y, vx, vy})
			id++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return robots, nil
}

func printRobots(robots []Robot) {
	grid := make([][]int, height)
	for i := range grid {
		grid[i] = make([]int, width)
	}
	for _, robot := range robots {
		grid[robot.x][robot.y]++
	}
	for _, row := range grid {
		for _, val := range row {
			if val != 0 {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\r\n")
	}
}

func solvePart1(inputRobots []Robot, steps int) int {
	robots := make([]Robot, len(inputRobots))
	copy(robots, inputRobots)

	// move
	for range steps {
		for r := range robots {
			robots[r].move()
		}
	}

	// quadrant count
	var q00, q01, q10, q11 int
	for _, robot := range robots {
		if robot.x < height/2 {
			if robot.y < width/2 {
				q00++
			} else if robot.y > width/2 {
				q01++
			}
		} else if robot.x > height/2 {
			if robot.y < width/2 {
				q10++
			} else if robot.y > width/2 {
				q11++
			}
		}
	}
	return q00 * q01 * q10 * q11
}

func solvePart2(inputRobots []Robot) int {
	robots := make([]Robot, len(inputRobots))
	copy(robots, inputRobots)

	cache := make(map[Coord2D]struct{})
	
	var step int
	for ; len(cache) != len(robots); step++ {
		clear(cache)
		for r := range robots {
			robots[r].move()
			cache[Coord2D{robots[r].x, robots[r].y}] = struct{}{}
		}
	}
	// printRobots(robots)
	return step
}

func main() {
	robots, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("Part 1:", solvePart1(robots, 100))
	fmt.Println("Part 2:", solvePart2(robots))
}
