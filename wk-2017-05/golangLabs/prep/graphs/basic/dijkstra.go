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
		fmt.Println("> dj > cv", curVertex)
	  // grab adj edges
    adjNodes := g.adj(curVertex)
    fmt.Println("> dj > adj", adjNodes)

    // compute new distances of adj vertexes.
    // update as necessary.
		for i:=0 ; i<len(adjNodes) ; i++ {
      if adjNodes[i].weight + spt[curVertex] <= spt[adjNodes[i].id] {
        // update new weight for the adj vertex.
        spt[adjNodes[i].id] = adjNodes[i].weight
        // update previous node.
        prevVertex[adjNodes[i].id] = curVertex
        fmt.Println("> dj > cv > spt", adjNodes[i].id, spt[adjNodes[i].id])
        fmt.Println("> dj > cv > prevVertex", adjNodes[i].id, prevVertex[adjNodes[i].id])
      }
    }

    // update visited.
    visited[curVertex] = true
    // remove cur vertex id from unvisited
    delete(unvisited, curVertex)

    // find next vertext to look at.
    curVertex = minDistVertex(spt, unvisited)
  }

  fmt.Println("> dj > spt", spt)
	fmt.Println("> dj > prevVertex", prevVertex)

  return spt, prevVertex
}

func printDijkstra(spt, prevVertex []int) {
  fmt.Println("index  | cost   | prev")
  fmt.Println("--------------------------")
  for i:=0 ; i<len(spt) ; i++ {
    fmt.Printf("%6d | %6d | %6d\n", i, spt[i], prevVertex[i])
  }
}

func main() {
  var g Graph

  // g.init(4)
  // g.insertEdge(1, 2, 20)
  // g.insertEdge(2 ,3, 20)
  // g.insertEdge(1, 3, 3)
  // spt, prevVertex := g.dijkstra(1)

  g.init(4)
  g.insertEdge(0, 1, 1)
  g.insertEdge(0 ,2, 2)
  g.insertEdge(0, 3, 1000)
	g.insertEdge(1, 3, 7)
	g.insertEdge(2, 3, 4)
  spt, prevVertex := g.dijkstra(0)

  printDijkstra(spt, prevVertex)
}

/*

Kinda close...
missing 2 things..
1. the first node.
2. the biggie. other nodes not working. Was not adding the previous best path...

Adjency list graph.
Array of distance and verticies.
Visited hash map..

https://www.youtube.com/watch?v=pVfj6mxhdMw
This is the guide that i'm trying to follow

*/