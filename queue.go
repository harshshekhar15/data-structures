package queue

import (
	"fmt"
)

type Node struct {
	val      int
	nextNode *Node
}

type Queue struct {
	rootNode *Node
}

// newNode return a new instance of Node
func newNode(value int) *Node {
	return &Node{
		val:      value,
		nextNode: nil,
	}
}

// New returns a new instance of Queue
func New() *Queue {
	return &Queue{
		rootNode: nil,
	}
}

// Push adds a new node to the queue
func (q *Queue) Push(value int) {
	item := newNode(value)

	// Pushing element for the first time
	if q.rootNode == nil {
		q.rootNode = item
	} else {
		node := q.rootNode
		for node.nextNode != nil {
			node = node.nextNode
		}
		node.nextNode = item
	}
}

// Pop deletes a node from the queue
func (q *Queue) Pop() {
	if q.rootNode == nil {
		fmt.Println("Nothing to pop. Queue empty!")
	} else {
		q.rootNode = q.rootNode.nextNode
	}
}

// Print prints the queue
func (q *Queue) Print() {
	if q.rootNode == nil {
		fmt.Println("Queue is empty!!")
	} else {
		iterator := q.rootNode
		for iterator.nextNode != nil {
			fmt.Printf(" %d ->", iterator.val)
			iterator = iterator.nextNode
		}
		fmt.Printf(" %d ->", iterator.val)
		fmt.Println()
	}
}
