package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var day = 1

func readInput() string {
	input_file_path := fmt.Sprintf("input_day_%d.txt", day)
	f, err := os.ReadFile(input_file_path)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) ([]int, []int) {
	var left, right []int
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		numbers := strings.SplitN(line, " ", 2)
		num1, _ := strconv.Atoi(strings.TrimSpace(numbers[0]))
		num2, _ := strconv.Atoi(strings.TrimSpace(numbers[1]))
		left = append(left, num1)
		right = append(right, num2)
	}
	return left, right
}

func main() {
	left, right := parseInput(readInput())
	slices.Sort(left)
	slices.Sort(right)
	sum := 0
	for i, l := range left {
		if l > right[i] {
			sum += l - right[i]
		} else {
			sum += right[i] - l
		}
	}

	rightCount := map[int]int{}
	for _, r := range right {
		rightCount[r]++
	}

	score := 0
	for _, l := range left {
		score += (l * rightCount[l])
	}

	fmt.Println(sum)
	fmt.Println(score)
}
