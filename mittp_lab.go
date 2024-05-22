package main

import (
	"fmt"
	"sync"
	"time"
)

const n = 10
const k = 5
const p = 0.5
const t = 10
const period = 5
const synced = false

type Number interface {
	int | float32 | float64
}

type Cell struct {
	mutex sync.Mutex
	count int
}

type CellArray [n]Cell
type signal chan interface{}

func clamp[T Number](x, lo, hi T) T {
	return min(max(x, lo), hi)
}

func printCells(cells *CellArray) {
	fmt.Printf("Current cells state: %v [", countCells(cells))
	for i := 0; i < n; i++ {
		fmt.Printf("%v", cells[i].count)
		if i != n-1 {
			fmt.Printf(", ")
		}
	}
	fmt.Printf("]")
	fmt.Println()
}

func countCells(cells *CellArray) int {
	sum := 0
	for i := 0; i < n; i++ {
		sum += cells[i].count
	}
	return sum
}

func main() {
	var cells CellArray

	// init
	for i := 0; i < n; i++ {
		cells[i] = Cell{
			count: 0,
			mutex: sync.Mutex{},
		}
	}

	cells[0].count = k

	printCells(&cells)

	done := make(signal)
	// start k goroutines for each of the k atoms
	for i := 0; i < k; i++ {
		if synced {
			go step2(0, &cells, done)
		} else {
			go step1(0, &cells, done)
		}
	}

	// log the state of the cells every `period` seconds
	go (func() {
	loop:
		for {
			time.Sleep(period * time.Second)
			select {
			case <-done:
				break loop
			default:
				for i := 0; i < n; i++ {
					cells[i].mutex.Lock()
				}

				printCells(&cells)

				for i := 0; i < n; i++ {
					cells[i].mutex.Unlock()
				}
			}
		}
	})()

	time.Sleep(t * time.Second)
	close(done)

	// check the final state of the cells
	fmt.Printf("Total cell count after simulation: %v", countCells(&cells))
	fmt.Println()
}
