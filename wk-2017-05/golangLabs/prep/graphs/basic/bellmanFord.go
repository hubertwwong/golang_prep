package main

import "fmt"

func printArrayI(vals [][]int) {
  for i := 0 ; i < len(vals) ; i++ {
    fmt.Printf("%3d", i)
    for j := 0 ; j < len(vals[0]) ; j++ {
      fmt.Printf("%8d", vals[i][j])
    }
    fmt.Println("")
  }
}

type NodeG struct {
  next *NodeG
  val int
  weight int
}

type Graph struct {
  vertexes []*NodeG
}

func (g *Graph) init(numVertex int) {
  if g.vertexes == nil {
    g.vertexes = make([]*NodeG, numVertex)
  }
}

func (g *Graph) insertEdge(src, dest, weight int) {
  newNode := NodeG{nil, dest, weight}
  curVertex := g.vertexes[src]
  
  if curVertex == nil {
    g.vertexes[src] = &newNode
  } else {
    // iterate thru the list
    for ; (*curVertex).next != nil ; curVertex = (*curVertex).next {}

    (*curVertex).next = &newNode
  }
}

func (g *Graph) adj(v int) []*NodeG {
  if g.vertexes[v] == nil {
    return nil
  } else {
    adjNodes := make([]*NodeG, 0)
    n := g.vertexes[v]
    for {
      adjNodes = append(adjNodes, n)
      n = n.next
      if n == nil {
        break
      }
      if n.next == nil {
        adjNodes = append(adjNodes, n)
        break
      }
    }
    return adjNodes
  }
}

func (g *Graph) bellmanFord(srcV int) [][]int {
  // construct the adj edges
  edges := make([][]int, 0)
  // iterate thru each vertex in the adj list.
  for i := 0 ; i < len(g.vertexes) ; i++ {
    adjV := g.adj(i)
    // for each neighbor vertex, add it to the edges list.
    for j := 0 ; j < len(adjV) ; j++ {
      edges = append(edges, []int{i, (*adjV[j]).val, (*adjV[j]).weight})
    }
  }

  // construct the memo table
  // index is the vertex.
  // col 1 is the best cost.
  // col 2 is the prev vertex.
  // 999999 is infinity.
  memo := make([][]int, 0)
  for i := 0 ; i < len(g.vertexes) ; i++ {
    if i == srcV {
      memo = append(memo, []int{0, -1})
    } else {
      memo = append(memo, []int{999999, -1})
    }
  }

  // bellman ford
  for i := 0 ; i + 1 < len(g.vertexes) ; i++ {
    for j := 0 ; j < len(edges) ; j++ {
      curEdgeCost := edges[j][2]
      curEdgeSrcV := edges[j][0]
      curEdgeDestV := edges[j][1]
      prevEdgeCostFromMemo := memo[curEdgeSrcV][0]
      curMemoCost := memo[curEdgeDestV][0]

      if curMemoCost > curEdgeCost + prevEdgeCostFromMemo {
        memo[curEdgeDestV][0] = curEdgeCost + prevEdgeCostFromMemo
        memo[curEdgeDestV][1] = curEdgeSrcV
      }
    }
  }

  //printArrayI(memo)

  return memo
}

func g1() {
  var g Graph
  g.init(2)
  g.insertEdge(0,1,20)
  g.insertEdge(1,0,15)
  result := g.bellmanFord(0)
  printArrayI(result)
}

func g2() {
  var g Graph
  g.init(3)
  g.insertEdge(0,1,20)
  g.insertEdge(1,2,-20)
  g.insertEdge(1,2,10)
  result := g.bellmanFord(1)
  printArrayI(result)
}

// 0 is the start
// https://www.youtube.com/watch?v=obWXjtg0L64
func g3() {
  var g Graph
  g.init(6)
  g.insertEdge(0,2,2)
  g.insertEdge(1,0,1)
  g.insertEdge(2,1,-2)
  g.insertEdge(3,2,-1)
  g.insertEdge(3,0,-4)
  g.insertEdge(4,3,1)
  g.insertEdge(5,4,8)
  g.insertEdge(5,0,10)
  result := g.bellmanFord(5)
  printArrayI(result)
}



func main() {
  g3()
}

/*
bellman ford.

04/27/2017

09:55
working on the actual algo..

10:03p
need to redo adj() to return the node g objects...
You have a weight now. you can't just return the vertex.

10:19p
on the actual algo.
taking a break to help out Allen.

10:26p
back to it.

02:02
after lunch and a bit of goofing off..

*/