package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strings"
)

var day = 14

func readInput(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

type pos struct {
	x, y int
}

type vel struct {
	x, y int
}

type bot struct {
	p pos
	v vel
}

func (b *bot) calcPosAfterSeconds(s int, m, n int) {
	b.p.x = (((b.p.x + b.v.x*s) % m) + m) % m
	b.p.y = (((b.p.y + b.v.y*s) % n) + n) % n
	// fmt.Println(b.p)
}

func parseInput(input string) []bot {
	bots := []bot{}
	lines := strings.Split(input, "\r\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		bots = append(bots, getBot(line))
	}
	return bots
}

func getBot(s string) bot {
	bot := bot{}
	fmt.Sscanf(s, "p=%d,%d v=%d,%d", &bot.p.x, &bot.p.y, &bot.v.x, &bot.v.y)
	return bot
}

func printGrid(grid [][]rune) {
	var sb strings.Builder
	for i := range grid {
		sb.WriteString(string(grid[i]))
		sb.WriteString("\n")
	}
	fmt.Print(sb.String())
}

func clearGrid(grid [][]rune) {
	for i := range grid {
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
}

func main() {
	// fileName := "sample.txt"
	fileName := fmt.Sprintf("input_day_%d.txt", day)
	bots := parseInput(readInput(fileName))
	// fmt.Printf("%+v\n", bots)

	solve_p1(bots, 100, 101, 103)

	// fileName := "sample.txt"
	fileName = fmt.Sprintf("input_day_%d.txt", day)
	bots = parseInput(readInput(fileName))
	solve_p2(bots, 103, 101)
}

func solve_p2(bots []bot, m, n int) {
	scores := make([]int, 10001)
	for i := range scores {
		scores[i] = math.MinInt
	}

	for i := 1; i <= 10000; i++ {
		scores[i] = gridAfter(slices.Clone(bots), i, n, m, false)
	}

	fmt.Println(slices.Index(scores, slices.Max(scores)))
	// gridAfter(bots, slices.Index(scores, slices.Max(scores)), n, m, true)
}

func gridAfter(bots []bot, s, m, n int, print bool) int {
	grid := make([][]rune, n)
	for i := range grid {
		grid[i] = make([]rune, m)
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}
	for i := range bots {
		bots[i].calcPosAfterSeconds(s, m, n)
		grid[bots[i].p.y][bots[i].p.x] = '*'
	}

	if print {
		printGrid(grid)
	}

	return calc_grid_symm_score(grid)
}

func calc_grid_symm_score(grid [][]rune) int {
	score := 0
	width := len(grid[0])
	mid := (width / 2)

	for i := range grid {
		for j := range mid {
			if grid[i][mid-j] == '*' && grid[i][mid+j] == '*' {
				score += 1
			}
		}
	}
	return score
}

func solve_p1(bots []bot, s, m, n int) {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for i := range bots {
		bots[i].calcPosAfterSeconds(s, m, n)
		p := bots[i].p
		if p.x == m/2 || p.y == n/2 {
			continue
		}
		if p.x < m/2 && p.y < n/2 {
			q1++
		}
		if p.x > m/2 && p.y < n/2 {
			q2++
		}
		if p.x < m/2 && p.y > n/2 {
			q3++
		}
		if p.x > m/2 && p.y > n/2 {
			q4++
		}
	}
	fmt.Println(q1 * q2 * q3 * q4)
}
