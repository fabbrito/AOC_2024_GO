// day13.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFile(filename string) (buttonAs, buttonBs, prizes [][2]int, quant int, err error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, nil, -1, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		if count%4 == 0 {
			lineA := scanner.Text()
			rightA := strings.Split(lineA, "X+")[1]
			xyA := strings.Split(rightA, ", Y+")
			xA, _ := strconv.Atoi(xyA[0])
			yA, _ := strconv.Atoi(xyA[1])
			buttonAs = append(buttonAs, [2]int{xA, yA})
		} else if count%4 == 1 {
			lineB := scanner.Text()
			rightB := strings.Split(lineB, "X+")[1]
			xyB := strings.Split(rightB, ", Y+")
			xB, _ := strconv.Atoi(xyB[0])
			yB, _ := strconv.Atoi(xyB[1])
			buttonBs = append(buttonBs, [2]int{xB, yB})
		} else if count%4 == 2 {
			linePrize := scanner.Text()
			rightPrize := strings.Split(linePrize, "X=")[1]
			xyPrize := strings.Split(rightPrize, ", Y=")
			xPrize, _ := strconv.Atoi(xyPrize[0])
			yPrize, _ := strconv.Atoi(xyPrize[1])
			prizes = append(prizes, [2]int{xPrize, yPrize})
		}
		count++
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, nil, -1, err
	}
	quant = len(buttonAs)
	if len(buttonAs) != len(buttonBs) || len(buttonAs) != len(prizes) {
		return nil, nil, nil, -1, fmt.Errorf("Error while parsing the file. Output slices should not have different lengths!")
	}

	return
}

func solvePart1(buttonAs, buttonBs, prizes [][2]int, quantMachines int) int {
	total := 0
	for i := 0; i < quantMachines; i++ {
		D := buttonAs[i][0]*buttonBs[i][1] - buttonAs[i][1]*buttonBs[i][0]
		Dx := prizes[i][0]*buttonBs[i][1] - prizes[i][1]*buttonBs[i][0]
		Dy := buttonAs[i][0]*prizes[i][1] - buttonAs[i][1]*prizes[i][0]
		if D == 0 {
			panic("Linear equation determinant cannot be zero!")
		}
		if Dx%D == 0 && Dy%D == 0 {
			a, b := Dx/D, Dy/D
			if a >= 0 && a <= 100 && b >= 0 && b <= 100 {
				total += 3*a + b
			}
		}
	}
	return total
}

func solvePart2(buttonAs, buttonBs, prizes [][2]int, quantMachines int) int {
	for i := 0; i < quantMachines; i++ {
		prizes[i][0] = prizes[i][0] + 10_000_000_000_000
		prizes[i][1] = prizes[i][1] + 10_000_000_000_000
	}
	total := 0
	for i := 0; i < quantMachines; i++ {
		D := buttonAs[i][0]*buttonBs[i][1] - buttonAs[i][1]*buttonBs[i][0]
		Dx := prizes[i][0]*buttonBs[i][1] - prizes[i][1]*buttonBs[i][0]
		Dy := buttonAs[i][0]*prizes[i][1] - buttonAs[i][1]*prizes[i][0]
		if D == 0 {
			panic("Linear equation determinant cannot be zero!")
		}
		if Dx%D == 0 && Dy%D == 0 {
			a, b := Dx/D, Dy/D
			total += 3*a + b
		}
	}
	return total
}

func main() {
	buttonAs, buttonBs, prizes, quantMachines, err := parseFile("input.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	fmt.Println("Part 1:", solvePart1(buttonAs, buttonBs, prizes, quantMachines))
	fmt.Println("Part 2:", solvePart2(buttonAs, buttonBs, prizes, quantMachines))
}
