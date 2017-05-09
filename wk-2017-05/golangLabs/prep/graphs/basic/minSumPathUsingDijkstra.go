package main

import (
  "errors" 
  "fmt"
)

func mergeSort(vals []NodeG) []NodeG {
  lenVals := len(vals)
  if lenVals <= 1 {
    return vals
  }

  left := mergeSort(vals[:lenVals/2])
  right := mergeSort(vals[lenVals/2:])
  return merge2(left, right)
}

func merge2(left, right []NodeG) []NodeG {
  lenLeft := len(left)
  lenRight := len(right)
  merged := make([]NodeG, lenLeft + lenRight)
  lenMerged := len(merged)

  for i,li,ri:=0,0,0 ; i<lenMerged ; i++ {
    if li == lenLeft {
      merged[i] = right[ri]
      ri++
    } else if ri == lenRight {
      merged[i] = left[li]
      li++
    } else if left[li].weight<right[ri].weight {
      merged[i] = left[li]
      li++
    } else {
      merged[i] = right[ri]
      ri++
    }
  }

  return merged
}

type NodeG struct {
  id int
  weight int
  next *NodeG
}

type Graph struct {
  adjVertexes []*NodeG
}

func (g *Graph) init(numVertex int) {
  if g.adjVertexes == nil {
    g.adjVertexes = make([]*NodeG, numVertex)
  }
}

func (g *Graph) insertEdge(s, d, w int) error {
  if (s >=0 && s < len(g.adjVertexes)) && (d >= 0 && d < len(g.adjVertexes)) {
    nodeD := &NodeG{d, w, nil}
    nodeCur := g.adjVertexes[s]

    if nodeCur == nil {
      // Note about this. initially this is nil, you can't deference this.
      // should refer to the original array...
      g.adjVertexes[s] = nodeD
    } else {
      for {
        if (*nodeCur).next != nil {
          nodeCur = (*nodeCur).next
        } else {
          (*nodeCur).next = nodeD
          break
        }
      }
    }
    return nil
  } else {
    return errors.New("insertEdge: Out of bounds.")
  }
}

// returns the adj NodesG items.
// usually you want just the ints but we have extra data in the nodeG.
func (g *Graph) adj(v int) []NodeG {
  if v < 0 || v >= len(g.adjVertexes) {
    return nil
  }

  adjVertexes := make([]NodeG, 0)

  for curNode:=g.adjVertexes[v] ; curNode != nil ; curNode = (*curNode).next {
    adjVertexes = append(adjVertexes, *curNode)
  }
  
  return adjVertexes
}

func (g *Graph) print() {
  fmt.Println("> Print Graph")
  for i := 0 ; i<len(g.adjVertexes) ; i++ {
    fmt.Println(i, "",g.adj(i))
  }
}

/*
What you are trying to do here...
Find the min spt that is in the unvisited list
i'm using spt index at the name of the vertex so when I sort it, i lose that id.
*/

// convert a list of ints to a list of nodes.
// assuming the values in the array of ints is
func convertToNodeG(spt []int) []NodeG {
  n := make([]NodeG, len(spt))
  for i:=0 ; i<len(spt) ; i++ {
    n[i] = NodeG{i, spt[i], nil}
  }
  return n
}

// hacked it to use NodeG
// merge sorting the results based off the weight.
func minDistVertex(spt []int, unvisited map[int]bool) int {
  sptNodes := convertToNodeG(spt)
  sortedSPT := mergeSort(sptNodes)

  //fmt.Println("> mdv", sortedSPT)

  // checks that each node
  for i:=0 ; i<len(sortedSPT) ; i++ {
    if _, ok := unvisited[sortedSPT[i].id]; ok {
      //fmt.Println("> mdv", sortedSPT[i].id)
      return sortedSPT[i].id
    }
  }

  return -1
}

func (g *Graph) dijkstra(s int) ([]int, []int) {
  spt := make([]int, len(g.adjVertexes))
  prevVertex := make([]int, len(g.adjVertexes))
  visited := make(map[int]bool)
  unvisited := make(map[int]bool)

  // 0. set distances to 999,999 or infinity.
  for i:=0 ; i<len(spt) ; i++ {
    spt[i] = 999999
  }
  // set source vertex to zero
  spt[s] = 0

  // load unvisited nodes.
  // cheating a little since the nodes are labelled 0-n-1.
  for i:=0 ; i<len(g.adjVertexes) ; i++ {
    unvisited[i] = true
  }

  // kicks off grabs source vertex.
  curVertex := s
  // assign the previous of the start node to itself.
  // not sure if you need to do this.
  prevVertex[curVertex] = curVertex

  // computes shortest path distance using dijkstra.
  for ; curVertex != -1 ; {
    //fmt.Println("> dj > cv", curVertex)
    // grab adj edges
    adjNodes := g.adj(curVertex)
    //fmt.Println("> dj > adj", adjNodes)

    // compute new distances of adj vertexes.
    // update as necessary.
    for i:=0 ; i<len(adjNodes) ; i++ {
      if adjNodes[i].weight + spt[curVertex] <= spt[adjNodes[i].id] {
        // update new weight for the adj vertex.
        spt[adjNodes[i].id] = adjNodes[i].weight
        // update previous node.
        prevVertex[adjNodes[i].id] = curVertex
        //fmt.Println("> dj > cv > spt", adjNodes[i].id, spt[adjNodes[i].id])
        //fmt.Println("> dj > cv > prevVertex", adjNodes[i].id, prevVertex[adjNodes[i].id])
      }
    }

    // update visited.
    visited[curVertex] = true
    // remove cur vertex id from unvisited
    delete(unvisited, curVertex)

    // find next vertext to look at.
    curVertex = minDistVertex(spt, unvisited)
  }

  // fmt.Println("> dj > spt", spt)
  // fmt.Println("> dj > prevVertex", prevVertex)

  return spt, prevVertex
}

