package main

import "os"
import "bufio"
import "fmt"

type Column map[rune]int

func process(filename string, best func(c Column) rune) {
	file, err := os.Open(filename)
	check(err)
	defer file.Close()

	peek, err := bufio.NewReader(file).ReadString('\n')
	check(err)
	columns := make([]Column, len(peek))
	for i := range columns {
		columns[i] = make(Column)
	}

	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for i, c := range line {
			columns[i][c]++
		}
	}

	for _, c := range columns {
		fmt.Printf("%c", best(c))
	}
	fmt.Print("\n")
}

func Puzzle6() {
	input := "data/6_test"

	// Most common
	process(input, func(c Column) rune {
		bestCount := 0
		var bestRune rune
		for k, v := range c {
			if v > bestCount {
				bestRune, bestCount = k, v
			}
		}
		return bestRune
	})

	// Least common
	process(input, func(c Column) rune {
		bestCount := 1 << 32
		var bestRune rune
		for k, v := range c {
			if v < bestCount {
				bestRune, bestCount = k, v
			}
		}
		return bestRune
	})
}
