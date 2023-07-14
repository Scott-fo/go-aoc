package day4

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Run() error {
	file, err := os.Open("./input/day4.txt")
	if err != nil {
		log.Fatal("Failed to open input")
	}
	defer file.Close()

	s := bufio.NewScanner(file)

	o := 0
	c := 0

	for s.Scan() {
		i := strings.Split(s.Text(), ",")

		l := strings.Split(i[0], "-")
		r := strings.Split(i[1], "-")

		ll, err := strToInt(l[0])
		if err != nil {
			return err
		}

		lh, err := strToInt(l[1])
		if err != nil {
			return err
		}

		rl, err := strToInt(r[0])
		if err != nil {
			return err
		}

		rh, err := strToInt(r[1])
		if err != nil {
			return err
		}

		if (ll <= rl && lh >= rh) || (rl <= ll && rh >= lh) {
			c += 1
		}

		if max(ll, rl) < min(lh, rh) {
			o += 1
		}
	}

	fmt.Println("Contains Count: ", c)
	fmt.Println("Overlap Count: ", o)

	return nil
}

func strToInt(s string) (int, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("Failed to convert string %s to int: %w", s, err)
	}

	return i, nil
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
