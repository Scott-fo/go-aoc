package day3

import (
	"bufio"
	"fmt"
	"os"
)

func partOne() error {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	bits := make(map[int]map[int]int)
	for s.Scan() {
		for idx, r := range s.Text() {
			i := int(r - '0')

			_, ok := bits[idx]
			if !ok {
				bits[idx] = map[int]int{0: 0, 1: 0}
			}

			bits[idx][i] += 1
		}
	}

	fmt.Println(bits)

	gr := 0
	er := 0

	for bit := len(bits) - 1; bit >= 0; bit-- {
		if bits[bit][0] < bits[bit][1] {
			gr |= 1 << (len(bits) - 1 - bit)
		} else {
			er |= 1 << (len(bits) - 1 - bit)
		}
	}

	pc := er * gr

	fmt.Println("Gamme rate: ", gr)
	fmt.Println("Epsilon rate: ", er)
	fmt.Println("Power consumption: ", pc)

	return nil
}

func getBitCount(bits *[][]int) map[int]map[int]int {
	bc := make(map[int]map[int]int)
	for _, bs := range *bits {
		for i, b := range bs {
			_, ok := bc[i]
			if !ok {
				bc[i] = map[int]int{0: 0, 1: 0}
			}

			bc[i][b] += 1
		}
	}

	return bc
}

func filterBits(bits *[][]int, criteria string) {
	for x := 0; x < len((*bits)[0]); x++ {
		bc := getBitCount(bits)

		var mc int
		var lc int

		if bc[x][0] > bc[x][1] {
			mc = 0
		} else {
			mc = 1
		}

		if bc[x][0] <= bc[x][1] {
			lc = 0
		} else {
			lc = 1
		}

		is := []int{}

		for y := 0; y < len(*bits); y++ {
			if criteria == "MC" && (*bits)[y][x] == mc {
				is = append(is, y)
			}

			if criteria == "LC" && (*bits)[y][x] == lc {
				is = append(is, y)
			}
		}

		fb := make([][]int, 0, len(is))
		for _, i := range is {
			fb = append(fb, (*bits)[i])
		}

		*bits = fb

		if len(*bits) == 1 {
			return
		}
	}
}

func binaryArrayToDecimal(a []int) int {
	out := 0
	for bit := len(a) - 1; bit >= 0; bit-- {
		if a[bit] == 1 {
			out |= 1 << (len(a) - 1 - bit)
		}
	}

	return out
}

func Run() error {
	file, err := os.Open("./input/day3.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	bits := [][]int{}
	for s.Scan() {
		br := []int{}
		for _, r := range s.Text() {
			i := int(r - '0')
			br = append(br, i)
		}

		bits = append(bits, br)
	}

	og := make([][]int, len(bits))
	copy(og, bits)
	filterBits(&og, "MC")

	cs := make([][]int, len(bits))
	copy(cs, bits)
	filterBits(&cs, "LC")

	fmt.Println("Oxygen Generator: ", og)
	fmt.Println("CS Scrubber: ", cs)

	ogV := binaryArrayToDecimal(og[0])
	csV := binaryArrayToDecimal(cs[0])

	fmt.Println("Oxygen Generator Decimal: ", ogV)
	fmt.Println("CS Scrubber Decimal: ", csV)

	p := ogV * csV
	fmt.Println("Product: ", p)

	return nil
}
