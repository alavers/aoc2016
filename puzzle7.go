package main

import "os"
import "bufio"
import "fmt"

type Abba struct {
	len         int
	insertPoint int
	buffer      []rune
}

func (a *Abba) ReadRune(r rune) bool {
	len := len(a.buffer)
	a.buffer[a.insertPoint] = r
	accepted := true

	for i := 0; i < a.insertPoint+1; i++ {
		left := a.buffer[i]
		right := a.buffer[len-1-i]
		if right != 0 && left != right {
			accepted = false
			break
		}
	}

	if a.insertPoint > 1 && a.buffer[0] == a.buffer[1] {
		accepted = false
	}

	if accepted {
		a.insertPoint++
	} else {
		a.buffer = append(a.buffer[1:], 0)
	}

	return a.insertPoint == len
}

func (a *Abba) Reset() {
	a.buffer = make([]rune, a.len)
	a.insertPoint = 0
}

func ValidateTls(address string) bool {
	len := 4
	abba := Abba{len, 0, make([]rune, len)}
	inBrackets := false
	valid := false
	foundAbba := false
	for _, r := range address {
		switch r {
		case '[':
			inBrackets = true
			abba.Reset()
			continue
		case ']':
			inBrackets = false
			abba.Reset()
			continue
		}

		valid = abba.ReadRune(r)

		if valid {
			abba.Reset()
			if inBrackets {
				return false
			} else {
				foundAbba = true
			}
		}
	}
	return foundAbba
}

func ContainsInverse(m map[string]bool, s string) bool {
	inverse := string([]byte{s[1], s[0], s[1]})
	_, ok := m[inverse]
	return ok
}

func ValidateSsl(s string) bool {
	valid := false
	hypertext := false
	abas := make(map[string]bool)
	babs := make(map[string]bool)
	for i, r := range s {
		if i > len(s)-3 {
			break
		}

		if r == '[' {
			hypertext = true
			continue
		} else if r == ']' {
			hypertext = false
			continue
		}

		if s[i+1] == '[' || s[i+1] == ']' || s[i+2] == '[' || s[i+2] == ']' {
			continue
		}

		if s[i] != s[i+1] && s[i] == s[i+2] {
			aba := string(s[i : i+3])
			if hypertext {
				babs[aba] = true
				if ContainsInverse(abas, aba) {
					valid = true
					break
				}
			} else {
				abas[aba] = true
				if ContainsInverse(babs, aba) {
					valid = true
					break
				}
			}
		}
	}
	return valid
}

func Puzzle7() {
	input := "data/7"
	file, err := os.Open(input)
	check(err)
	defer file.Close()

	validTls := 0
	validSsl := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if ValidateTls(text) {
			validTls++
		}
		if ValidateSsl(text) {
			validSsl++
		}
	}

	fmt.Println(validTls)
	fmt.Println(validSsl)
}
