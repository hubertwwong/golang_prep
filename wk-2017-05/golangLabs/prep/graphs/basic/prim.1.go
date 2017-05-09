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

// hacked it to use NodeG
// merge sorting the results based off the weight.
// return the first node that is in unvisited.
func minDistVertex(adj []NodeG, unvisited map[int]bool) int {
  sortedSPT := mergeSort(adj)

  for i:=0 ; i<len(sortedSPT) ; i++ {
    if _, ok := unvisited[sortedSPT[i].id]; ok {
      return sortedSPT[i].id
    }
  }

  return -1
}

/*
  given a list of vertexes.
  find a non duplicate list of adj vertexes.
  pick the min weighted one.

  takes pointers so you can adjust the params...
  probably not the best way to do this.
*/
func allAdjVertex(visted *map[int]bool, unvisited *map[int]bool, visitedAllAdj *map[int]bool) int {
  return -1
}

// This assumes vertexes start at zero.
func (g *Graph) prim(s int) [][]int {
  // list of visted and visited nodes
  visited := make(map[int]bool)
  unvisited := make(map[int]bool)
  visitedAllAdj := make(map[int]bool)
  // this is just an optimization.

  // probably should use this... because you know that length of paths are length of vertex - -1.
  // edges := make([][]int, len(g.vertexes)-1)
  edges := make([][]int, 0)
  
  // load unvisited with all vertexes.
  for i:=0 ; i<len(g.adjVertexes) ; i++ {
    unvisited[i] = true
  }

  // WORKING ON THIS.....
  curVertex := s
  for ; curVertex != -1 ; {
    fmt.Println("> loop start ===============")
    
    // this needs to be get all adj nodes of visited.
    // get adj nodes.
    adjNodes := g.adj(curVertex)
    fmt.Println("> ver", curVertex, "adj", adjNodes)
    
    // figure out min dist.
    nextVertex := minDistVertex(adjNodes, unvisited)
    fmt.Println("> next from minDist", nextVertex)

    // if no next vertex is found,
    // you are done. break out.
    if nextVertex == -1 {
      break
    }

    // add edge to mst.
    edges = append(edges, []int{curVertex, nextVertex})

    // update visited and unvisited nodes.
    visited[curVertex] = true
    delete(unvisited, curVertex)
    delete(unvisited, nextVertex)

    // update current node
    curVertex = nextVertex
  }

  return edges
}

func main() {
  var g Graph

  // redo this so start at vertex 0

  // g.init(4)
  // g.insertEdge(1, 2, 10)
  // g.insertEdge(2, 1, 10)
  // g.insertEdge(1 ,3, 20)
  // g.insertEdge(3 ,1, 20)
  // g.insertEdge(2, 3, 30)
  // g.insertEdge(3, 2, 30)
  // edges := g.prim(1)

  g.init(5)

  g.insertEdge(0, 1, 10)
  g.insertEdge(1, 0, 10)
  
  g.insertEdge(1 ,2, 20)
  g.insertEdge(2 ,1, 20)
  
  g.insertEdge(2, 3, 30)
  g.insertEdge(3, 2, 30)
  
  g.insertEdge(2, 4, 40)
  g.insertEdge(4, 2, 40)
  
  g.insertEdge(3, 4, 50)
  g.insertEdge(4, 3, 50)

  // sanity check
  fmt.Println(g.adj(1))
  
  edges := g.prim(1)

  fmt.Println(edges)

  // g.init(4)
  // g.insertEdge(0, 1, 1)
  // g.insertEdge(0 ,2, 2)
  // g.insertEdge(0, 3, 1000)
	// g.insertEdge(1, 3, 7)
	// g.insertEdge(2, 3, 4)
  // spt, prevVertex := g.dijkstra(0)

  //printDijkstra(spt, prevVertex)
}

/*

Giving prim a shot...
Curious on what happend on a DAG.



*/