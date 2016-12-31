package main

import "fmt"
import "strconv"
import "io/ioutil"
import "strings"

func ParseTag(line []byte, i int) (len, mult, next int) {
	tokenStart := i + 1

	for line[i] != 'x' {
		i++
	}
	len, err := strconv.Atoi(string(line[tokenStart:i]))

	tokenStart = i + 1
	for line[i] != ')' {
		i++
	}
	mult, err = strconv.Atoi(string(line[tokenStart:i]))

	next = i + 1
	check(err)
	return
}

func FillBuffer(line []byte, len, i int) (buffer []byte, nested bool, next int) {
	buffer = make([]byte, len)
	nested = false
	for j := 0; j < len; j++ {
		if line[i] == '(' {
			nested = true
		}
		buffer[j] = line[i]
		i++
	}
	next = i
	return
}

func Decompress(line []byte) ([]byte, bool) {
	var m_len, m_mult, i int
	var out []byte
	var buffer []byte
	var nest bool
	r_nest := false

	for i < len(line) {
		if line[i] == '(' {
			m_len, m_mult, i = ParseTag(line, i)
			buffer, nest, i = FillBuffer(line, m_len, i)
			r_nest = r_nest || nest

			for j := 0; j < m_mult; j++ {
				out = append(out, buffer...)
			}
		} else {
			out = append(out, line[i])
			i++
		}
	}

	return out, r_nest
}

func DecompressCount(line []byte, offset int, length int) int {
	count := 0
	var m_len, m_mult, i int
	i = offset

	for i < offset+length {
		if line[i] == '(' {
			m_len, m_mult, i = ParseTag(line, i)
			subcount := DecompressCount(line, i, m_len)
			count = count + m_mult*subcount
			i = i + m_len
		} else {
			count++
			i++
		}
	}

	return count
}

func Puzzle9() {
	input := "data/9"
	raw, err := ioutil.ReadFile(input)
	check(err)
	compressed := []byte(strings.TrimSpace(string(raw)))

	bytes, _ := Decompress(compressed)
	fmt.Println(len(bytes))

	count := DecompressCount(compressed, 0, len(compressed))
	fmt.Println(count)
}
