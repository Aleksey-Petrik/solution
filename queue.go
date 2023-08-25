package main

type Queue struct {
	symbols []string
}

func NewQueue() Queue {
	return Queue{
		symbols: []string{},
	}
}

func (q *Queue) Add(symbol string) {
	q.symbols = append(q.symbols, symbol)
}

func (q *Queue) Remove() string {
	symbol := q.symbols[len(q.symbols)-1]
	q.symbols = q.symbols[:len(q.symbols)-1]
	return symbol
}

func (q *Queue) Size() int {
	return len(q.symbols)
}
