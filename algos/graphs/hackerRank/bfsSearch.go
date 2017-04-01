package main

// 10:49a
import "fmt"

// ==== linked list
type NodeL struct {
	next *NodeL
	val *NodeG
}

func newNodeL(val *NodeG) (*NodeL) {
	n := NodeL{nil, val}
	return &n
}

// 10:54ap
// queue takes 2 pointers...
// enter at tail
// pop off front

type Queue struct {
	head *NodeL
	tail *NodeL
}

func (q *Queue) enqueue(val *NodeG) {
	n := newNodeL(val)
		
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		(*q.tail).next = n
		q.tail = n
	}
}

func (q *Queue) dequeue() (val *NodeG) {
	retNode := q.head

	// move the head pointer.
	if (*q.head).next != nil {
		q.head = (*q.head).next
	}

	// return first item of queue.
	(*retNode).next = nil
	return (*retNode).val
}

 
// ==== graph

type NodeG struct {
	vertex int
	next *NodeG
}

type AdjList struct {
	vertexList []*NodeG
}

func newNodeG(val int, next *NodeG) (*NodeG) {
	n := NodeG{val, next}
	return &n
}

// create adj list...
func createAdjListNode(val int, adjEdge []int) (*NodeG) {
	root := newNodeG(val, nil)
	if len(adjEdge) == 0 {
		return root
	}
	
	result := newNodeG(val, root)
	tail := result
	for i:=0 ; i<len(adjEdge) ; i++ {
		n := newNodeG(adjEdge[i], nil)
		tail.next = n
		tail = n
	}

	return result
}




func main() {
	var q Queue
	q.enqueue(nil)
	fmt.Println(q.head)

	v := createAdjListNode(1, []int{2,3})
	fmt.Println((*v).next)
}


/*
graphs..
use list
and do a bfs search with cycles..
*/