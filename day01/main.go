// day01.go
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readArrays(filename string, numColumns int) ([][]int, error) {
	// Open and read the file
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %v", err)
	}

	// Initialize a slice to hold arrays for each column
	columns := make([][]int, numColumns)
	for i := range columns {
		columns[i] = []int{}
	}

	// Split the file into lines
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, line := range lines {
		// Split each line into parts
		parts := strings.Fields(line)
		if len(parts) != numColumns {
			return nil, fmt.Errorf("invalid line format: %q (expected %d columns)", line, numColumns)
		}

		// Convert each part to an integer and append it to the corresponding column
		for i := 0; i < numColumns; i++ {
			num, err := strconv.Atoi(parts[i])
			if err != nil {
				return nil, fmt.Errorf("error converting number %q on line: %q", parts[i], line)
			}
			columns[i] = append(columns[i], num)
		}
	}
	return columns, nil
}

// Helper function to calculate the absolute value of a number
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func solvePart1(arr1, arr2 []int) int {
	// Sort both slices
	sort.Ints(arr1)
	sort.Ints(arr2)
	// Calculate the absolute differences and sum them
	acc := 0
	for i := 0; i < len(arr1) && i < len(arr2); i++ {
		acc += abs(arr1[i] - arr2[i])
	}
	return acc
}

func countFrequencyBinary(source []int, target []int) []int {
	result := make([]int, len(source))

	for i, num := range source {
		// Find first occurrence
		left := findFirst(target, num)
		if left == -1 {
			result[i] = 0
			continue
		}

		// Find last occurrence
		right := findLast(target, num)
		result[i] = right - left + 1
	}

	return result
}

func findFirst(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if (mid == 0 || arr[mid-1] < target) && arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func findLast(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if (mid == len(arr)-1 || arr[mid+1] > target) && arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func solvePart2(arr1, arr2 []int) int {
	sort.Ints(arr2)
	res := countFrequencyBinary(arr1, arr2)
	acc := 0
	for i := 0; i < len(arr1) && i < len(res); i++ {
		acc += arr1[i] * res[i]
	}
	return acc
}

func main() {
	// Read the input arrays
	columns, err := readArrays("input.txt", 2)
	if err != nil {
		log.Fatal(err)
	}
	arr1 := columns[0]
	arr2 := columns[1]

	fmt.Println("Part 1:", solvePart1(arr1, arr2))
	fmt.Println("Part 2:", solvePart2(arr1, arr2))
}
