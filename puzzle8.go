package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

type Screen [][]byte

func NewScreen(x, y int) Screen {
	var s Screen
	s = make([][]byte, y)
	for i, _ := range s {
		s[i] = make([]byte, x)
		for j, _ := range s[i] {
			s[i][j] = ' '
		}
	}
	return s
}

func (s Screen) Draw() {
	for _, row := range s {
		fmt.Println(string(row))
	}
}

func (s Screen) Rect(x, y int) {
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			s[i][j] = '#'
		}
	}
}

func (s Screen) RotateColumn(col, n int) {
	n = n % len(s)
	n = len(s) - n
	column := make([]byte, len(s))
	for i := range s {
		column[i] = s[i][col]
	}
	column = append(column[n:], column[:n]...)
	for i := range column {
		s[i][col] = column[i]
	}
}

func (s Screen) RotateRow(row, n int) {
	r := s[row]
	n = n % len(s[row])
	n = len(s[row]) - n
	s[row] = append(r[n:], r[:n]...)
}

func (s Screen) Count() int {
	count := 0
	for i, _ := range s {
		for j, _ := range s[i] {
			if s[i][j] == '#' {
				count++
			}
		}
	}
	return count
}

func Puzzle8() {
	s := NewScreen(50, 6)
	file, err := os.Open("data/8")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tokens := strings.Split(scanner.Text(), " ")
		switch tokens[0] {
		case "rect":
			args := strings.Split(tokens[1], "x")
			x, err := strconv.Atoi(args[0])
			y, err := strconv.Atoi(args[1])
			check(err)
			s.Rect(x, y)
		case "rotate":
			i, err := strconv.Atoi(tokens[2][2:])
			by, err := strconv.Atoi(tokens[4])
			check(err)
			switch tokens[1] {
			case "column":
				s.RotateColumn(i, by)
			case "row":
				s.RotateRow(i, by)
			}
		}
	}
	fmt.Println(s.Count())
	s.Draw()
}
