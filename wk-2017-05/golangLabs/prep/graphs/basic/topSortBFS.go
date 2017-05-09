package main

import "fmt"

// ========== QUEUE STUFF ==========

type NodeQ struct {
  next *NodeQ
  val int
}

type Queue struct {
  root *NodeQ
  tail *NodeQ
}

func (q *Queue) enqueue(val int) {
  n := &NodeQ{nil, val}
  if q.root == nil {
    q.root = n
    q.tail = n
  } else {
    (*q.tail).next = n
    q.tail = n
  }
}

func (q *Queue) dequeue() int {
  if q.root != nil {
    retNode := q.root
    q.root = (*q.root).next
    (*retNode).next = nil
    return (*retNode).val
  }
  return -1
}

func (q *Queue) isEmpty() bool {
  if q.tail == nil {
    return false
  } else {
    return true
  }
}

// ========== GRAPH STUFF ==========

type NodeG struct {
  next *NodeG
  val int
}

type Graph struct {
  vertex []*NodeG
}

// constrct a node.
// assumes that its linear. in the vertex labeling...
func (g *Graph) init(numVertex int) {
  if g.vertex == nil {
    g.vertex = make([]*NodeG, numVertex)
  }
}

// insert stuff into the Graph
func (g *Graph) insert(adjList [][]int) {
  //fmt.Println(">insert")
  for i:=0 ; i<len(adjList) ; i++ {
    if g.vertex[adjList[i][0]] == nil {
      //fmt.Println(">insert new")
      g.vertex[adjList[i][0]] = &NodeG{nil, adjList[i][1]}
    } else {
      //fmt.Println(">insert old")
      n := g.vertex[adjList[i][0]]
      for ; n.next != nil ; {
        n = n.next
      }
      (*n).next = &NodeG{nil, adjList[i][1]}
    }
  }
}

// check if a node is adjacent.
func (g *Graph) isAdj(orig, adj int) bool {
  if g.vertex[orig] == nil {
    return false
  } else {
    n := g.vertex[orig]
    for ; n.next != nil ; {
      if (*n).val == adj {
        return true
      }
      n = n.next
    }
    return false
  }
}

// helper to see if node has adj vertex.
func (g *Graph) hasAdj(orig int) bool {
  if g.vertex[orig] == nil {
    return false
  } else {
    return true
  }
}

// return all adj vertex as an array of ints if you want to process it.
func (g *Graph) adj(v int) []int {
  if g.vertex[v] == nil {
    return []int{}
  } else {
    adjNodes := make([]int, 0)
    n := g.vertex[v]
    //fmt.Println("adj", n)
    for {
      adjNodes = append(adjNodes, (*n).val)
      n = n.next
      if n == nil {
        break
      } else if n.next == nil {
        adjNodes = append(adjNodes, (*n).val)
        break
      }
    }
    return adjNodes
  }
}

// insert a single edge into the list.
func (g *Graph) addEdge(s, d int) bool {
  if s >=0 && s< len(g.vertex) && d >= 0 && d < len(g.vertex) {
    if g.vertex[s] == nil {
      g.vertex[s] = &NodeG{nil, d}
    } else {
      for c := g.vertex[s] ; c != nil ; c = (*c).next {
        if (*c).next == nil {
          (*c).next = &NodeG{nil, d}
          break
        }
      }
    }

    return true
  } else {
    return false
  }
}

func (g *Graph) topSortBFS() []int {
  // visited list
  // queue of node...
  return nil
}



func main() {
  //var g Graph
  //g.init(2)
  //g.addEdge(0,1)
  //g.addEdge(1,0)
  //fmt.Println( *g.vertex[1] )

  //var q Queue
  //q.enqueue(4)
  //q.enqueue(2)
  //fmt.Println(q.dequeue())
  //fmt.Println(q.dequeue())
  //fmt.Println(q.dequeue())
}