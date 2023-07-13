package main

import (
	"aoc/day1"
	"aoc/day2"
	"log"
	"os"
	"strconv"
)

func main() {
	day := os.Args[1]
	i, err := strconv.Atoi(day)
	if err != nil {
		log.Fatal("Incorrect arguments. Pass day number")
	}

	switch i {
	case 1:
		day1.Run()
	case 2:
		day2.Run()
	}
}
