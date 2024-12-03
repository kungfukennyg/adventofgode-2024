package day3

import (
	"regexp"
	"strconv"
	"strings"
)

// match instructions matching `mul(0-999, 0-999)`
const multPattern = "mul\\((?:\\d{1,3})?,(?:\\d{1,3})?\\)"
const multConditionalPattern = "mul\\((?:\\d{1,3})?,(?:\\d{1,3})?\\)|do\\(\\)|don't\\(\\)"
const enable = "do()"
const disable = "don't()"

func scanCorrupted(input string, withConditionals bool) int {
	total := 0
	for _, product := range findProducts(input, withConditionals) {
		total += product
	}
	return total
}

func findProducts(input string, withConditionals bool) []int {
	var regex *regexp.Regexp
	if !withConditionals {
		regex = regexp.MustCompile(multPattern)
	} else {
		regex = regexp.MustCompile(multConditionalPattern)
	}

	products := []int{}
	enabled := true
	for _, instr := range regex.FindAllString(input, -1) {
		if withConditionals && (instr == enable || instr == disable) {
			enabled = instr == enable
			continue
		}

		if !enabled {
			continue
		}

		instr = strings.ReplaceAll(instr, "mul(", "")
		instr = strings.ReplaceAll(instr, ")", "")
		parts := strings.Split(instr, ",")
		a, _ := strconv.Atoi(parts[0])
		b, _ := strconv.Atoi(parts[1])

		products = append(products, a*b)
	}
	return products
}
