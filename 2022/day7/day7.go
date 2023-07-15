package day7

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const TOTAL_DISK_SPACE = 70000000
const REQUIRED_SPACE = 30000000

func Run() error {
	file, err := os.Open("./input/day7.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	total := 0
	stack := []int{0}
	dirs := []int{}

	for s.Scan() {
		t := s.Text()
		if t == "$ cd /" || t == "$ ls" {
			continue
		}

		if strings.HasPrefix(t, "$ cd ") {
			dir := t[5:]
			if dir == ".." {
				var a int
				a, stack = stack[len(stack)-1], stack[:len(stack)-1]

				if a <= 100000 {
					total += a
				}

				stack[len(stack)-1] += a
				dirs = append(dirs, a)
			} else {
				stack = append(stack, 0)
			}

			continue
		}

		meta := strings.Split(t, " ")

		if meta[0] == "dir" {
			continue
		}

		size := meta[0]
		if size == "" {
			return fmt.Errorf("Didn't get size of file as expected")
		}

		sizeInt, err := strconv.Atoi(size)
		if err != nil {
			return fmt.Errorf("Expected file size to be int")
		}

		stack[len(stack)-1] += sizeInt
	}

	for len(stack) > 1 {
		var a int
		a, stack = stack[len(stack)-1], stack[:len(stack)-1]
		dirs = append(dirs, a)

		if a <= 100000 {
			total += a
		}

		stack[len(stack)-1] += a
	}

	dirs = append(dirs, stack[0])

	fmt.Println("Sum under 100k: ", total)

	fs := TOTAL_DISK_SPACE - stack[0]
	tf := REQUIRED_SPACE - fs

	min := math.MaxInt
	for _, a := range dirs {
		if a >= tf && min > a {
			min = a
		}
	}

	fmt.Println("Delete dir size: ", min)

	return nil
}
