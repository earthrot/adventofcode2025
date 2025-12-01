package day01

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var startPosition int = 50
var maxSteps int = 100
var stepParse *regexp.Regexp

func init() {
	stepParse = regexp.MustCompile(`([LR])(\d+)`)
}

func Day01a() {
	fmt.Println("Day 1a")

	rotations, err := LoadRotations("input/rotations.txt")
	if err != nil {
		panic(err)
	}

	value := startPosition
	zeroCount := 0

	for _, rotation := range rotations {
		step := parseRotation(rotation)
		step = step % maxSteps
		value += step
		if value < 0 {
			value += maxSteps
		}

		if value > maxSteps {
			value -= maxSteps
		}

		if value == 0 || value == 100 {
			zeroCount++
		}
		fmt.Printf("Rotation: %d, value: %d\n", step, value)
	}

	fmt.Println("Final Value:", zeroCount)
}

func Day01b() {
	fmt.Println("Day 1b")

	rotations, err := LoadRotations("input/rotations.txt")
	if err != nil {
		panic(err)
	}

	value := startPosition
	zeroCount := 0

	for _, rotation := range rotations {
		newValue := value
		step := parseRotation(rotation)

		fullRound := step / maxSteps
		fullRound = max(fullRound, -fullRound)
		zeroCount += fullRound

		moddedStep := step % maxSteps
		newValue += moddedStep

		if newValue == maxSteps {
			newValue = 0
		}

		if newValue < 0 {
			if newValue-moddedStep != 0 {
				zeroCount++
			}
			newValue += maxSteps
		}

		if newValue > maxSteps {
			if newValue-moddedStep != 0 {
				zeroCount++
			}
			newValue -= maxSteps
		}

		if newValue == 0 {
			zeroCount++
		}

		fmt.Printf("value: %d \tstep: %d \tmodstep: %d \tfullround: %d \tnewvalue: %d \tzeroes:%d\n", value, step, moddedStep, fullRound, newValue, zeroCount)
		value = newValue
	}

	fmt.Println("Final Value:", zeroCount)
}

func LoadRotations(filename string) ([]string, error) {
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

func parseRotation(step string) int {
	matches := stepParse.FindStringSubmatch(step)
	if len(matches) != 3 {
		panic("invalid step: " + step)
	}

	val, err := strconv.Atoi(matches[2])
	if err != nil {
		panic(err)
	}

	if matches[1] == "R" {
		return val
	}
	return -val
}
