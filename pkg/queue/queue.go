package queue

import (
	"fmt"

	"github.com/harshshekhar15/queue/internal/node"
)

type Queue struct {
	rootNode *node.Node
}

// New returns a new instance of Queue
func New() *Queue {
	return &Queue{
		rootNode: nil,
	}
}

// Push adds a new node to the queue
func (q *Queue) Push(value int) {
	item := node.New(value)

	// Pushing element for the first time
	if q.rootNode == nil {
		q.rootNode = item
	} else {
		node := q.rootNode
		for node.NextNode != nil {
			node = node.NextNode
		}
		node.NextNode = item
	}
}

// Pop deletes a node from the queue
func (q *Queue) Pop() {
	if q.rootNode == nil {
		fmt.Println("Nothing to pop. Queue empty!")
	} else {
		q.rootNode = q.rootNode.NextNode
	}
}

// Print prints the queue
func (q *Queue) Print() {
	if q.rootNode == nil {
		fmt.Println("Queue is empty!!")
	} else {
		iterator := q.rootNode
		for iterator.NextNode != nil {
			fmt.Printf(" %d ->", iterator.Val)
			iterator = iterator.NextNode
		}
		fmt.Printf(" %d ->", iterator.Val)
		fmt.Println()
	}
}
