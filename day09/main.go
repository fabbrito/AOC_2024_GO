// day08.go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			data = append(data, int(char)-int('0'))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func solvePart1(data []int) int {
	res := 0
	pos := 0
	pLast := len(data) - 1

	leftover := data[pLast]

	totalSize := 0
	for i := 0; i < len(data); i += 2 {
		totalSize += data[i]
	}

	for p := 0; pos < totalSize && p < len(data); p++ {
		if p%2 == 0 { // normal file
			for i := 0; pos < totalSize && i < data[p]; i++ {
				res += (p / 2) * pos
				pos++
			}
		} else {
			// runs until no space left
			for i := 0; pos < totalSize && i < data[p]; i++ {
				if leftover == 0 {
					pLast -= 2
					leftover = data[pLast]
				}
				res += (pLast / 2) * pos
				pos++
				leftover--
			}
		}
	}
	return res
}

type File struct {
	id, startPosition, size int
	empty                   bool
}

type FS []File

func solvePart2(data []int) int {
	fs := FS{}
	startPosition := 0
	for i := 0; i <= len(data)-1; i++ {
		fs = append(fs, File{i / 2, startPosition, data[i], i%2 != 0})
		startPosition += data[i]
	}
	for pLast := len(fs) - 1; pLast > 0; pLast -= 2 {
		for pSpace := 1; pSpace < pLast; pSpace += 2 {
			if fs[pSpace].empty && (fs[pSpace].size >= fs[pLast].size) {
				fs[pSpace].size -= fs[pLast].size
				fs[pLast].startPosition = fs[pSpace].startPosition
				fs[pSpace].startPosition += fs[pLast].size
				break
			}
		}
	}
	checksum := 0
	for i, file := range fs {
		if i%2 != 0 {
			continue
		}
		checksum += file.id * (file.size * (file.startPosition + file.startPosition + file.size - 1) / 2)
	}
	return checksum
}

func main() {
	data, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(data))
	fmt.Println("Part 2:", solvePart2(data))
}
