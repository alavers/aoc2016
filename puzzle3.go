package main

import "fmt"
import "strings"
import "os"
import "bufio"
import "regexp"
import "strconv"

func validate(sides []int) bool {
	if sides[0]+sides[1] <= sides[2] {
		return false
	}

	if sides[0]+sides[2] <= sides[1] {
		return false
	}

	if sides[1]+sides[2] <= sides[0] {
		return false
	}

	return true
}

func scan(f func(*bufio.Scanner)) {
	file, err := os.Open("data/3")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)

	f(scanner)
}

func row(scanner *bufio.Scanner) {
	good := 0
	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " ")
		splits := regexp.MustCompile(" +").Split(text, 3)
		sides := make([]int, len(splits))

		for i, split := range splits {
			sides[i], _ = strconv.Atoi(split)
		}

		if validate(sides) {
			good++
		}
	}

	fmt.Println(good)
}

func col(scanner *bufio.Scanner) {
	good := 0
	row := 0
	triangles := [][]int{make([]int, 3), make([]int, 3), make([]int, 3)}

	for scanner.Scan() {
		text := strings.Trim(scanner.Text(), " ")
		splits := regexp.MustCompile(" +").Split(text, 3)
		for i, split := range splits {
			triangles[i][row], _ = strconv.Atoi(split)
		}

		if row == 2 {
			for _, triangle := range triangles {
				if validate(triangle) {
					good++
				}
			}
		}

		row = (row + 1) % 3
	}

	fmt.Println(good)
}

func Puzzle3() {
	scan(row)
	scan(col)
}
