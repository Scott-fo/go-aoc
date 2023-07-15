package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CPU struct {
	X int
}

type Instruction string

const (
	ADD_X Instruction = "add_x"
	NO_OP Instruction = "no_op"
)

func parseInstruction(s string) (Instruction, int) {
	if strings.HasPrefix(s, "addx") {
		si := strings.Split(s, " ")
		i, _ := strconv.Atoi(si[1])

		return ADD_X, i
	}

	return NO_OP, 0
}

type Process struct {
	Instruction Instruction
	Value       int
}

func NewProcess(s string) Process {
	i, v := parseInstruction(s)
	return Process{
		Instruction: i,
		Value:       v,
	}
}

type Queue []Process

func (q *Queue) Dequeue() Process {
	i := (*q)[0]
	*q = (*q)[1:]

	return i
}

func (q *Queue) Enqueue(p Process) {
	*q = append(*q, p)
}

func Run() error {
	file, err := os.Open("./input/day10.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	var q Queue

	s := bufio.NewScanner(file)
	for s.Scan() {
		p := NewProcess(s.Text())
		if p.Instruction == ADD_X {
			q.Enqueue(NewProcess("noop"))
		}

		q.Enqueue(p)
	}

	cycle := 0
	c := CPU{X: 1}
	signalStrengthSum := 0

	for len(q) > 0 {
		drawCrt(c.X, cycle)

		cycle++

		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			signalStrengthSum += cycle * c.X
		}

		p := q.Dequeue()
		c.X += p.Value
	}

	fmt.Println()
	fmt.Println("SignalStrengthSUm: ", signalStrengthSum)

	return nil
}

func drawCrt(x int, cycle int) {
	position := cycle % 40

	if position == 0 {
		fmt.Println()
	}

	if x == position || x-1 == position || x+1 == position {
		fmt.Print("#")
	} else {
		fmt.Print(".")
	}
}
