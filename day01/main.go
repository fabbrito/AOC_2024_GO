// day01.go
package main

import (
	"aoc2024/utils"
	"fmt"
	"log"
	"sort"
)

func solvePart1(arr1, arr2 []int) int {
	// Sort both slices
	sort.Ints(arr1)
	sort.Ints(arr2)
	// Calculate the absolute differences and sum them
	acc := 0
	for i := 0; i < len(arr1) && i < len(arr2); i++ {
		acc += utils.Abs(arr1[i] - arr2[i])
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
	columns, err := utils.ReadArrays("input.txt", 2)
	if err != nil {
		log.Fatal(err)
	}
	arr1 := columns[0]
	arr2 := columns[1]

	fmt.Println("Part 1:", solvePart1(arr1, arr2))
	fmt.Println("Part 2:", solvePart2(arr1, arr2))
}
