package main

import "fmt"



// GRAPH CONSTRUCTIOn
// ==================

// 2 nodes.
// 2, 4
//
// 2 points to 2 and 4
// 4 points ot 2
func createGraph1() *NodeG {
  n1 := NodeG{2, nil}
  n2 := NodeG{4, nil}
  n1.neighbors = append(n1.neighbors, &n2)
  n1.neighbors = append(n1.neighbors, &n1)
  n2.neighbors = append(n2.neighbors, &n1)

  // fmt.Println(">>>>>>1")
  // fmt.Println(n1)
  // fmt.Println(n2)
  // fmt.Println(">>>>>>2")

  return &n1
}

// 2 nodes
// 2,4
//
// 2 points to 4
// 4 points to 2
func createGraph2() *NodeG {
  n1 := NodeG{2, nil}
  n2 := NodeG{4, nil}
  n1.neighbors = append(n1.neighbors, &n2)
  n2.neighbors = append(n2.neighbors, &n1)

  return &n1
}

// 3 nodes
// 1,2,3
//
// 1 -> 2 -> 3
func createGraph3() *NodeG {
  n1 := NodeG{1, nil}
  n2 := NodeG{2, nil}
  n3 := NodeG{3, nil}
  n1.neighbors = append(n1.neighbors, &n2)
  n2.neighbors = append(n2.neighbors, &n3)

  return &n1
}

// GRAPH ACTUAL
// ============

type NodeG struct {
  val int
  neighbors []*NodeG
}

func cloneGraph(root *NodeG) *NodeG {
  visited := make(map[int]*NodeG)
  return cloneGraphDFS(root, visited)
}

func cloneGraphDFS(root *NodeG, visited map[int]*NodeG) *NodeG {
  // if nil check
  if root == nil {
    fmt.Println("> root is nil")
    return nil
  }

  // if existing in list check.
  rootVal := (*root).val
  if val, ok := visited[rootVal]; ok {
    fmt.Println("> root is visited. returning pointer from hash table that contains the new one.")
    return val
  }

  // clone children current node.
  newRoot := NodeG{rootVal, nil}

  // mark as true when you pass it in in visited.
  visited[rootVal] = &newRoot

  // DFS on neighbors and return attach pointers to neighbors.
  for i := 0 ; i < len((*root).neighbors) ; i++ {
    curNode := (*root).neighbors[i]
    fmt.Println("> dfs on curNode", newRoot, &newRoot)
    retNode := cloneGraphDFS(curNode, visited)
    newRoot.neighbors = append(newRoot.neighbors, retNode)
  }

  fmt.Println("> root returning")
  return &newRoot
}



func main() {
  g1 := createGraph3()
  fmt.Println("g1 neighbors", (*g1).neighbors[0])
  fmt.Println("g1", g1)
  
  fmt.Println(">Cloned")

  g2 := cloneGraph(g1)
  fmt.Println("g2 neighbors", (*g2).neighbors[0])
  fmt.Println("g2", g2)
}

/*

Clone graph.
The basic idea is to mark each visited node in a visited table.
Visited table contain a list of vertex id and the pointers that correspond to the id.
We need the pointers becase we are traversing on the original graph. If we return those nodes, we didn't actually clone anything.
Traverse in a DFS fashion.
If you see the item in visited. return the pointer stored.

*/