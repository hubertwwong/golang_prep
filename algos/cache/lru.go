package main

import "fmt"

// === DOUBLE LINKED LIST.
type Node struct {
  key int
  val int
  next *Node
  prev *Node
}

type LinkedList struct {
  head *Node
  tail *Node
}

func (this *LinkedList) delete(cur *Node) (int, int, bool) {
  // guard
  fmt.Println("> LL > delete start", cur)
  if cur == nil {
    return -1, -1, false
  }

  n := (*cur).next
  p := (*cur).prev

  if n == nil && p == nil {
    fmt.Println("> ll > delete > 1")
    // this is the only node.
    this.head = nil
    this.tail = nil
    (*cur).next = nil
    (*cur).prev = nil
  } else if n == nil {
    fmt.Println("> ll > delete > 2")
    // this is the tail node
    this.tail = (*cur).prev
    (*this.tail).next = nil
    (*cur).prev = nil
  } else if p == nil {
    fmt.Println("> ll > delete > 3")
    // this is the head node.
    this.head = (*cur).next
    (*this.head).prev = nil
    (*cur).next = nil
  } else {
    fmt.Println("> ll > delete > 4")
    // you have a previous and next node.
    (*p).next = n
    (*n).prev = p
    (*cur).next = nil
    (*cur).prev = nil
  }

  return (*cur).key, (*cur).val, true
}

// nice shortcut to delete oldest item from the lru cache.
func (this *LinkedList) deleteHead() (int, int, bool) {
  if this.head != nil {
    oldHead := this.head
    this.head = (*this.head).next
    if this.head != nil {
      (*this.head).prev = nil
    }
    (*oldHead).next = nil
    return (*oldHead).key, (*oldHead).val, true
  } else {
    return -1, -1, false
  }
}

func (this *LinkedList) insertTail(key, val int) *Node {
  n := Node{key, val, nil, nil}
  if this.head == nil {
    this.head = &n
    this.tail = &n
  } else {
    (*this.tail).next = &n
    n.prev = this.tail
    this.tail = (*this.tail).next
  }

  return this.tail
} 

func (this *LinkedList) print() {
  fmt.Print("> LL > ")
  for cur := this.head ; cur != nil ; cur = (*cur).next {
    fmt.Print((*cur).key, (*cur).val, " | ")
  }
  fmt.Println("")
}

// === LRU CACHE


type LRUCache struct {  
  h map[int]*Node
  ll LinkedList
  cap int
}

func Constructor(capacity int) LRUCache {
  var lru LRUCache
  lru.h = make(map[int]*Node)
  lru.cap = capacity
  return lru
}

func (this *LRUCache) Get(key int) int {
  fmt.Println("\n> get >", key)
  this.ll.print()  
  
  if v, ok := this.h[key]; ok {
    // stash the value
    curVal := (*v).val
    
    // delete the ll item.
    this.ll.delete(v)

    // insert at the end.
    this.ll.insertTail(key, curVal)

    // redo the key in the hash map.
    this.h[key] = this.ll.tail
    this.ll.print()  
    return curVal
  } else {
    this.ll.print()  
    return -1
  }
}

func (this *LRUCache) Put(key int, value int)  {
  fmt.Println("\n> put >", key, value)
  this.ll.print()  
  //fmt.Println("> put > 1 > key", key, "value", value, "len", len(this.h))
  if v, ok := this.h[key]; ok {
    // replace a key
    // delete ll
    delK, _, _ := this.ll.delete(v)
    // delete the old key
    delete(this.h, delK)
    fmt.Println("> put > updated key", delK)
  } else if len(this.h) == this.cap {
    // eviction check
    // delete the old head node
    delK, _, _ := this.ll.deleteHead()
    // delete the old key
    delete(this.h, delK)
    fmt.Println("> put > evict > ", delK)
    //this.ll.print()
    //fmt.Println(">", this.h)
  }

  // insert new key/value
  ptr := this.ll.insertTail(key, value)
  this.h[key] = ptr
  
  //fmt.Println("> put > 2 > key", key, "value", value, "len", len(this.h))
}




func main() {
  // var ll LinkedList
  // ll.insertTail(3,1)
  // ll.insertTail(4,1)
  // ll.insertTail(5,1)
  // ll.insertTail(6,1)

  // ptr := ll.head
  // ptr = (*ptr).next
  // ptr = (*ptr).next
  // ptr = (*ptr).next
  // ll.delete(ptr)
  
  // ll.print()

//  cache := Constructor(2)

  // cache.Put(1, 1)
  // cache.Put(2, 2)
  // fmt.Println("> m >", cache.Get(1))       // returns 1
  // cache.Put(3, 3)    // evicts key 2
  // fmt.Println("> m >", cache.Get(2))       // returns -1 (not found)
  // cache.Put(4, 4)    // evicts key 1
  // fmt.Println("> m >", cache.Get(1))      // returns -1 (not found)
  // fmt.Println("> m >", cache.Get(3))      // returns 3
  // fmt.Println("> m >", cache.Get(4))      // returns 4

  // fmt.Println("> m >", cache.Get(2))
  // cache.Put(2, 6)
  // fmt.Println("> m >", cache.Get(1))
  // cache.Put(1, 5)
  // cache.Put(1, 2)
  // fmt.Println("> m >", cache.Get(1))
  // fmt.Println("> m >", cache.Get(2))

  // cache := Constructor(3)

  // cache.Put(1, 1)
  // cache.Put(2, 2)
  // cache.Put(3, 3)
  // cache.Put(4, 4)
  // fmt.Println("> m >", cache.Get(4))
  // fmt.Println("> m >", cache.Get(3))
  // fmt.Println("> m >", cache.Get(2))
  // fmt.Println("> m >", cache.Get(1))
  // cache.Put(5, 5)
  // fmt.Println("> m >", cache.Get(1))
  // fmt.Println("> m >", cache.Get(2))
  // fmt.Println("> m >", cache.Get(3))
  // fmt.Println("> m >", cache.Get(4))
  // fmt.Println("> m >", cache.Get(5))

  cache := Constructor(1)
  cache.Put(2, 1)
  fmt.Println("> m >", cache.Get(2))
  cache.Put(3, 2)
  fmt.Println("> m >", cache.Get(2))
  fmt.Println("> m >", cache.Get(3))

  // fmt.Println(cache)

  //fmt.Println(ll.head)
}


/*
["LRUCache","put","get","put","get","get"]
[[1],[2,1],[2],[3,2],[2],[3]]

["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
[[3],[1,1],[2,2],[3,3],[4,4],[4],[3],[2],[1],[5,5],[1],[2],[3],[4],[5]]

[null,null,null,null,null,4,3,2,-1,null,-1,2,-1,4,5]
[null,null,null,null,null,4,3,2,-1,null,-1,2,3,-1,5]

["LRUCache","get","put","get","put","put","get","get"]
[[2],[2],[2,6],[1],[1,5],[1,2],[1],[2]]

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