package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var day = 13

type button struct {
	X, Y int64
}

type prize struct {
	X, Y int64
}

type machine struct {
	A, B   button
	Target prize
}

func readInput(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) []machine {
	machines := []machine{}
	paragraphs := strings.Split(input, "\r\n\r\n")
	for _, para := range paragraphs {
		if strings.TrimSpace(para) == "" {
			continue
		}
		lines := strings.Split(para, "\r\n")
		if len(lines) != 3 {
			panic("parse para")
		}

		machines = append(machines, machine{
			A:      getButtonA(lines[0]),
			B:      getButtonB(lines[1]),
			Target: getPrize(lines[2]),
		})
	}
	return machines
}

func getButtonA(s string) button {
	var x, y int64
	fmt.Sscanf(s, "Button A: X+%d, Y+%d", &x, &y)
	return button{X: x, Y: y}
}
func getButtonB(s string) button {
	var x, y int64
	fmt.Sscanf(s, "Button B: X+%d, Y+%d", &x, &y)
	return button{X: x, Y: y}
}

func getPrize(s string) prize {
	var x, y int64
	fmt.Sscanf(s, "Prize: X=%d, Y=%d", &x, &y)
	return prize{X: x, Y: y}
}

func main() {
	// fileName := "sample.txt"
	fileName := fmt.Sprintf("input_day_%d.txt", day)
	machines := parseInput(readInput(fileName))
	// fmt.Printf("%+v\n", machines)
	var sum1 int64 = 0
	for _, m := range machines {
		sum1 += solve_p1(m)
	}
	fmt.Println(sum1)

	// p2
	var sum2 int64 = 0
	for i := range machines {
		machines[i].Target.X += 1e13
		machines[i].Target.Y += 1e13
		// fmt.Println(machines[i])
		sum2 += solve_p2(machines[i])
	}
	fmt.Println(sum2)
}

func solve_p2(m machine) int64 {
	// iAx + jBx = targetX - 1
	// iAy + jBy = targetY - 2
	// mutiple 1 by Ay and 2 by Ax
	// iAxAy + jBxAy = targetXAy - 1
	// iAyAx + jByAx = targetYAx - 2
	// subtract the 1 - 2 => i CANCELS OUT
	// coeff of j = BxAy-ByAx
	// target = targetXAy - targetYAx

	coeffJ := (m.B.X * m.A.Y) - (m.B.Y * m.A.X)
	target := (m.Target.X * m.A.Y) - (m.Target.Y * m.A.X)
	var floatJ = float64(target) / float64(coeffJ)
	j := int64(floatJ)
	i := (m.Target.X - (j * m.B.X)) / m.A.X

	if i*m.A.X+j*m.B.X != m.Target.X {
		j++
		i = (m.Target.X - (j * m.B.X)) / m.A.X
	}

	if i >= 0 && j >= 0 && (i*m.A.X)+(j*m.B.X) == m.Target.X && (i*m.A.Y)+(j*m.B.Y) == m.Target.Y {
		// fmt.Println(i, j)
		return 3*i + j

	}
	return 0
}

func solve_p1(m machine) int64 {
	sols := [][]int64{}
	for i := range 1_00 {
		for j := range 1_00 {
			if (m.A.X*int64(i)+m.B.X*int64(j)) == m.Target.X && (m.A.Y*int64(i)+m.B.Y*int64(j)) == m.Target.Y {
				sols = append(sols, []int64{int64(i), int64(j)})
			}
		}
	}

	if len(sols) == 0 {
		return 0
	}

	return (3*sols[0][0] + sols[0][1])
}