func minPathSum(grid [][]int) int {
  if len(grid) == 0 {
    return 0
  } else if len(grid[0]) == 0 {
    return 0
  }

  // shift the weights + 1
  for i := 0 ; i < len(grid) ; i++ {
    for j := 0 ; j < len(grid[0]) ; j++ {
      grid[i][j]++
    }
  }

  lenX := len(grid)
  lenY := len(grid[0])

  // init graph to x+y
  // not putting the top left node in the graph. We add that value at the end of the graph. 
  var g Graph
  g.init(lenX*lenY+1)
  g.insertEdge(0, 1, grid[0][0])

  vCur := 1
  for i := 0 ; i < lenX ; i++ {
    for j := 0 ; j < lenY ; j++ {
      vDown := vCur + lenY
      vAccross := vCur + 1
      //fmt.Println("\n> Insert i", i, "j", j, "vCur [", vCur, "] vDown [", vDown, "] vAccross [", vAccross)

      if i + 1 < lenX {
        //fmt.Println("> down >", grid[i+1][j], "at", vCur, vDown)
        g.insertEdge(vCur, vDown, grid[i+1][j])
      }
      if j + 1 < lenY {
        //fmt.Println("> across >", grid[i][j+1], "at", vCur, vAccross)
        g.insertEdge(vCur, vAccross, grid[i][j+1])
      }

      // increment the vertex.
      vCur++

      //fmt.Println("> i", i, "j" ,j, "> grid[ij]", grid[i][j], "vCur", vCur, "vDown", vDown, "vAccross", vAccross)
    }
  }

  // run dikstra algo.
  spt, prevVertex := g.dijkstra(0)

  // compute the smallest cost edge.
  edges := minPathEdge(1, len(g.adjVertexes)-1, spt, prevVertex)
  minSum := 0
  for i := 0 ; i < len(edges) ; i++ {
    minSum += edges[i] - 1
  }

  //g.print()
  //printDijkstra(spt, prevVertex)

  return minSum
}

// given a source and dest.
// construct a list on cost.
// spt is the edge cost
// prevVertex is the previous vertex.
func minPathEdge(src, dest int, spt, prevVertex []int) []int {
  edgeCost := make([]int, 0)
  curV := dest
  edgeCost = append(edgeCost, spt[curV])
  for curV != src {
    curV = prevVertex[curV]
    edgeCost = append(edgeCost, spt[curV])
  }
  return edgeCost
}

func printDijkstra(spt, prevVertex []int) {
  fmt.Println("index  | cost   | prev")
  fmt.Println("--------------------------")
  for i:=0 ; i<len(spt) ; i++ {
    fmt.Printf("%6d | %6d | %6d\n", i, spt[i], prevVertex[i])
  }
}

func main() {
  // grid := [][]int{
  //   []int{11,22,33},
  //   []int{44,55,66},
  //   []int{77,88,99},
  // }

  // grid := [][]int{
  //   []int{11,3},
  //   []int{20,1},
  // }

  // 6
  // grid := [][]int{
  //   []int{1,2,5},
  //   []int{3,2,1},
  // }

  // grid := [][]int{
  //   []int{1,2,5},
  //   []int{20,2,20},
  //   []int{1,20,3},
  // }

  // grid := [][]int{
  //   []int{1,2,5},
  //   []int{20,2,20},
  //   []int{1,20,3},
  // }

  grid := [][]int{
    []int{1,2,5},
    []int{0,2,1},
    []int{0,0,3},
  }

  //var grid [][]int

  fmt.Println(minPathSum(grid))

  // var g Graph

  // g.init(4)
  // g.insertEdge(1, 2, 20)
  // g.insertEdge(2 ,3, 20)
  // g.insertEdge(1, 3, 3)
  // spt, prevVertex := g.dijkstra(1)

  // g.init(4)
  // g.insertEdge(0, 1, 1)
  // g.insertEdge(0 ,2, 2)
  // g.insertEdge(0, 3, 1000)
  // g.insertEdge(1, 3, 7)
  // g.insertEdge(2, 3, 4)
  // spt, prevVertex := g.dijkstra(0)

  // printDijkstra(spt, prevVertex)
}

/*

a
a

MinSumPath
Using dikstra...
Convert the array to edges.
And just use the answer to find the answer.
0 edge is going to the top left node.

i don't think you can have 0 costed edges..
boo...
+1

Added 1 on the graph
and sub 1 after you got the correct edges.

17/61 passed...
Maybe come back to it later.
04/25/2017



a
a
a
a
a

*/