package main

import "fmt"

// === STACK stuff

type Node struct {
  Next *Node
  Val int
  MinV int
}

type Stack struct {
  Root *Node
}

func (s *Stack) Push(val int) {
  //fmt.Println(s.Root)
  if s.Root == nil {
    n := Node{nil, val, val}
    s.Root = &n
    //fmt.Println(">>>",s.Root)
  }
  newMin := val
  if val > (*(s.Root)).MinV {
    v := s.Root
    newMin = (*v).MinV
  }
  newRoot := Node{s.Root, val, newMin}
  s.Root = &newRoot
}

func (s *Stack) Pop() (*Node) {
  if s.Root == nil {
    return nil
  }

  retNode := s.Root
  s.Root = (*(s.Root)).Next
  (*retNode).Next = nil
  return retNode
}

func (s *Stack) Min() int {
  if s.Root == nil {
    return -1
  }
  return (*(s.Root)).MinV
}


// === GRAPH STUFF

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

func (g *Graph) topSort() Stack {
  visited := make([]bool, len(g.vertex))
  var stk Stack
  for i:=0 ; i<len(g.vertex) ; i++ {
    stk = g.topSortPath(i, &visited, stk)
  }
  return stk
}

// Traverse the path and return a list of nodes.
// Use s for memo.
// you should be using a stack.
func (g *Graph) topSortPath(v int, visited *[]bool, stk Stack) Stack {
  fmt.Println("> tsp", v)

  // mark item as visited.
  if (*visited)[v] == false { 
    (*visited)[v] = true

    // visit non visited edge.
    if g.hasAdj(v) == true {
      //fmt.Println("has adj")
      // get a list of adjacent edges
      adjEdges := g.adj(v)
      for i:=0 ; i<len(adjEdges) ; i++ {
        curAdjVer := adjEdges[i]
        if (*visited)[curAdjVer] == false {
          fmt.Println("> tsp > visiting")
          stk = g.topSortPath(curAdjVer, visited, stk)
        }
      }
    }

    // push the current item to the stack
    fmt.Println("> pushing", v)
    stk.Push(v)
  }

  return stk
}

func main() {
  var g Graph
  g.init(5)
  adjList := [][]int{[]int{1,2}, []int{2,3}, []int{2,4}}
  g.insert(adjList)
  //fmt.Println((*g.vertex[2]).next)
  //fmt.Println(g.isAdj(2,1))
  //fmt.Println(g.adj(2))

  fmt.Println("> Topological sort")
  topStk := g.topSort()
  fmt.Println("> Topological sort output")
  for i:=topStk.Root ; topStk.Root != nil ; i=topStk.Pop() { 
    fmt.Println("> ", i)
  }
}

/*
  basically 2 problems
  
  1. if you init a list of verticies. how do you tell if that edge is being used?
  2. on the recursion calls going back up.... how do you not push on an item multiple times....
*/