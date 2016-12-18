package main

import "fmt"
import "os"
import "bufio"
import "io"

func decode(fname string, move func(int, rune) int) {
	file, err := os.Open(fname)
	check(err)
	defer file.Close()
	reader := bufio.NewReader(file)

	n := 5
	for {
		if r, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				check(err)
			}
		} else {
			if r == '\n' {
				fmt.Printf("%X", n)
			}
			n = move(n, r)
		}
	}
	fmt.Println()
}

// 1 2 3
// 4 5 6
// 7 8 9
func gridMove(n int, dir rune) int {
	switch dir {
	case 'U':
		if n > 3 {
			n = n - 3
		}
	case 'D':
		if n < 7 {
			n = n + 3
		}
	case 'L':
		if (n-1)%3 != 0 {
			n--
		}
	case 'R':
		if n%3 != 0 {
			n++
		}
	}
	return n
}

func Puzzle2a() {
	decode("data/2a", gridMove)
}

//     1
//   2 3 4
// 5 6 7 8 9
//   A B C
//     D
var starPad = struct {
	stuck map[rune]map[int]bool
	uStep map[int]int
	dStep map[int]int
}{
	stuck: map[rune]map[int]bool{
		'U': {1: true, 2: true, 4: true, 5: true, 9: true},
		'D': {5: true, 9: true, 0xA: true, 0xD: true, 0xC: true},
		'L': {1: true, 2: true, 5: true, 0xA: true, 0xD: true},
		'R': {1: true, 4: true, 9: true, 0xC: true, 0xD: true},
	},
	// 3, 0xD: sub 2 to move up
	// For the rest sub 4
	uStep: map[int]int{
		3:   -2,
		0xD: -2,
	},
	// 1, 0xB: add 2 to move down
	// For the rest, add 4
	dStep: map[int]int{
		1:   2,
		0xB: 2,
	},
}

func starMove(n int, dir rune) int {
	if _, ok := starPad.stuck[dir][n]; ok {
		return n
	}

	switch dir {
	case 'U':
		step, ok := starPad.uStep[n]
		if !ok {
			step = -4
		}
		n = n + step
	case 'D':
		step, ok := starPad.dStep[n]
		if !ok {
			step = 4
		}
		n = n + step
	case 'L':
		n = n - 1
	case 'R':
		n = n + 1
	}
	return n
}

func Puzzle2b() {
	decode("data/2a", starMove)
}
