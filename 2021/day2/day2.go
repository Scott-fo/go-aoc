package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Submarine struct {
	Horizontal int
	Depth      int
	Aim        int
}

func (s *Submarine) move(i Instruction) {
	switch i.Direction {
	case UP:
		s.Aim -= i.Steps
	case DOWN:
		s.Aim += i.Steps
	case FORWARD:
		s.Horizontal += i.Steps
		s.Depth += s.Aim * i.Steps
	}
}

type Direction string

const (
	FORWARD Direction = "forward"
	DOWN    Direction = "down"
	UP      Direction = "up"
)

func direction(s string) (Direction, error) {
	switch s {
	case "forward":
		return FORWARD, nil
	case "down":
		return DOWN, nil
	case "up":
		return UP, nil
	}

	return "", fmt.Errorf("Unexpected direction")
}

type Instruction struct {
	Direction Direction
	Steps     int
}

func parseInstruction(s string) (Instruction, error) {
	sp := strings.Split(s, " ")
	d, err := direction(sp[0])
	if err != nil {
		return Instruction{}, err
	}

	st, err := strconv.Atoi(sp[1])
	if err != nil {
		return Instruction{}, fmt.Errorf("Failed to parse steps")
	}

	return Instruction{
		Direction: d,
		Steps:     st,
	}, nil
}

func Run() error {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	s := bufio.NewScanner(file)

	sub := Submarine{
		Horizontal: 0,
		Depth:      0,
	}

	for s.Scan() {
		i, err := parseInstruction(s.Text())
		if err != nil {
			return err
		}

		sub.move(i)
	}

	prod := sub.Depth * sub.Horizontal
	fmt.Println("Final horiztonal and depth product: ", prod)

	return nil
}
