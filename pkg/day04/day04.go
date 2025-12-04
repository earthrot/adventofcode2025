package day04

import (
	"bufio"
	"fmt"
	"os"
)

func Day04a() {
	rolls, err := loadRolls("input/rolls.txt")
	if err != nil {
		panic(err)
	}

	count, _ := parseRolls(rolls)
	fmt.Println("Total valid rolls:", count)
}

func Day04b() {
	rolls, err := loadRolls("input/rolls.txt")
	if err != nil {
		panic(err)
	}

	totalCount := 0
	keepGoing := true
	inRolls := rolls
	for keepGoing {
		var count int
		count, inRolls = parseRolls(inRolls)
		totalCount += count
		if count == 0 {
			keepGoing = false
		}
	}

	fmt.Println("Total removed rolls:", totalCount)
}

func parseRolls(rolls [][]bool) (int, [][]bool) {
	outRolls := rolls
	count := 0
	for x := 0; x < len(rolls[0]); x++ {
		for y := 0; y < len(rolls); y++ {
			if rolls[x][y] {
				fmt.Println("Checking position:", x, y)
				if checkPosition(rolls, x, y) {
					count++
					outRolls[x][y] = false
				}
			}
		}
	}
	return count, outRolls
}

func checkPosition(rolls [][]bool, check_x int, check_y int) bool {
	totalCount := 0
	startX := max(0, check_x-1)
	endX := min(check_x+1, len(rolls[0])-1)
	startY := max(0, check_y-1)
	endY := min(check_y+1, len(rolls)-1)

	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if x == check_x && y == check_y {
				continue
			}
			if rolls[x][y] {
				totalCount++
			}
		}
	}
	fmt.Printf("%d,%d '@' count: %d\n", check_x, check_y, totalCount)
	return totalCount < 4
}

func loadRolls(filename string) ([][]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, parseLine(scanner.Text()))
	}

	return lines, scanner.Err()
}

func parseLine(line string) []bool {
	var result []bool
	for _, char := range line {
		if char == '@' {
			result = append(result, true)
		} else {
			result = append(result, false)
		}
	}
	return result
}
