package iterator

import "fmt"

// Iterator defines traversal methods
type Iterator interface {
	HasNext() bool
	Next() int
}

// Concrete Aggregate
type IntSlice struct {
	items []int
}

func NewIntSlice(items []int) *IntSlice {
	return &IntSlice{items}
}

func (s *IntSlice) Iterator() Iterator {
	return &IntSliceIterator{
		slice: s,
		index: 0,
	}
}

// Concrete Iterator
type IntSliceIterator struct {
	slice *IntSlice
	index int
}

func (it *IntSliceIterator) HasNext() bool {
	return it.index < len(it.slice.items)
}

func (it *IntSliceIterator) Next() int {
	if !it.HasNext() {
		panic("No more elements")
	}
	val := it.slice.items[it.index]
	it.index++
	return val
}

// Client
func RunArrayIterator() {
	numbers := NewIntSlice([]int{2, 4, 6, 8, 10})
	it := numbers.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
