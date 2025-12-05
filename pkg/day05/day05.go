package day05

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type IngredientRange struct {
	Min int
	Max int
}

func Day05a() {
	ranges, ids := parseIngredients("input/ingredients.txt")

	freshCount := 0
	for _, id := range ids {
		if checkIngredient(id, &ranges) {
			freshCount++
		}
	}
	fmt.Println("Total invalid ingredients:", freshCount)
}

func Day05b() {
	ranges, _ := parseIngredients("input/ingredients.txt")

	outRanges := []IngredientRange{}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Min < ranges[j].Min
	})

	for _, rnge := range ranges {
		if len(outRanges) == 0 {
			outRanges = append(outRanges, rnge)
			continue
		}

		last := &outRanges[len(outRanges)-1]
		if rnge.Min <= last.Max+1 {
			if rnge.Max > last.Max {
				last.Max = rnge.Max
			}
		} else {
			outRanges = append(outRanges, rnge)
		}
	}

	ingredientCount := 0
	for _, rnge := range outRanges {
		ingredientCount += (rnge.Max - rnge.Min + 1)
	}

	fmt.Println("Total unique ingredients:", ingredientCount)
}

func parseIngredients(filename string) ([]IngredientRange, []int) {
	firstHalf := true

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var ranges []IngredientRange
	var ids []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			firstHalf = false
			continue
		}
		if firstHalf {
			line := scanner.Text()
			var min, max int
			fmt.Sscanf(line, "%d-%d", &min, &max)
			ranges = append(ranges, IngredientRange{Min: min, Max: max})
		} else {
			line := scanner.Text()
			var id int
			fmt.Sscanf(line, "%d", &id)
			ids = append(ids, id)
		}
	}

	return ranges, ids
}

func checkIngredient(id int, ranges *[]IngredientRange) bool {
	for _, r := range *ranges {
		if id >= r.Min && id <= r.Max {
			return true
		}
	}
	return false
}
