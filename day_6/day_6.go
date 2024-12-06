package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strings"
)

var day = 6

func readInput() string {
	input_file_path := fmt.Sprintf("input_day_%d.txt", day)
	f, err := os.ReadFile(input_file_path)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) [][]rune {
	lines := strings.Split(input, "\r\n")
	out := [][]rune{}
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			out = append(out, []rune(line))
		}
	}
	return out
}

var maze [][]rune

type pos struct {
	x, y int
}

func PrintMaze() {
	for _, row := range maze {
		fmt.Println(string(row))
	}
}

func main() {
	maze = parseInput(readInput())
	solve()
}

func part1(x, y int) (bool, map[pos][]int) {
	steps := map[pos][]int{}
	direction := 0
	prevX, prevY := x, y
	for withinBoundary(x, y) {
		if isObstacle(x, y) {
			direction = (direction + 1) % 4
			x, y = prevX, prevY
		}
		prevX, prevY = x, y
		key := pos{x: prevX, y: prevY}
		if v, ok := steps[key]; ok && slices.Contains(v, direction) {
			return true, nil
		}
		steps[key] = append(steps[key], direction)
		x, y = moveAhead(x, y, direction)
	}
	return false, steps
}

func solve() {
	startX, startY := 0, 0
outer:
	for i, row := range maze {
		for j, char := range row {
			if char == '^' {
				startX = i
				startY = j
				break outer
			}
		}
	}
	_, steps := part1(startX, startY)
	fmt.Println(len(steps))

	// part 2
	loops := map[pos]struct{}{}
	for key := range maps.Keys(steps) {
		maze[key.x][key.y] = '#'
		if ok, _ := part1(startX, startY); ok {
			loops[key] = struct{}{}
		}
		maze[key.x][key.y] = 'X'
	}
	for key := range loops {
		maze[key.x][key.y] = 'O'
	}
	fmt.Println(len(loops))
	// PrintMaze()
}

func moveAhead(i, j, dir int) (int, int) {
	switch dir {
	case 0:
		i--
	case 1:
		j++
	case 2:
		i++
	case 3:
		j--
	default:
		panic("wrong direction")
	}
	return i, j
}

func withinBoundary(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(maze) && j < len(maze[0])
}

func isObstacle(i, j int) bool {
	return withinBoundary(i, j) && maze[i][j] == '#'
}
