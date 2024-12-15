package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var G map[complex128]rune

func move(p, d complex128, test bool) bool {
	// test/check is a flag which checks if boxes can be moved
	// if it's not a check call, then do a check call (via short circuting)
	// if check result is false, return false
	// otherwise proceed, since check is passed and this is a move call
	if !test && !move(p, d, true) {
		return false
	}

	// starting positions for move
	// slice stores the multiple boxes to move in resursion
	ps := []complex128{p}
	if imag(d) != 0 { // bounds 0 check for row/col
		// [][]
		//  []
		if G[p] == '[' {
			ps = append(ps, p+1) // append the box pos of right side for row/col
		}
		if G[p] == ']' {
			ps = append(ps, p-1) // append the box pos of left side for row/col
		}
	}

	// ps now contains the position of bot or
	// empty space or
	// box pairs []

	for _, p := range ps {
		// for all chars in p -> bot, space, box []
		// test true always triggers first
		// => if we can't move the boxes recursively return false
		// => test = false is called when test result is true

		if test {
			// if it's a check call
			// for all ps, we are bfsing
			// p + d is  moving a step for all directions as per map
			// if you encounter a box again, recursively check to see if you can move
			// if you cant recursively move or you hit a wall, return false
			if (G[p+d] == '[' || G[p+d] == ']') && !move(p+d, d, true) || G[p+d] == '#' {
				return false
			}
		} else {
			// if not a check, need to perform move
			// before moving, recursively move the other blocks
			// if current part is a block
			if G[p+d] == '[' || G[p+d] == ']' {
				move(p+d, d, false)
			}
			// actual move after block moved or empty space
			// swap chars constitues a move, no matter what the char is
			G[p+d], G[p] = G[p], G[p+d]
		}
	}

	return true
}

func main() {
	file, err := os.Open("input_day_15.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	parts := strings.Split(string(data), "\r\n\r\n")
	grid := parts[0]
	moves := parts[1]

	grid = strings.ReplaceAll(grid, "#", "##")
	grid = strings.ReplaceAll(grid, ".", "..")
	grid = strings.ReplaceAll(grid, "O", "[]")
	grid = strings.ReplaceAll(grid, "@", "@.")

	G = make(map[complex128]rune)
	rows := strings.Split(grid, "\n")

	// parsing and storing the characters as a map
	// map[complex coordinate] = character
	// 2d grid -> math notation conversion
	// imagine the grid flipped across the x axis
	// grid row (i [0-m]) -> math +y axis flip of as it looks
	// grid col (j [0-n]) -> math +x axis as it looks
	for i, row := range rows {
		for j, c := range row {
			G[complex(float64(j), float64(i))] = c
		}
	}

	// find starting position of the bot
	var p complex128
	for pos, c := range G {
		if c == '@' {
			p = pos
			break
		}
	}

	// direction map follows math notation
	// +x is 1 +y is -i instead of i
	// moving in ^ corresponds to moving up in rows
	// which corresponds to moving down in +y axis
	// as per the grid to complex plane conversion
	moveMap := map[rune]complex128{
		'<': -1,
		'>': 1,
		'^': -1i,
		'v': 1i,
	}

	for _, m := range moves {
		d := moveMap[m]
		if move(p, d, false) { // moves the bot test=false flag,
			// it calls itself with true first to check if move is possible
			// and moves the box when it is
			p += d // updates position of the bot
		}
	}

	// calculate gps score
	sum := 0.0
	for pos, c := range G {
		if c == '[' {
			sum += real(pos) + 100*imag(pos)
		}
	}

	fmt.Printf("%.0f\n", sum)
}
