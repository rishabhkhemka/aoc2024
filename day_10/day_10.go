package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var day = 10

func readInput(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) [][]int {
	out := [][]int{}
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			chars := strings.Split(line, "")
			nums := []int{}
			for _, char := range chars {
				if strings.TrimSpace(char) != "" {
					num, err := strconv.Atoi(char)
					if err != nil {
						nums = append(nums, -2)
					} else {
						nums = append(nums, num)
					}
				}
			}
			out = append(out, nums)
		}
	}
	return out
}

var maze [][]int

func PrintInput() {
	for _, v := range maze {
		fmt.Println(v)
	}
}

func main() {
	// fileName := "sample.txt"
	fileName := fmt.Sprintf("input_day_%d.txt", day)
	maze = parseInput(readInput(fileName))
	// PrintInput()
	solve_p1()
}

type pos struct {
	i, j int
}

func withinBounds(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(maze) && j < len(maze[0])
}

func canContinue(height, x, y int) bool {
	return withinBounds(x, y) && (maze[x][y]-height == 1)
}

func solve_p1() {
	trailheads := []pos{}
	for i, row := range maze {
		for j, el := range row {
			if el == 0 {
				trailheads = append(trailheads, pos{i: i, j: j})
			}
		}
	}
	// fmt.Println(len(trailheads), trailheads)

	score := 0
	for _, start := range trailheads {
		trails := p1_helper(start)
		// fmt.Println(start, trails)
		score += len(trails)
	}
	fmt.Println(score)

	rating := 0
	for _, start := range trailheads {
		trails := p2_helper(start)
		// fmt.Println(start, trails)
		rating += len(trails)
	}
	fmt.Println(rating)
}

func p2_helper(start pos) [][]pos {
	if !withinBounds(start.i, start.j) || (maze[start.i][start.j] == -1) {
		return nil
	}

	if maze[start.i][start.j] == 9 {
		return [][]pos{{start}}
	}

	// PrintInput()
	tmp := maze[start.i][start.j]
	maze[start.i][start.j] = -1

	trails := [][]pos{}
	if canContinue(tmp, start.i+1, start.j) {
		trails = append(trails, p2_helper(pos{start.i + 1, start.j})...)
	}
	if canContinue(tmp, start.i-1, start.j) {
		trails = append(trails, p2_helper(pos{start.i - 1, start.j})...)
	}
	if canContinue(tmp, start.i, start.j+1) {
		trails = append(trails, p2_helper(pos{start.i, start.j + 1})...)
	}
	if canContinue(tmp, start.i, start.j-1) {
		trails = append(trails, p2_helper(pos{start.i, start.j - 1})...)
	}

	for i := range trails {
		trails[i] = append(trails[i], start)
	}
	maze[start.i][start.j] = tmp
	return trails
}

func p1_helper(start pos) []pos {
	if !withinBounds(start.i, start.j) || (maze[start.i][start.j] == -1) {
		return nil
	}

	if maze[start.i][start.j] == 9 {
		return []pos{start}
	}

	// PrintInput()
	tmp := maze[start.i][start.j]
	maze[start.i][start.j] = -1

	trails := []pos{}
	if canContinue(tmp, start.i+1, start.j) {
		trails = append(trails, p1_helper(pos{start.i + 1, start.j})...)
	}
	if canContinue(tmp, start.i-1, start.j) {
		trails = append(trails, p1_helper(pos{start.i - 1, start.j})...)
	}
	if canContinue(tmp, start.i, start.j+1) {
		trails = append(trails, p1_helper(pos{start.i, start.j + 1})...)
	}
	if canContinue(tmp, start.i, start.j-1) {
		trails = append(trails, p1_helper(pos{start.i, start.j - 1})...)
	}

	trails = getUniquePos(trails) // disable this for p2 also works
	maze[start.i][start.j] = tmp
	return trails
}

func getUniquePos(trails []pos) []pos {
	unique := map[pos]struct{}{}
	for _, v := range trails {
		unique[v] = struct{}{}
	}
	trails = nil
	for k := range unique {
		trails = append(trails, k)
	}
	return trails
}
