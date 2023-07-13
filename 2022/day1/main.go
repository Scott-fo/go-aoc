package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	var elfMap []int
	elfVal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			elfMap = append(elfMap, elfVal)
			elfVal = 0
			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		elfVal += val
	}

	if elfVal > 0 {
		elfMap = append(elfMap, elfVal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sort.Slice(elfMap, func(i, j int) bool {
		return elfMap[i] > elfMap[j]
	})

	fmt.Println(elfMap[0])
	sum := 0
	for _, val := range elfMap[0:3] {
		sum += val
	}

	fmt.Println(sum)
}
