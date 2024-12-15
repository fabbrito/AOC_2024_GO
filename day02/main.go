// day02.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) ([][]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var result [][]int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var lineInts []int
		numbers := strings.Fields(scanner.Text())

		for _, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				continue // Skip invalid numbers
			}
			lineInts = append(lineInts, n)
		}
		result = append(result, lineInts)
	}

	return result, scanner.Err()
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
func isValid(slice []int) bool {
	if (len(slice) <= 2) && (slice[0] != slice[1]) {
		return true
	}
	increasing := slice[1] > slice[0]
	for i := 1; i < len(slice); i++ {
		diff := abs(slice[i] - slice[i-1])
		if (diff < 1) || (diff > 3) || (increasing && slice[i] <= slice[i-1]) || (!increasing && slice[i] >= slice[i-1]) {
			return false
		}
	}
	return true
}

func solvePart1(slices [][]int) int {
	countValid := 0
	for _, slice := range slices {
		if isValid(slice) {
			countValid++
		}
	}
	return countValid
}

func solvePart2(slices [][]int) int {
	countValid := 0
	for _, slice := range slices {
		if isValid(slice) {
			countValid++
		} else {
			for i := 0; i < len(slice); i++ {
				var modSlice []int
				modSlice = append(modSlice, slice[:i]...)
				modSlice = append(modSlice, slice[i+1:]...)
				if isValid(modSlice) {
					countValid++
					break
				}
			}
		}
	}
	return countValid
}

func main() {
	slices, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(slices))
	fmt.Println("Part 2:", solvePart2(slices))
}
