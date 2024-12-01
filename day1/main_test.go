package day1

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_totalDistance(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	left, right, err := parseInput("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				left:  left,
				right: right,
			},
			want: 2192892,
		},
		{
			args: args{
				left:  []int{3, 4, 2, 1, 3, 3},
				right: []int{4, 3, 5, 3, 9, 3},
			},
			want: 11,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := totalDistance(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("totalDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func parseInput(path string) ([]int, []int, error) {
	input, err := os.ReadFile(path)
	if err != nil {
		return nil, nil, err
	}

	lines := strings.Split(string(input), "\n")
	left, right := make([]int, 0, len(lines)), make([]int, 0, len(lines))

	for _, line := range lines {
		parts := strings.Split(line, "   ")
		l, err := strconv.Atoi(strings.TrimSpace(parts[0]))
		if err != nil {
			return nil, nil, err
		}
		r, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return nil, nil, err
		}

		left = append(left, l)
		right = append(right, r)
	}
	return left, right, nil
}

func Test_similarityScore(t *testing.T) {
	type args struct {
		left  []int
		right []int
	}
	left, right, err := parseInput("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				left:  left,
				right: right,
			},
			want: 22962826,
		},
		{
			args: args{
				left:  []int{3, 4, 2, 1, 3, 3},
				right: []int{4, 3, 5, 3, 9, 3},
			},
			want: 31,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := similarityScore(tt.args.left, tt.args.right); got != tt.want {
				t.Errorf("similarityScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
