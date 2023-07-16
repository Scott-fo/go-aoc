package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func partOne() error {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	p := 0
	c := -1

	s := bufio.NewScanner(file)
	for s.Scan() {
		str := s.Text()
		i, err := strconv.Atoi(str)
		if err != nil {
			return fmt.Errorf("Failed to parse depth")
		}

		if i > p {
			c++
		}

		p = i
	}

	fmt.Println("Number of increases: ", c)

	return nil
}

func Run() error {
	file, err := os.Open("./input/day1.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	ints := []int{}

	s := bufio.NewScanner(file)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			return fmt.Errorf("Failed to parse depth")
		}

		ints = append(ints, i)
	}

	h := 3
	t := 0

	ps := 0
	c := -1

	for h <= len(ints) {
		s := sumSlice(ints[t:h])
		if s > ps {
			c++
		}

		ps = s

		h++
		t++
	}

	fmt.Println("Number of increases: ", c)

	return nil
}

func sumSlice(is []int) int {
	s := 0
	for _, i := range is {
		s += i
	}

	return s
}
