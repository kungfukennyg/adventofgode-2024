package day2

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test_countSafeReports(t *testing.T) {
	type args struct {
		reports      [][]int
		withDampener bool
	}
	input, err := loadReport("input.txt")
	if err != nil {
		t.Fatal(err)
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args: args{
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
				withDampener: false,
			},
			want: 2,
		},
		{
			args: args{
				reports:      input,
				withDampener: false,
			},
			want: 524,
		},
		{
			args: args{
				reports: [][]int{
					{7, 6, 4, 2, 1},
					{1, 2, 7, 8, 9},
					{9, 7, 6, 2, 1},
					{1, 3, 2, 4, 5},
					{8, 6, 4, 4, 1},
					{1, 3, 6, 7, 9},
				},
				withDampener: true,
			},
			want: 4,
		},
		{
			args: args{
				reports:      input,
				withDampener: true,
			},
			want: 569,
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if got := countSafe(tt.args.reports, tt.args.withDampener); got != tt.want {
				t.Errorf("countSafeReports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func loadReport(path string) ([][]int, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	reports := [][]int{}
	for _, line := range strings.Split(string(data), "\n") {
		report := []int{}
		for _, level := range strings.Split(line, " ") {
			l, err := strconv.Atoi(level)
			if err != nil {
				return nil, err
			}
			report = append(report, l)
		}
		reports = append(reports, report)
	}
	return reports, nil
}
