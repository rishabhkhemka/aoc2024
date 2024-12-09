package main

import (
	"fmt"
	"log"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"
)

var day = 9

func readInput(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) [][]int {
	chars := strings.Split(input, "")
	out := [][]int{nil, nil}
	toggle := 0
	for _, char := range chars {
		if strings.TrimSpace(char) != "" {
			num, _ := strconv.Atoi(char)
			out[toggle] = append(out[toggle], num)
			toggle = (toggle + 1) % 2
		}
	}
	return out
}

var fileLengths []int
var spaceLengths []int

func PrintInput() {
	fmt.Println(len(fileLengths), len(spaceLengths))
	for id, v := range spaceLengths {
		fmt.Print(strings.Repeat(fmt.Sprintf("%d ", id), fileLengths[id]))
		fmt.Print(strings.Repeat(". ", v))
	}
	fmt.Print(strings.Repeat(fmt.Sprintf("%d ", len(fileLengths)-1), fileLengths[len(fileLengths)-1]))
	fmt.Println()
}

func main() {
	// fileName := "sample.txt"
	fileName := fmt.Sprintf("input_day_%d.txt", day)
	input := parseInput(readInput(fileName))
	fileLengths = input[0]
	spaceLengths = input[1]
	// PrintInput()
	solve_p1()

	input = parseInput(readInput(fileName))
	fileLengths = input[0]
	spaceLengths = input[1]
	solve_p2()
}

func solve_p1() {
	checksum := 0
	pos := 0
	left := 0
	right := len(fileLengths) - 1
	spaceIndex := 0

	for spaceIndex < len(spaceLengths) && right > left {
		for i := 0; i < fileLengths[left]; i++ {
			checksum += (left * pos)
			pos++
		}
		left++
	fillSpace:
		if spaceIndex < len(spaceLengths) && right > left {
			minVal := min(fileLengths[right], spaceLengths[spaceIndex])
			for i := 0; i < minVal; i++ {
				// fmt.Println(right, pos)
				checksum += (right * pos)
				pos++
			}

			if fileLengths[right] > spaceLengths[spaceIndex] {
				fileLengths[right] -= spaceLengths[spaceIndex]
				spaceIndex++
			} else {
				spaceLengths[spaceIndex] -= fileLengths[right]
				right--
				goto fillSpace
			}
		}
	}
	for i := 0; i < fileLengths[right]; i++ {
		checksum += (pos * left)
		pos++
	}
	fmt.Println(checksum)
}

type block struct {
	id   int
	size int
}

func solve_p2() {
	blocksToMove := map[int][]block{}
	for id, blockSize := range slices.Backward(fileLengths) {
		for spaceIndex := 0; spaceIndex < id; spaceIndex++ {
			if blockSize <= spaceLengths[spaceIndex] {
				spaceLengths[spaceIndex] -= blockSize
				blocksToMove[spaceIndex] = append(blocksToMove[spaceIndex], block{id: id, size: blockSize})
				break
			}
		}
	}

	blocksMovedIds := map[int]int{}
	for blocks := range maps.Values(blocksToMove) {
		for _, b := range blocks {
			blocksMovedIds[b.id] = fileLengths[b.id]
			fileLengths[b.id] = 0
		}
	}

	pos := 0
	checksum := 0
	for id, length := range fileLengths {
		for range length {
			checksum += (id * pos)
			pos++
			// fmt.Print(id, " ")
		}

		for range blocksMovedIds[id] {
			pos++
			// fmt.Print(". ")
		}

		spaceConsumed := 0
		for _, b := range blocksToMove[id] {
			for range b.size {
				checksum += (b.id * pos)
				pos++
				// fmt.Print(b.id, " ")
			}
			spaceConsumed += b.size
		}

		if id < len(spaceLengths) {
			for range spaceLengths[id] {
				pos++
				// fmt.Print(". ")
			}
		}
	}
	fmt.Println(checksum)
}
