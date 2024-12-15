package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var day = 15

func readInput(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) ([][]rune, []rune) {
	para := strings.Split(input, "\r\n\r\n")
	grid := [][]rune{}
	for _, line := range strings.Split(para[0], "\r\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		grid = append(grid, []rune(line))
	}

	moves := []rune{}
	for _, line := range strings.Split(para[1], "\r\n") {
		if strings.TrimSpace(line) == "" {
			continue
		}
		moves = append(moves, []rune(line)...)
	}
	return grid, moves
}

func printGrid(grid [][]rune) {
	for i := range grid {
		fmt.Println(string(grid[i]))
	}
}

func getBotLoc() (int, int) {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				return i, j
			}
		}
	}
	return -1, -1
}

func withinBounds(i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])
}

func isWall(i, j int) bool {
	return !withinBounds(i, j) || grid[i][j] == '#'
}

func isBox(i, j int) bool {
	return withinBounds(i, j) && grid[i][j] == 'O'
}

func calcSumGPS() int {
	sum := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'O' || grid[i][j] == '[' {
				sum += (i*100 + j)
			}
		}
	}
	return sum
}

func getEmptySpaces(i, j int, direction rune) [][]int {
	emptySpaces := [][]int{}
	switch direction {
	case '>':
		for k := 1; k < len(grid[0]) && !isWall(i, j+k); k++ {
			if grid[i][j+k] == '.' {
				emptySpaces = append(emptySpaces, []int{i, j + k})
			}
		}
	case '<':
		for k := 1; k < len(grid[0]) && !isWall(i, j-k); k++ {
			if grid[i][j-k] == '.' {
				emptySpaces = append(emptySpaces, []int{i, j - k})
			}
		}
	case '^':
		for k := 1; k < len(grid) && !isWall(i-k, j); k++ {
			if grid[i-k][j] == '.' {
				emptySpaces = append(emptySpaces, []int{i - k, j})
			}
		}
	case 'v':
		for k := 1; k < len(grid) && !isWall(i+k, j); k++ {
			if grid[i+k][j] == '.' {
				emptySpaces = append(emptySpaces, []int{i + k, j})
			}
		}
	}
	return emptySpaces
}

func chunkSubMoves() {
	subMoves = [][]rune{}
	currentMove := []rune{}
	for _, move := range moves {
		if len(currentMove) == 0 || currentMove[len(currentMove)-1] == move {
			currentMove = append(currentMove, move)
		} else {
			subMoves = append(subMoves, currentMove)
			currentMove = []rune{move}
		}
	}
	if len(currentMove) != 0 {
		subMoves = append(subMoves, currentMove)
	}
	// printGrid(subMoves)
}

func moveBot(i, j int, times int, direction rune) (int, int, int) {
	spaces := getEmptySpaces(i, j, direction)
	if len(spaces) == 0 {
		return i, j, 0
	}
	switch direction {
	case '>':
		for times != 0 && len(spaces) > 0 && !isWall(i, j+1) {
			spacei, spacej := spaces[0][0], spaces[0][1]
			nexti, nextj := i, j+1
			grid[nexti][nextj], grid[spacei][spacej] = grid[spacei][spacej], grid[nexti][nextj]

			spaces = spaces[1:]
			times--
			grid[i][j] = '.'
			i, j = nexti, nextj
			grid[i][j] = '@'
		}
	case '<':
		for times != 0 && len(spaces) > 0 && !isWall(i, j-1) {
			spacei, spacej := spaces[0][0], spaces[0][1]
			nexti, nextj := i, j-1
			grid[nexti][nextj], grid[spacei][spacej] = grid[spacei][spacej], grid[nexti][nextj]

			spaces = spaces[1:]
			times--
			grid[i][j] = '.'
			i, j = nexti, nextj
			grid[i][j] = '@'
		}
	case '^':
		for times != 0 && len(spaces) > 0 && !isWall(i-1, j) {
			spacei, spacej := spaces[0][0], spaces[0][1]
			nexti, nextj := i-1, j
			grid[nexti][nextj], grid[spacei][spacej] = grid[spacei][spacej], grid[nexti][nextj]

			spaces = spaces[1:]
			times--
			grid[i][j] = '.'
			i, j = nexti, nextj
			grid[i][j] = '@'
		}
	case 'v':
		for times != 0 && len(spaces) > 0 && !isWall(i+1, j) {
			spacei, spacej := spaces[0][0], spaces[0][1]
			nexti, nextj := i+1, j
			grid[nexti][nextj], grid[spacei][spacej] = grid[spacei][spacej], grid[nexti][nextj]

			spaces = spaces[1:]
			times--
			grid[i][j] = '.'
			i, j = nexti, nextj
			grid[i][j] = '@'
		}
	}
	return i, j, times
}

func solve_p1() {
	i, j := getBotLoc()
	for _, subMove := range subMoves {
		dir := subMove[0]
		times := len(subMove)
		i, j, times = moveBot(i, j, times, dir)
		// printGrid(grid)
	}
}

func get2Xgrid() [][]rune {
	newGrid := [][]rune{}
	for i := range grid {
		row := []rune{}
		for j := range grid[i] {
			switch grid[i][j] {
			case '#':
				row = append(row, '#', '#')
			case 'O':
				row = append(row, '[', ']')
			case '.':
				row = append(row, '.', '.')
			case '@':
				row = append(row, '@', '.')
			}
		}
		newGrid = append(newGrid, row)
	}
	return newGrid
}

var grid [][]rune
var moves []rune
var subMoves [][]rune

func main() {
	// fileName := "sample.txt"
	fileName := fmt.Sprintf("input_day_%d.txt", day)
	grid, moves = parseInput(readInput(fileName))
	chunkSubMoves()
	solve_p1()
	fmt.Println(calcSumGPS())

	// printGrid(grid)
	// grid, moves = parseInput(readInput(fileName))
	// grid = get2Xgrid()
	// chunkSubMoves()
	// printGrid(grid)
	// fmt.Println(calcSumGPS())
}
