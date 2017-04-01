package main

import "fmt"

type PriorityQueue struct {
	queue []int
	max int
}

// max heap
func (pq *PriorityQueue) sink(k int) {
	for ; 2*k < pq.max ; {
		leftC := 2*k
		rightC := 2*k + 1
		maxC := leftC
		
		// figure out largest child
		if leftC < pq.max && pq.less(leftC, rightC) {
			maxC = rightC
		}

		// check property
		if !pq.less(k, maxC) {
			break
		}

		// swap out stuff
		pq.exch(k, maxC)

		// move to next item.
		k = maxC
	}
}

func (pq *PriorityQueue) swim(k int) {
	for ; k>1 ; {
		if pq.less(k/2, k) {
			pq.exch(k/2, k)
		}
		k = k/2
	}
}

func (pq *PriorityQueue) add(priority int) {
	pq.max++
	pq.queue[pq.max] = priority
	pq.swim(pq.max)
}

func (pq *PriorityQueue) remove() (int) {
	retVal := pq.queue[1]

	pq.exch(1, pq.max)
	pq.queue[pq.max] = -1
	pq.max--

	pq.sink(1)

	return retVal
}

// return true is val1 < val2
func (pq *PriorityQueue) less(val1, val2 int) (bool) {
	if pq.queue[val1] < pq.queue[val2] {
		return true
	} else {
		return false
	}
}

func (pq *PriorityQueue) exch(val1, val2 int) {
	temp := pq.queue[val1]
	pq.queue[val1] = pq.queue[val2]
	pq.queue[val2] = temp
}



func main() {
	var pq PriorityQueue
	pq.queue = make([]int,8)
	pq.add(1)
	pq.add(2)
	pq.add(3)
	pq.add(4)
	pq.add(5)
	pq.add(6)
	pq.add(7)
	fmt.Println(pq.queue)
	r:=pq.remove()
	r=pq.remove()
	pq.add(6)
	fmt.Println(pq.queue, r)
}



/*

http://algs4.cs.princeton.edu/24pq/

*/