package main

import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

const (
	North = iota
	East
	South
	West
)

func turn(d int, t byte) (r int) {
	switch t {
	case 'L':
		r = (d + 3) % 4
	case 'R':
		r = (d + 1) % 4
	}
	return
}

type Pos struct {
	x   int
	y   int
	dir int
}

func (p *Pos) move(turn, blocks int) {
	switch turn {
	case 'L':
		p.dir = (p.dir + 3) % 4
	case 'R':
		p.dir = (p.dir + 1) % 4
	}

	switch p.dir {
	case North:
		p.y = p.y + blocks
	case East:
		p.x = p.x + blocks
	case South:
		p.y = p.y - blocks
	case West:
		p.x = p.x - blocks
	}
}

func Abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func Puzzle1a() {
	b, err := ioutil.ReadFile("data/1a")
	check(err)
	steps := strings.Split(string(b), ", ")
	var pos = Pos{0, 0, North}
	fmt.Printf("Staring position: %v\n", pos)

	for _, step := range steps {
		step = strings.TrimSpace(step)
		turn := step[0]
		blocks, err := strconv.ParseInt(step[1:], 10, 32)
		check(err)
		pos.move(int(turn), int(blocks))
	}

	fmt.Printf("Ending position: %v\n", pos)
	fmt.Printf("Blocks: %v\n", Abs(pos.x)+Abs(pos.y))
}

type Grid struct {
	m map[int]map[int]bool
}

func (g *Grid) mark(x, y int) bool {
	if _, ok := g.m[x]; !ok {
		g.m[x] = make(map[int]bool)
	}
	if _, ok := g.m[x][y]; !ok {
		g.m[x][y] = true
		return false
	}
	return true
}
func Puzzle1b() {
	b, err := ioutil.ReadFile("data/1a")
	check(err)
	steps := strings.Split(string(b), ", ")
	var pos = Pos{0, 0, North}
	fmt.Printf("Staring position: %v\n", pos)

	g := &Grid{make(map[int]map[int]bool)}

Loop:
	for _, step := range steps {
		step = strings.TrimSpace(step)
		turn := step[0]
		blocks, err := strconv.ParseInt(step[1:], 10, 32)
		check(err)

		switch turn {
		case 'L':
			pos.dir = (pos.dir + 3) % 4
		case 'R':
			pos.dir = (pos.dir + 1) % 4
		}

		for i := 0; i < int(blocks); i++ {
			switch pos.dir {
			case North:
				pos.y = pos.y + 1
			case East:
				pos.x = pos.x + 1
			case South:
				pos.y = pos.y - 1
			case West:
				pos.x = pos.x - 1
			}

			if g.mark(pos.x, pos.y) {
				fmt.Println("Already visited: ", pos)
				break Loop
			}
		}
	}

	fmt.Printf("Ending position: %v\n", pos)
	fmt.Printf("Blocks: %v\n", Abs(pos.x)+Abs(pos.y))
}
