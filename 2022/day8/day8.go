package day8

import (
	"bufio"
	"fmt"
	"os"
)

func Run() error {
	file, err := os.Open("./input/day8.txt")
	if err != nil {
		return fmt.Errorf("Failed to read input")
	}

	defer file.Close()

	var v [][]int

	s := bufio.NewScanner(file)
	for s.Scan() {
		iv := []int{}
		for _, r := range s.Text() {
			i := int(r - '0')
			iv = append(iv, i)
		}
		v = append(v, iv)

	}

	// trees := (2 * len(v)) + 2*(len(v[0])-2)
	highScore := 0

	for y := 0; y < len(v); y++ {
		for x := 0; x < len(v[0]); x++ {
			score := checkVisible(&v, x, y)
			if score > highScore {
				highScore = score
			}
		}
	}

	fmt.Println("Score : ", highScore)
	return nil
}

func checkVisible(v *[][]int, x int, y int) int {
	t := (*v)[y][x]
	score := 1

	ud, _ := checkUp(v, x, y-1, t)
	score *= ud + 1

	ld, _ := checkLeft(v, x-1, y, t)
	score *= ld + 1

	rd, _ := checkRight(v, x+1, y, t)
	score *= rd + 1

	dd, _ := checkDown(v, x, y+1, t)
	score *= dd + 1

	return score
}

func checkUp(v *[][]int, x int, y int, t int) (int, bool) {
	for i := y; i >= 0; i-- {
		h := (*v)[i][x]
		if h >= t {
			return y - i, false
		}
	}

	return y, true
}

func checkDown(v *[][]int, x int, y int, t int) (int, bool) {
	for i := y; i <= len(*v)-1; i++ {
		h := (*v)[i][x]
		if h >= t {
			return i - y, false
		}
	}

	return len(*v) - y - 1, true
}

func checkLeft(v *[][]int, x int, y int, t int) (int, bool) {
	for i := x; i >= 0; i-- {
		h := (*v)[y][i]
		if h >= t {
			return x - i, false
		}
	}

	return x, true
}

func checkRight(v *[][]int, x int, y int, t int) (int, bool) {
	for i := x; i <= len((*v)[0])-1; i++ {
		h := (*v)[y][i]
		if h >= t {
			return i - x, false
		}
	}

	return len((*v)[0]) - x - 1, true
}
