package day2

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Move string

const (
	Rock     Move = "rock"
	Paper    Move = "paper"
	Scissors Move = "scissors"
)

type Result string

const (
	Win  Result = "win"
	Lose Result = "lose"
	Draw Result = "draw"
)

var winMap = map[Move]Move{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var loseMap = map[Move]Move{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

func parseMove(s string) (Move, error) {
	switch s {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	}

	return "", errors.New("invalid move")
}

func parseStrategy(s string) (Result, error) {
	switch s {
	case "X":
		return Lose, nil
	case "Y":
		return Draw, nil
	case "Z":
		return Win, nil
	}

	return "", errors.New("unexpected strategy")
}

func getMoveScore(m Move) (int, error) {
	switch m {
	case Rock:
		return 1, nil
	case Paper:
		return 2, nil
	case Scissors:
		return 3, nil
	}

	return 0, errors.New("invalid move")
}

func getResultScore(r Result) (int, error) {
	switch r {
	case Win:
		return 6, nil
	case Lose:
		return 0, nil
	case Draw:
		return 3, nil
	}

	return 0, errors.New("invalid result")
}

func getResult(y, o Move) Result {
	switch {
	case y == o:
		return Draw
	case (y == Rock && o == Scissors) || (y == Paper && o == Rock) || (y == Scissors && o == Paper):
		return Win
	default:
		return Lose
	}
}

func getMoveToPlay(o Move, r Result) (Move, error) {
	switch r {
	case Draw:
		return o, nil
	case Win:
		return winMap[o], nil
	case Lose:
		return loseMap[o], nil
	}

	return "", errors.New("unexpected strategy")
}

func Run() error {
	file, err := os.Open("./input/day2.txt")
	if err != nil {
		return fmt.Errorf("Failed to open input")
	}

	defer file.Close()

	score := 0

	s := bufio.NewScanner(file)
	for s.Scan() {
		r := strings.Split(s.Text(), " ")

		o, err := parseMove(r[0])
		if err != nil {
			return err
		}

		st, err := parseStrategy(r[1])
		if err != nil {
			return err
		}

		m, err := getMoveToPlay(o, st)
		if err != nil {
			return err
		}

		rs, err := getResultScore(st)
		if err != nil {
			return err
		}

		ms, err := getMoveScore(m)
		if err != nil {
			return err
		}

		score += rs + ms
	}

	fmt.Println("Score: ", score)

	return nil
}
