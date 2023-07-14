package day6

import (
	"fmt"
	"os"
)

func Run() error {
	data, err := os.ReadFile("./input/day6.txt")
	if err != nil {
		return fmt.Errorf("Failed to read file")
	}

	head := 14
	tail := 0

	runes := []rune(string(data))

	for head < len(runes) {
		if !repeats(runes[tail:head]) {
			fmt.Println("Marker: ", head)
			return nil
		}

		head += 1
		tail += 1
	}

	return fmt.Errorf("Marker not found")
}

func repeats(rs []rune) bool {
	s := make(map[rune]bool)
	for _, r := range rs {
		if s[r] {
			return true
		} else {
			s[r] = true
		}
	}

	return false
}
