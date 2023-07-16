package main

import (
	"2021/day1"
	"2021/day2"
	"2021/day3"
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
	}

	if err != nil {
		log.Fatal(err)
	}
}
