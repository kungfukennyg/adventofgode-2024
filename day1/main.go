package day1

import (
	"slices"
)

func similarityScore(left, right []int) int {
	unique := map[int]struct{}{}
	for _, l := range left {
		unique[l] = struct{}{}
	}

	counts := map[int]int{}
	for _, r := range right {
		if _, ok := unique[r]; !ok {
			continue
		}

		counts[r] += 1
	}

	similarity := 0
	for _, l := range left {
		count := counts[l]
		similarity += l * count
	}
	return similarity
}

func totalDistance(left, right []int) int {
	slices.Sort(left)
	slices.Sort(right)

	total := 0
	for i := range left {
		a, b := left[i], right[i]

		total += absDiff(a, b)
	}

	return total
}

func absDiff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}
