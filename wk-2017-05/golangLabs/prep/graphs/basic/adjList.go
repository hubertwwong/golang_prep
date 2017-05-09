package main

import "fmt"

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
    g.vertex = make([]*NodeG, numVertex+1)
  }
}

// insert stuff into the Graph
func (g *Graph) insert(adjList [][]int) {
  fmt.Println(">insert")
  for i:=0 ; i<len(adjList) ; i++ {
    if g.vertex[adjList[i][0]] == nil {
      fmt.Println(">insert new")
      g.vertex[adjList[i][0]] = &NodeG{nil, adjList[i][1]}
    } else {
      fmt.Println(">insert old")
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

// return all adj vertex as an array of ints if you want to process it.
func (g *Graph) adj(v int) []int {
  if g.vertex[v] == nil {
    return nil
  } else {
    adjNodes := make([]int, 0)
    n := g.vertex[v]
    for {
      adjNodes = append(adjNodes, (*n).val)
      n = n.next
      if n.next == nil {
        adjNodes = append(adjNodes, (*n).val)
        break
      }
    }
    return adjNodes
  }
}

func main() {
  var g Graph
  g.init(4)
  adjList := [][]int{[]int{1,2}, []int{2,3}, []int{2,4}}
  g.insert(adjList)
  fmt.Println((*g.vertex[2]).next)
  fmt.Println(g.isAdj(2,1))
  fmt.Println(g.adj(2))
}