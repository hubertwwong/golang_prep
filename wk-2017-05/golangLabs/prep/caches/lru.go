package main

import "fmt"

// === DOUBLE LINKED LIST.
type Node struct {
  val int
  next *Node
  prev *Node
}

type LinkedList struct {
  head *Node
  tail *Node
}

func (this *LinkedList) delete(cur *Node) bool {
  // guard
  if cur == nil {
    return false
  }

  n := (*cur).next
  p := (*cur).prev

  if n == nil && p == nil {
    // this is the only node.
    this.head = nil
    this.tail = nil
    (*cur).next = nil
    (*cur).prev = nil
  } else if n == nil {
    // this is the tail node
    this.tail = (*cur).prev
    (*this.tail).next = nil
    (*cur).prev = nil
  } else if p == nil {
    // this is the head node.
    this.head = (*cur).next
    (*this.head).prev = nil
    (*cur).next = nil
  } else {
    // you have a previous and next node.
    (*p).next = nil
    (*n).prev = nil
    (*cur).next = nil
    (*cur).prev = nil
  }

  return true
}

// nice shortcut to delete oldest item from the lru cache.
func (this *LinkedList) deleteHead() bool {
  oldHead := this.head
  this.head = (*this.head).next

  // 
}

func (this *LinkedList) insertTail(val int) bool {

} 

// === LRU CACHE

// type LRUCache struct {
  
// }


// func Constructor(capacity int) LRUCache {
  
// }


// func (this *LRUCache) Get(key int) int {

// }


// func (this *LRUCache) Put(key int, value int)  {

// }

func main() {

}


/*

SO LRU

My guess:

You get a hash map that points into the double linked list.
hash map contains key.
double link list contains the value.

get just reorders the linked list.
pushes things to the end of the list.

put
can either add to the DLL
or evict somethign off it.
this is basically delete first item
and add the last item.

linked list...
most recent is at the end of the list.


*/