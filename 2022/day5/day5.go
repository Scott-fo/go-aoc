package day5

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	quantity    int
	origin      int
	destination int
}

func partOne() error {
	data, err := os.ReadFile("./input/day5.txt")
	if err != nil {
		return err
	}

	sections := bytes.Split(data, []byte("\n\n"))
	stacks := parseStacks(string(sections[0]))

	lines := strings.Split(string(sections[1]), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		inst, err := parseInstruction(string(line))
		if err != nil {
			return err
		}

		s := stacks[inst.origin]

		for k := 0; k < inst.quantity; k++ {
			var crate rune
			crate, s = s[len(s)-1], s[:len(s)-1]
			stacks[inst.destination] = append(stacks[inst.destination], crate)
		}

		stacks[inst.origin] = s
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
	fmt.Println()

	return nil
}

func partTwo() error {
	data, err := os.ReadFile("./input/day5.txt")
	if err != nil {
		return err
	}

	sections := bytes.Split(data, []byte("\n\n"))
	stacks := parseStacks(string(sections[0]))

	lines := strings.Split(string(sections[1]), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		inst, err := parseInstruction(string(line))
		if err != nil {
			return err
		}

		s := stacks[inst.origin]

		if inst.quantity == 1 {
			var crate rune
			crate, s = s[len(s)-1], s[:len(s)-1]
			stacks[inst.destination] = append(stacks[inst.destination], crate)
		} else {
			movedCrates := s[len(s)-inst.quantity:]
			s = s[:len(s)-inst.quantity]
			stacks[inst.destination] = append(stacks[inst.destination], movedCrates...)
		}
		stacks[inst.origin] = s
	}

	for _, stack := range stacks {
		fmt.Print(string(stack[len(stack)-1]))
	}
	fmt.Println()

	return nil
}

func Run() error {
	return partTwo()
}

func parseStacks(input string) [][]rune {
	lines := strings.Split(input, "\n")

	numStacks := len(strings.Fields(lines[len(lines)-1]))
	stacks := make([][]rune, numStacks)

	for i := len(lines) - 2; i >= 0; i-- {
		line := lines[i]
		for k := 1; k <= len(line); k += 4 {
			if line[k] != ' ' {
				stackIdx := (k - 1) / 4
				crate := rune(line[k])
				stacks[stackIdx] = append(stacks[stackIdx], crate)
			}
		}
	}

	printStacks(stacks)

	return stacks
}

func parseInstruction(input string) (Instruction, error) {
	sl := strings.Split(input, " ")
	quantity, err := strconv.Atoi(sl[1])
	if err != nil {
		return Instruction{}, fmt.Errorf("Failed to parse quantity")
	}

	origin, err := strconv.Atoi(sl[3])
	if err != nil {
		return Instruction{}, fmt.Errorf("Failed to parse origin")
	}

	destination, err := strconv.Atoi(sl[5])
	if err != nil {
		return Instruction{}, fmt.Errorf("Failed to parse destination")
	}

	return Instruction{
		quantity:    quantity,
		origin:      origin - 1,
		destination: destination - 1,
	}, nil
}

func printStacks(stacks [][]rune) {
	for i, stack := range stacks {
		fmt.Printf("Stack %d: ", i+1)
		for _, crate := range stack {
			fmt.Printf("%c ", crate)
		}
		fmt.Println()
	}
}
