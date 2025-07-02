package iterator

import (
	"bufio"
	"fmt"
	"os"
)

// Iterator interface
type LineIterator interface {
	HasNext() bool
	Next() (string, error)
}

// Concrete Iterator
type FileLineIterator struct {
	scanner *bufio.Scanner
	hasNext bool
}

func NewFileLineIterator(path string) (*FileLineIterator, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	sc := bufio.NewScanner(f)
	has := sc.Scan()
	return &FileLineIterator{scanner: sc, hasNext: has}, nil
}

func (it *FileLineIterator) HasNext() bool {
	return it.hasNext
}

func (it *FileLineIterator) Next() (string, error) {
	if !it.hasNext {
		return "", fmt.Errorf("no more lines")
	}
	line := it.scanner.Text()
	it.hasNext = it.scanner.Scan()
	return line, it.scanner.Err()
}

// Client
func RunFileIterator() {
	it, err := NewFileLineIterator("data.txt")
	if err != nil {
		panic(err)
	}
	for it.HasNext() {
		line, err := it.Next()
		if err != nil {
			panic(err)
		}
		fmt.Println(line)
	}
}
