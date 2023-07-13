package main

import (
	"bufio"
	"bytes"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	heap := *h
	length := len(heap)
	out := heap[length-1]
	*h = heap[0 : length-1]
	return out
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	h := &IntHeap{}
	heap.Init(h)

	elfVal := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := bytes.TrimSpace(scanner.Bytes())
		if len(line) == 0 {
			if h.Len() < 3 {
				heap.Push(h, elfVal)
			} else if elfVal > (*h)[0] {
				heap.Pop(h)
				heap.Push(h, elfVal)
			}
			elfVal = 0
			continue
		}

		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		elfVal += val
	}

	if h.Len() < 3 {
		heap.Push(h, elfVal)
	} else if elfVal > (*h)[0] {
		heap.Pop(h)
		heap.Push(h, elfVal)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	max := math.MinInt
	for h.Len() > 0 {
		curr := heap.Pop(h).(int)
		if curr > max {
			max = curr
		}
		sum += curr
	}

	fmt.Println(max)
	fmt.Println(sum)
}
