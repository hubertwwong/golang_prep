package main

import "fmt"

type PQueue struct {
	queue []int
	max int
	pos int
}

// using +3 for a few things.
// +1 for the start at 1 for a pq.
// +1 for extra slot.
func (pq *PQueue) init(max int) {
	pq.queue = make([]int, max+3)
	pq.max = max
	pq.pos = 1
}

func (pq *PQueue) getMin() int {
	return pq.queue[1]
}

// insert item at the end of list.
// bubble the value up to respect the heap property.
func (pq *PQueue) insert(val int) {
	if pq.pos <= pq.max {
		pq.queue[pq.pos] = val
		
		for i := pq.pos ; i>1 ; i/=2 {
			if pq.queue[i] > pq.queue[i/2] {
				pq.queue[i], pq.queue[i/2] = pq.queue[i/2], pq.queue[i]
			}
		}
		
		pq.pos++
	} else {
		// going to use the extra slot.
		pq.queue[pq.pos] = val

		for i := pq.pos ; i>1 ; i/=2 {
			if pq.queue[i] > pq.queue[i/2] {
				pq.queue[i], pq.queue[i/2] = pq.queue[i/2], pq.queue[i]
			}
		}

		// remove ever item except the last item.
		tmp := make([]int, 0)
		for i := 0 ; i < pq.pos ; i++ {
			tmp = append(tmp, pq.removeMax())
		}

		// insert again.
		for k := 0 ; k < len(tmp) ; k++ {
			pq.queue[pq.pos] = tmp[k]
		
			for i := pq.pos ; i>1 ; i/=2 {
				if pq.queue[i] > pq.queue[i/2] {
					pq.queue[i], pq.queue[i/2] = pq.queue[i/2], pq.queue[i]
				}
			}
			
			pq.pos++
		}
	}
}

// remove the max item
// copy the last time to the first item
// push down. You have to sub out the larget of the 2 childs.
func (pq *PQueue) removeMax() int {
	if pq.pos >= 1 {
		maxI := pq.queue[1]
		pq.queue[1] = pq.queue[pq.pos]
		pq.queue[pq.pos] = -1
		pq.pos--

		for i := 1 ; i*2 < pq.pos ; i*=2 {
			if pq.queue[2*i] > pq.queue[2*i+1] {
				if pq.queue[i] < pq.queue[2*i] {
					pq.queue[i], pq.queue[2*i] = pq.queue[2*i], pq.queue[i]
				}	
			} else {
				if pq.queue[i] < pq.queue[2*i+1] {
					pq.queue[i], pq.queue[2*i+1] = pq.queue[2*i+1], pq.queue[i]
				}
			}
		}

		return maxI
	} else {
		return -1
	}
}

func main() {
	var pq PQueue
	pq.init(3)

	pq.insert(1)
	pq.insert(2)
	pq.insert(4)
	pq.insert(5)
	pq.insert(20)
	// pq.insert(7)
	// pq.insert(9)
	// pq.insert(20)
	// pq.insert(11)
	// pq.insert(15)
	fmt.Println(pq.queue)

	fmt.Println(pq.removeMax())
	fmt.Println(pq.queue)
	fmt.Println(pq.removeMax())
	fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())
	// fmt.Println(pq.removeMax())

	fmt.Println(pq.queue)
}

/*
Trying to implment this again.
*/