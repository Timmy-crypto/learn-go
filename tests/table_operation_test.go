package tests

import (
	"fmt"
	"sync"
)

type TableUpdate struct {
	RawNumber     int
	ColNumber     int
	SumValue      int
	Values        [][]int
	SumValueCache map[string]int
	cacheMutex    sync.RWMutex
	valuesMutex   sync.RWMutex
}

func (t *TableUpdate) init(rawNumber int, colNumber int) {
	if rawNumber < 0 || t.ColNumber < 0 {
		panic("raw or col number is negative")
	}
	t.RawNumber = rawNumber
	t.ColNumber = colNumber
	t.Values = make([][]int, rawNumber)
	for i := 0; i < t.RawNumber; i++ {
		t.Values[i] = make([]int, colNumber)
	}
	t.SumValueCache = make(map[string]int)
}

func (t *TableUpdate) set(x int, y int, value int) {
	if x < 0 || y < 0 {
		panic("x or y is negative")
	}
	if x >= t.ColNumber || y >= t.RawNumber {
		panic("table value error")
	}
	var originValue int
	t.valuesMutex.Lock()
	originValue = t.Values[y][x]
	t.Values[y][x] = value
	t.valuesMutex.Unlock()

	changeValue := value - originValue
	t.cacheMutex.Lock()
	//update sumValueCache
	for i := y; i < t.RawNumber; i++ {
		for j := x; j < t.ColNumber; j++ {
			key := fmt.Sprintf("%d:%d", y, x)
			t.SumValueCache[key] += changeValue
		}
	}
	t.cacheMutex.Unlock()
}

func (t *TableUpdate) sum(x int, y int) int {
	if x < 0 || y < 0 {
		panic("x or y is negative")
	}
	if x >= t.ColNumber || y >= t.RawNumber {
		panic("table value error")
	}
	key := fmt.Sprintf("%d:%d", y, x)
	t.cacheMutex.RLock()
	t.SumValue = t.SumValueCache[key]
	t.cacheMutex.RUnlock()

	return t.SumValue
}
