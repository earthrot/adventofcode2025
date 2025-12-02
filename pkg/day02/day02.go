package day02

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Day02a() {
	ranges, err := LoadRanges("input/prodid.txt")
	if err != nil {
		panic(err)
	}

	invalidSum := 0

	for _, r := range ranges {
		start, stop, err := parseRange(r)
		if err != nil {
			fmt.Printf("Error parsing range %s: %v\n", r, err)
			continue
		}

		for i := start; i <= stop; i++ {
			code := strconv.Itoa(i)
			if !validateProductCodeA(code) {
				fmt.Printf("Invalid product code: %s\n", code)
				invalidSum += i
			}
		}
	}

	fmt.Println("Sum of invalid product codes:", invalidSum)
}

func Day02b() {
	ranges, err := LoadRanges("input/prodid.txt")
	if err != nil {
		panic(err)
	}

	invalidSum := 0

	for _, r := range ranges {
		start, stop, err := parseRange(r)
		if err != nil {
			fmt.Printf("Error parsing range %s: %v\n", r, err)
			continue
		}

		for i := start; i <= stop; i++ {
			code := strconv.Itoa(i)
			if !validateProductCodeB(code) {
				fmt.Printf("Invalid product code: %s\n", code)
				invalidSum += i
			}
		}
	}

	fmt.Println("Sum of invalid product codes:", invalidSum)
}

func parseRange(r string) (int, int, error) {
	parts := strings.Split(r, "-")
	if len(parts) != 2 {
		return 0, 0, errors.New("invalid range format")
	}
	min, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, err
	}
	max, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, err
	}
	return min, max, nil
}

func validateProductCodeA(code string) bool {
	idLen := len(code)
	if idLen%2 != 0 {
		return true
	}

	return code[0:idLen/2] != code[idLen/2:idLen]
}

func validateProductCodeB(code string) bool {
	for i := 0; i < len(code)/2; i++ {
		re := regexp.MustCompile(fmt.Sprintf(`(?m)^(%s)+$`, code[0:i+1]))
		if re.MatchString(code) {
			return false
		}
	}

	return true
}

func LoadRanges(filename string) ([]string, error) {
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

	if len(lines) > 0 {
		return strings.Split(lines[0], ","), nil
	}

	return nil, errors.New("no ranges found in file")
}
