package main

import (
	"aoc/day1"
	"aoc/day2"
	"aoc/day3"
	"aoc/day4"
	"aoc/day5"
	"aoc/day6"
	"aoc/day7"
	"aoc/day8"
	"aoc/day9"
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
		err = day1.Run()
	case 2:
		err = day2.Run()
	case 3:
		err = day3.Run()
	case 4:
		err = day4.Run()
	case 5:
		err = day5.Run()
	case 6:
		err = day6.Run()
	case 7:
		err = day7.Run()
	case 8:
		err = day8.Run()
	case 9:
		err = day9.Run()
	}

	if err != nil {
		log.Fatal(err)
	}
}
