package main

import (
	"math/rand"
)

func step1(index int, cells *CellArray, done signal) {
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			next := cellStep(index)
			if index != next {
				cells[index].count--
				cells[next].count++
				index = next
			}
		}
	}
}

func step2(index int, cells *CellArray, done signal) {
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			next := cellStep(index)
			if index != next {
				cells[index].mutex.Lock()
				cells[next].mutex.Lock()

				cells[index].count--
				cells[next].count++

				cells[next].mutex.Unlock()
				cells[index].mutex.Unlock()

				index = next
			}
		}
	}
}

func cellStep(i int) int {
	if rand.Float64() > p {
		i = i + 1
	} else {
		i = i - 1
	}

	return clamp(i, 0, n-1)
}
