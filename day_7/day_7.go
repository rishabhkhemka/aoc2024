package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
	"slices"
	"strings"
	"sync"
)

var day = 7

func readInput(fileName string) string {
	f, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}
	return string(f)
}

func parseInput(input string) map[*big.Int][]*big.Int {
	lines := strings.Split(input, "\r\n")
	out := map[*big.Int][]*big.Int{}
	for _, line := range lines {
		if strings.TrimSpace(line) != "" {
			numbers := strings.Split(line, ":")
			result := new(big.Int)
			result.SetString(numbers[0], 10)
			operands := []*big.Int{}
			for _, num := range strings.Split(numbers[1], " ") {
				if strings.TrimSpace(num) != "" {
					n := new(big.Int)
					n.SetString(num, 10)
					operands = append(operands, n)
				}
			}
			out[result] = operands
		}
	}
	return out
}

var equations map[*big.Int][]*big.Int
var opts = []rune{'+', '*'}

func PrintEquations() {
	for result, operands := range equations {
		fmt.Printf("%s: %v\n", result.String(), operands)
	}
}

func main() {
	fileName := fmt.Sprintf("input_day_%d.txt", day)
	// fileName := "sample.txt"
	equations = parseInput(readInput(fileName))
	solve()
}

func solve() {
	sum1 := new(big.Int)
	sum2 := new(big.Int)
	wg1, wg2 := sync.WaitGroup{}, sync.WaitGroup{}
	sum1chan := make(chan *big.Int, len(equations))
	sum2chan := make(chan *big.Int, len(equations))
	wg1.Add(2)
	go func() {
		for v := range sum1chan {
			sum1.Add(sum1, v)
		}
		wg1.Done()
	}()

	go func() {
		for v := range sum2chan {
			sum2.Add(sum2, v)
		}
		wg1.Done()
	}()
	for n, ops := range equations {
		wg2.Add(2)
		ops = slices.Clone(ops)
		slices.Reverse(ops)

		go func() {
			// part 1
			defer wg2.Done()
			allCombosP1 := tryCombosP1(ops)
			// fmt.Println(n, ops, allCombosP1)
			for _, combo := range allCombosP1 {
				if combo.Cmp(n) == 0 {
					sum1chan <- n
					break
				}
			}
		}()

		go func() {
			// part 2
			defer wg2.Done()
			allCombosP2 := tryCombosP2(ops)
			// fmt.Println(n, ops, allCombosP2)
			for _, combo := range allCombosP2 {
				if combo.Cmp(n) == 0 {
					sum2chan <- n
					break
				}
			}
		}()
	}
	wg2.Wait()
	close(sum1chan)
	close(sum2chan)
	wg1.Wait()
	fmt.Println(sum1.String())
	fmt.Println(sum2.String())
}

func tryCombosP1(ops []*big.Int) []*big.Int {
	if len(ops) == 1 {
		return []*big.Int{ops[0]}
	}

	partialCombos := tryCombosP1(ops[1:])
	n := len(partialCombos)
	for i := 0; i < n; i++ {
		addResult := new(big.Int).Add(partialCombos[i], ops[0])
		mulResult := new(big.Int).Mul(partialCombos[i], ops[0])
		partialCombos = append(partialCombos, mulResult)
		partialCombos[i] = addResult
	}
	return partialCombos
}

func tryCombosP2(ops []*big.Int) []*big.Int {
	if len(ops) == 1 {
		return []*big.Int{ops[0]}
	}

	partialCombos := tryCombosP2(ops[1:])
	n := len(partialCombos)
	for i := 0; i < n; i++ {
		addResult := new(big.Int).Add(partialCombos[i], ops[0])
		mulResult := new(big.Int).Mul(partialCombos[i], ops[0])
		concatResult := concatBigInts(ops[0], partialCombos[i])
		partialCombos = append(partialCombos, mulResult, concatResult)
		partialCombos[i] = addResult
	}
	return partialCombos
}

func concatBigInts(a, b *big.Int) *big.Int {
	aStr := a.String()
	bStr := b.String()
	concatStr := bStr + aStr
	concatResult := new(big.Int)
	concatResult.SetString(concatStr, 10)
	return concatResult
}
