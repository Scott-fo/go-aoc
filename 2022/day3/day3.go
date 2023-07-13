package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func getIntValue(r rune) int {
	if unicode.IsUpper(r) {
		return int(r%32) + 26
	} else {
		return int(r % 32)
	}
}

func part1() {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	sum := 0

	s := bufio.NewScanner(file)
	for s.Scan() {
		str := s.Text()
		l := len(str)
		m := l / 2

		c := map[byte]bool{}

		for i := 0; i < m; i++ {
			char := str[i]

			_, ok := c[char]
			if !ok {
				c[char] = true
			}
		}

		for j := m; j < l; j++ {
			char := str[j]
			_, ok := c[char]
			if ok {
				sum += getIntValue(rune(char))
				delete(c, char)
			}
		}
	}

	fmt.Println(sum)
}

func part2() {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		log.Fatal("Failed to open file")
	}

	sum := 0
	groupTracker := 0

	groupMaps := [3]map[byte]bool{
		make(map[byte]bool),
		make(map[byte]bool),
		make(map[byte]bool),
	}

	s := bufio.NewScanner(file)
	for s.Scan() {
		str := s.Text()

		for i := 0; i < len(str); i++ {
			char := str[i]

			if groupTracker == 2 && groupMaps[0][char] && groupMaps[1][char] {
				sum += getIntValue(rune(char))
				break
			}

			_, ok := groupMaps[groupTracker][char]
			if !ok {
				groupMaps[groupTracker][char] = true
			}
		}

		if groupTracker == 2 {
			groupTracker = 0
			for i := range groupMaps {
				groupMaps[i] = make(map[byte]bool)
			}
		} else {
			groupTracker++
		}
	}

	fmt.Println(sum)
}

func Run() {
	part2()
}
