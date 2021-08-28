package main

import (
	"github.com/harshshekhar15/queue/pkg/queue"
)

func main() {
	q := queue.New()
	q.Print()
	q.Push(10)
	q.Print()
	q.Push(2)
	q.Print()
	q.Pop()
	q.Print()
	q.Pop()
	q.Print()
	q.Pop()
	q.Pop()
}
