package main

import (
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

var day = 5

func readInput() string {
	input_file_path := fmt.Sprintf("input_day_%d.txt", day)
	f, err := os.ReadFile(input_file_path)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) (map[int][]int, [][]int) {
	rules := map[int][]int{}
	updates := [][]int{}
	update := false
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		if line == "" {
			update = true
			continue
		}
		if !update {
			parts := strings.Split(line, "|")
			page, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
			dep, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
			rules[page] = append(rules[page], dep)

		} else {
			numbers := strings.Split(line, ",")
			tmp := []int{}
			for _, x := range numbers {
				n, _ := strconv.Atoi(strings.TrimSpace(x))
				tmp = append(tmp, n)
			}
			updates = append(updates, tmp)
		}
	}
	return rules, updates
}

func main() {
	rules, updates := parseInput(readInput())
	incorrect := []int{}

	// part 1
	sum1 := 0

outer:
	for p2, update := range updates {
		for i, page := range update {
			for _, dep := range rules[page] {
				if slices.Contains(update[i+1:], dep) {
					incorrect = append(incorrect, p2) // for part 2
					continue outer
				}
			}
		}
		sum1 += update[(len(update)-1)/2]
	}
	fmt.Println(sum1)

	// part 2
	sum2 := 0
	for _, index := range incorrect {
		update := updates[index]
		for i := len(update) - 1; i >= 0; i-- {
			page := update[i]
			for j := i - 1; j >= 0; j-- {
				remaining := update[j]
				if slices.Contains(rules[remaining], page) {
					update[j], update[i] = update[i], update[j]
					i++
					break
				}
			}
		}
		sum2 += update[(len(update)-1)/2]
	}
	fmt.Println(sum2)
}
