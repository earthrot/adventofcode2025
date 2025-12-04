package day03

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Day03a() {
	batteries, err := LoadBatteries("input/batteries.txt")
	if err != nil {
		panic(err)
	}

	totalCapacity := 0
	for _, battery := range batteries {
		capacity := parseBattery(battery, 2)
		totalCapacity += capacity
	}

	fmt.Println("Total battery capacity:", totalCapacity)
}

func Day03b() {
	batteries, err := LoadBatteries("input/batteries.txt")
	if err != nil {
		panic(err)
	}

	totalCapacity := 0
	for _, battery := range batteries {
		capacity := parseBattery(battery, 12)
		totalCapacity += capacity
	}

	fmt.Println("Total battery capacity:", totalCapacity)
}

func parseBattery(battery string, numdigits int) int {
	output := ""
	lastPos := -1

	for x := 0; x < numdigits; x++ {
		highest := 0
		highestStr := ""
		for i := lastPos + 1; i < len(battery)-(numdigits-1-x); i++ {
			intVal := int(battery[i] - '0')
			strVal := string(battery[i])
			if intVal > highest {
				highest = intVal
				highestStr = strVal
				lastPos = i
			}
		}

		fmt.Println("Found highest digit:", highestStr, "at position", lastPos)
		output += highestStr
	}

	retval, err := strconv.Atoi(output)
	if err != nil {
		panic(err)
	}

	return retval
}

func LoadBatteries(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, scanner.Err()
}
