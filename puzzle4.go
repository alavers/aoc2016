package main

import "os"
import "bufio"
import "fmt"
import "regexp"
import "strconv"
import "sort"
import "strings"

type Occur struct {
	letter rune
	count  int
}
type OccurList []Occur

func (os OccurList) Len() int      { return len(os) }
func (os OccurList) Swap(i, j int) { os[i], os[j] = os[j], os[i] }
func (os OccurList) Less(i, j int) bool {
	if os[i].count == os[j].count {
		return os[i].letter < os[j].letter
	}
	return os[i].count > os[j].count
}

type room struct {
	name     string
	sector   int
	checksum string
}

var roomPattern = regexp.MustCompile("([a-z-]+)-(\\d+)\\[([a-z]+)\\]")

func parse(s string) *room {
	matches := roomPattern.FindStringSubmatch(s)
	sector, err := strconv.Atoi(matches[2])
	check(err)

	return &room{
		name:     matches[1],
		sector:   sector,
		checksum: matches[3],
	}
}

func (r room) valid() bool {
	stripped := regexp.MustCompile("-").ReplaceAllString(r.name, "")

	counts := make(map[rune]int)
	for _, l := range stripped {
		ru := rune(l)
		counts[ru]++
	}
	occurList := make(OccurList, len(counts))
	occurIndex := 0
	for k, v := range counts {
		occurList[occurIndex] = Occur{k, v}
		occurIndex++
	}

	sort.Sort(occurList)
	checksum := make([]rune, 5)
	for i, occur := range occurList[0:5] {
		checksum[i] = occur.letter
	}

	return 0 == strings.Compare(string(checksum), r.checksum)
}

func (r room) decrypt() string {
	shift := rune(r.sector % 26)
	plaintext := make([]rune, len(r.name))
	for i, c := range r.name {
		if c != '-' {
			num := c - 'a'
			num = (num + shift) % 26
			plaintext[i] = num + 'a'
		} else {
			plaintext[i] = '-'
		}
	}
	return string(plaintext)
}

func Puzzle4() {
	file, err := os.Open("data/4")
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	tally := 0
	for scanner.Scan() {
		line := scanner.Text()
		room := parse(line)
		if room.valid() {
			tally += room.sector
			fmt.Println(room.decrypt(), room.sector)
		}
	}
	fmt.Println(tally)
}
