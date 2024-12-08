package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var day = 8

func readInput(fileName string) string {
	f, err := os.ReadFile(fileName)
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

var mat [][]rune
var antinodes_p1 map[point]int
var antinodes_p2 map[point]int

func PrintMatrix() {
	for _, line := range mat {
		fmt.Println(string(line))
	}
}

func main() {
	// fileName := "sample.txt"
	fileName := fmt.Sprintf("input_day_%d.txt", day)
	mat = parseInput(readInput(fileName))
	// PrintMatrix()
	antinodes_p1 = map[point]int{}
	antinodes_p2 = map[point]int{}

	p1()
}

type point struct {
	x, y int
}

func p1() {
	fMap := map[string][]point{}
	for i, line := range mat {
		for j, char := range line {
			if char != '.' {
				fMap[string(char)] = append(fMap[string(char)], point{x: i, y: j})
			}
		}
	}
	for _, antennas := range fMap {
		for i, p1 := range antennas {
			for _, p2 := range antennas[i+1:] {
				if p1.x > p2.x {
					// p1 is a lower point
					mark_antinodes_1(p1, p2)
					mark_antinodes_2(p1, p2)
				} else {
					mark_antinodes_1(p2, p1)
					mark_antinodes_2(p2, p1)
				}
			}
		}
	}

	fmt.Println(len(antinodes_p1))
	fmt.Println(len(antinodes_p2))
}

func mark_antinodes_2(p1, p2 point) {
	// fmt.Println(p1, p2)
	dy := abs(p1.y - p2.y)
	dx := abs(p1.x - p2.x)
	slope := float64(p2.y-p1.y) / float64(p2.x-p1.x)
	// fmt.Println(slope)
	if slope > 0 {
		// p2-p1 is like \
		for i := range 50 {
			mark_point_p2('#', point{x: p1.x + i*dx, y: p1.y + i*dy})
			mark_point_p2('#', point{x: p1.x - i*dx, y: p1.y - i*dy})
		}
	} else {
		// p2-p1 is like /
		for i := range 50 {
			mark_point_p2('#', point{x: p1.x + i*dx, y: p1.y - i*dy})
			mark_point_p2('#', point{x: p1.x - i*dx, y: p1.y + i*dy})
		}
	}

}

func mark_antinodes_1(p1, p2 point) {
	// fmt.Println(p1, p2)
	dy := abs(p1.y - p2.y)
	dx := abs(p1.x - p2.x)
	if p1.x == p2.x { // input doesn't have this case
		// same row
		mark_point_p1('#', point{x: p1.x, y: min(p1.y, p2.y) - dy})
		mark_point_p1('#', point{x: p1.x, y: max(p1.y, p2.y) + dy})
	} else if p1.y == p2.y { // input doesn't have this case
		// same column
		mark_point_p1('#', point{y: p1.y, x: p1.x - dx})
		mark_point_p1('#', point{y: p1.y, x: p2.x + dx})
	} else {
		slope := float64(p2.y-p1.y) / float64(p2.x-p1.x)
		// fmt.Println(slope)
		if slope > 0 {
			// p2-p1 is like \
			mark_point_p1('#', point{x: p1.x + dx, y: p1.y + dy})
			mark_point_p1('#', point{x: p2.x - dx, y: p2.y - dy})
		} else {
			// p2-p1 is like /
			mark_point_p1('#', point{x: p1.x + dx, y: p1.y - dy})
			mark_point_p1('#', point{x: p2.x - dx, y: p2.y + dy})
		}
	}
}

func mark_point_p1(char rune, points ...point) {
	for _, p := range points {
		if withinBounds(p) {
			mat[p.x][p.y] = char
			antinodes_p1[p]++
		}
	}
}

func mark_point_p2(char rune, points ...point) {
	for _, p := range points {
		if withinBounds(p) {
			mat[p.x][p.y] = char
			antinodes_p2[p]++
		}
	}
}

func withinBounds(p point) bool {
	return p.x >= 0 && p.y >= 0 && p.x < len(mat) && p.y < len(mat[0])
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
