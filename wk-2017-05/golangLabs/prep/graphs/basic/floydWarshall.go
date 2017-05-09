package main

import "fmt"

// HELPER FUNCS
// ===========================================================================

func printArrayI(vals [][]int) {
  for i := 0 ; i < len(vals) ; i++ {
    fmt.Printf("%3d |", i)
    for j := 0 ; j < len(vals[0]) ; j++ {
      fmt.Printf("%8d", vals[i][j])
    }
    fmt.Println("")
  }
}

// return a 2d slice of ints.
func gen2dArray(x, y int) [][]int {
    memo := make([][]int, y)

    for i := 0 ; i < y ; i++ {
      memo[i] = make([]int, x)
    }

    return memo
}

// Graphs
// ===========================================================================

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

// return a array of sourceVertex, destVertex, weightOfEdge
func (g *Graph) edges() [][]int {
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
  return edges
}

func (g *Graph) floydWarshall() [][]int {
  edges := g.edges()

  // generate the initial memo table..
  memo := gen2dArray(len(g.vertexes), len(g.vertexes))
  for i := 0 ; i < len(edges) ; i++ {
    //fmt.Println(edges[i])
    srcV := edges[i][0]
    destV := edges[i][1]
    weight := edges[i][2]
    memo[srcV][destV] = weight
  }

  // print the start of the memo table
  printArrayI(memo)
  fmt.Println("> floyd start")
  
  // run the floyd warshall.
  numVertex := len(g.vertexes)
  for i := 0 ; i < numVertex ; i++ {
    for j := 0 ; j < numVertex ; j++ {
      for k := 0 ; k < numVertex ; k++ {
        if memo[i][j] > memo[i][k] + memo[k][j] {
          memo[i][j] = memo[i][k] + memo[k][j]
        }
      }
    }
  }

  return memo
}

// EXAMPLES..
// ===========================================================================


func g1() {
  var g Graph

  g.init(2)
  g.insertEdge(0,1,20)
  g.insertEdge(1,0,15)
  result := g.floydWarshall()

  printArrayI(result)
}

func g2() {
  var g Graph

  g.init(3)
  g.insertEdge(0,1,20)
  g.insertEdge(1,2,-20)
  g.insertEdge(1,2,10)
  result := g.floydWarshall()

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

  result := g.floydWarshall()
  printArrayI(result)
}

func g4() {
  var g Graph

  g.init(3)
  g.insertEdge(0,1,4)
  g.insertEdge(1,2,10)
  g.insertEdge(0,2,20)
  result := g.floydWarshall()

  printArrayI(result)
}





func main() {
  g4()
}

/*

Floyd warshall..
Shortest path to all edges from all vertex.

The basic algorithm is this.
0. grab a list of edges and weights.
1. init a memo table that is a adjacency matrix.
2. use edge list to populate the memo table.
3. triple loop through the memo table.
4. The basic check is this. Is there a pair of edges who cost is smaller than the single edge. Update the edge cost. Do this to every edge.

*/