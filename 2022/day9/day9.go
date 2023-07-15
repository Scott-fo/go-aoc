package day9

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Direction string

const (
	L Direction = "L"
	R Direction = "R"
	U Direction = "U"
	D Direction = "D"
)

func parseDirection(s string) Direction {
	switch s {
	case "L":
		return L
	case "R":
		return R
	case "U":
		return U
	case "D":
		return D
	}

	return ""
}

type Instruction struct {
	Direction Direction
	Steps     int
}

type Node struct {
	X int
	Y int
}

func (n *Node) Step(d Direction) {
	switch d {
	case L:
		n.X -= 1
	case R:
		n.X += 1
	case U:
		n.Y += 1
	case D:
		n.Y -= 1
	}
}

func (n *Node) StepTowards(target Node) {
	if target.X > n.X {
		n.Step(R)
	} else if target.X < n.X {
		n.Step(L)
	} else if target.Y > n.Y {
		n.Step(U)
	} else if target.Y < n.Y {
		n.Step(D)
	}
}

func (n *Node) StepDiagonalTowards(target Node) {
	if target.X > n.X {
		n.Step(R)
	} else {
		n.Step(L)
	}

	if target.Y > n.Y {
		n.Step(U)
	} else {
		n.Step(D)
	}
}

func parseInstruction(s string) Instruction {
	sp := strings.Split(s, " ")
	steps, _ := strconv.Atoi(sp[1])
	direction := parseDirection(sp[0])

	return Instruction{
		Direction: direction,
		Steps:     steps,
	}
}

func adjacent(head Node, tail Node) bool {
	if abs(head.X-tail.X) <= 1 && abs(head.Y-tail.Y) <= 1 {
		return true
	}

	return false
}

func diagonal(head Node, tail Node) bool {
	if head.X != tail.X && head.Y != tail.Y {
		return true
	}

	return false

}

func abs(i int) int {
	if i < 0 {
		return -i
	}

	return i
}

func Run() error {
	file, err := os.Open("./input/day9.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	nodes := []Node{}
	for i := 0; i < 10; i++ {
		nodes = append(nodes, Node{
			X: 0,
			Y: 0,
		})
	}

	seen := make(map[Node]bool)

	for s.Scan() {
		inst := parseInstruction(s.Text())
		for i := 0; i < inst.Steps; i++ {
			nodes[0].Step(inst.Direction)
			for n := 1; n < len(nodes); n++ {
				if !adjacent(nodes[n-1], nodes[n]) {
					if diagonal(nodes[n-1], nodes[n]) {
						nodes[n].StepDiagonalTowards(nodes[n-1])
					} else {
						nodes[n].StepTowards(nodes[n-1])
					}
				}
			}
			if !seen[nodes[len(nodes)-1]] {
				seen[nodes[len(nodes)-1]] = true
			}
		}
	}

	fmt.Println("Tail visited: ", len(seen))
	return nil
}

func partOne() error {
	file, err := os.Open("./input/day9.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	s := bufio.NewScanner(file)

	head := Node{X: 0, Y: 0}
	tail := Node{X: 0, Y: 0}

	seen := make(map[Node]bool)

	for s.Scan() {
		inst := parseInstruction(s.Text())
		for i := 0; i < inst.Steps; i++ {
			head.Step(inst.Direction)
			if !adjacent(head, tail) {
				if diagonal(head, tail) {
					tail.StepDiagonalTowards(head)
				} else {
					tail.Step(inst.Direction)
				}
			}

			if !seen[tail] {
				seen[tail] = true
			}
		}
	}

	fmt.Println("Tail visited: ", len(seen))
	return nil
}
