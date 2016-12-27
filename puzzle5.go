package main

import "crypto/md5"
import "strconv"
import "fmt"

func computeA(prefix string) {
	n := -1
	password := ""
	for i := 0; i < 8; i++ {
		var test string
		var next byte
		for {
			n++
			test = prefix + strconv.Itoa(n)
			hash := md5.Sum([]byte(test))
			use := hash[0] == 0x0 &&
				hash[1] == 0x0 &&
				hash[2]&0xF0 == 0x0
			if use {
				next = hash[2] & 0x0F
				break
			}
		}

		password = password + fmt.Sprintf("%X", next)
		fmt.Println(password)
	}
}

func computeB(prefix string) {
	password := make([]rune, 8)
	n := -1
	remaining := make([]int, 8)
	for i := 0; i < 8; i++ {
		remaining[i] = i
		password[i] = '_'
	}

	for {
		n++
		test := prefix + strconv.Itoa(n)
		hash := md5.Sum([]byte(test))

		interesting := hash[0] == 0x0 &&
			hash[1] == 0x0 &&
			hash[2]&0xF0 == 0x0

		if !interesting {
			continue
		}

		var good bool
		var r, i int
		for _, r = range remaining {
			if good = hash[2] == byte(r); good {
				break
			}
			i++
		}

		if good {
			b := (hash[3] & 0xF0) >> 4
			password[remaining[i]] = rune(fmt.Sprintf("%X", b)[0])
			fmt.Println(string(password))
			remaining = append(remaining[:i], remaining[i+1:]...)

			if 0 == len(remaining) {
				break
			}
		}
	}
}

func Puzzle5() {
	input := "uqwqemis"
	computeA(input)
	computeB(input)
}
