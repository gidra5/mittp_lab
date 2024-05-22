package main

import (
	"math/rand"
	"sync"
)

type Change struct {
	from int
	to   int
}
type Collector chan Change

type StepFunc func(from int, to int)

func cellStep(i int) int {
	if rand.Float64() > p {
		i = i + 1
	} else {
		i = i - 1
	}

	return clamp(i, 0, n-1)
}

func simulate(index int, done signal, step StepFunc) {
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			next := cellStep(index)
			if index != next {
				step(index, next)
				index = next
			}
		}
	}
}

func step1(index int, next int, cells *CellArray) {
	cells[index].count--
	cells[next].count++
}

func step2(index int, next int, cells *CellArray) {
	cells[index].mutex.Lock()
	cells[index].count--
	cells[index].mutex.Unlock()

	cells[next].mutex.Lock()
	cells[next].count++
	cells[next].mutex.Unlock()
}

func step3(index int, next int, mutex *sync.Mutex, cells *CellArray) {
	mutex.Lock()

	cells[index].count--
	cells[next].count++

	mutex.Unlock()
}

func step4(index int, next int, cells *CellArray) {
	cells[index].mutex.Lock()
	cells[next].mutex.Lock()

	cells[index].count--
	cells[next].count++

	cells[index].mutex.Unlock()
	cells[next].mutex.Unlock()
}

func step5(index int, next int, cells *CellArray) {
	cells[index].mutex.Lock()

	cells[index].count--
	cells[next].count++

	cells[index].mutex.Unlock()
}

func step6(index int, next int, collector Collector) {
	collector <- Change{index, next}
}
