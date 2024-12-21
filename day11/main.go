// day11.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([]int64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	stones := []int64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, stone := range strings.Split(line, " ") {
			stoneNumber, err := strconv.ParseInt(stone, 10, 64)
			if err != nil {
				return nil, err
			}
			stones = append(stones, stoneNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return stones, nil
}

type StoneIndex map[int64]int64

func (index StoneIndex) merge(other StoneIndex, quant int64) {
	for key := range other {
		index[key] += other[key] * quant
	}
}

func (index StoneIndex) populate(stones []int64) {
	for _, v := range stones {
		index[v] += 1
	}
}

var powersOf10 = []int64{1, 10, 100, 1_000, 10_000, 100_000, 1_000_000, 10_000_000, 100_000_000, 1_000_000_000}

func splitInteger(num int64) (left, right int64, digits int) {
	temp := num
	for temp > 0 {
		digits++
		temp /= 10
	}
	divisor := int64(powersOf10[digits/2])
	left = num / divisor
	right = num % divisor
	return
}

func blink(in []int64) []int64 {
	result := make([]int64, 0, len(in)*2)
	for _, s := range in {
		if s == 0 {
			result = append(result, 1)
		} else if left, right, digits := splitInteger(s); digits%2 == 0 {
			result = append(result, left, right)
		} else {
			result = append(result, s*2024)
		}
	}
	return result
}

func sumValues(index StoneIndex) (sum int64) {
	for _, value := range index {
		sum += value
	}
	return
}

func solve(stones []int64, quantity int64) (result int64) {
	index := make(StoneIndex)
	index.populate(stones)
	iterCount := []int64{5, quantity / 5}
	for range iterCount[0] {
		nextIndex := make(StoneIndex)
		for stone, quant := range index {
			currStones := []int64{stone}
			for range iterCount[1] {
				currStones = blink(currStones)
			}
			currIndex := make(StoneIndex)
			currIndex.populate(currStones)
			nextIndex.merge(currIndex, quant)
		}
		index = nextIndex
		result = sumValues(nextIndex)
	}
	return
}

func main() {
	stones, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	fmt.Println("Part 1:", solve(stones, 25))
	fmt.Println("Part 2:", solve(stones, 75))
}
