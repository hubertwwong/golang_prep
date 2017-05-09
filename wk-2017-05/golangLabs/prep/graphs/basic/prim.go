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

func (g *Graph) minDistVertex2(visited *map[int]bool, unvisited *map[int]bool, visitedAllAdj *map[int]bool) (int, int) {
  // need to store a list of adj that you want to compare against.
  validAdjHash := make(map[int]NodeG)
  // store edges of NodeG
  // You need this to figure out the source vertex after you pick the min edge.
  // dest v, src v is the ordering...
  edgePairs := make([][]NodeG, 0)

  // construct a set of adj edges using the visited vertex as start points.
  for k, _ := range (*visited) {
    fmt.Println("> mdv2 > visited v >", k)
    // checks to see if start vertex is not in visitedAllAdj.
    if _, ok := (*visitedAllAdj)[k] ; !ok {
      possibleAdj := g.adj(k)
      // add to the validAdj if not in visited.
      // you are using the hash map to take of the duplicate.
      for _, v2 := range possibleAdj {
        // check to see a possible dest vertex is not in the visted list.
        if _, ok := (*visited)[v2.id] ; !ok {
          validAdjHash[v2.id]=v2
          destNode := v2
          srcNode := NodeG{k, -1, nil}
          // you
          edgePairs = append(edgePairs, []NodeG{srcNode, destNode})
          fmt.Println("> mdv2 > visted adj >", v2)
        }
      }
    }
  }

  // find min cost edge using merge sort.
  // convert a hash to an array.
  validAdj := make([]NodeG, 0)
  for _, v := range validAdjHash {
    validAdj = append(validAdj, v)
  }
  sortedVertex := mergeSort(validAdj)
  fmt.Println("> mdv2 > sorted", sortedVertex)

  // checks to see if you actually found the edge.
  // return -1 if you didn't find anything.
  if len(sortedVertex) > 0 {
    minVertex := sortedVertex[0]

    // update visited, unvisted
    (*visited)[minVertex.id] = true
    delete((*unvisited), minVertex.id)

    // updated visitedAllAdj
    tempAdj := g.adj(minVertex.id)
    allVisited := true
    for _, v := range tempAdj {
      if _, ok := (*visited)[v.id] ; !ok {
        allVisited = false
        break
      }
    }
    if allVisited {
      (*visitedAllAdj)[minVertex.id] = true
    }

    // figure out which src belong to the min dest vertex.
    minSrcVertexId:=-1
    minWeight:=99999
    for i:=0 ; i<len(edgePairs) ; i++ {
      curEdge := edgePairs[i]
      
      // found a possible min id vertex.
      // compute if the weight is min for the min vertex.
      // you can have the same vertex with multiple weights by different edges.
      // we want to pick one with the smallest weight.
      if curEdge[1].id == minVertex.id {
        // if this was the first one.
        if minSrcVertexId == -1 {
          minSrcVertexId = curEdge[0].id
          minWeight = curEdge[0].weight
        } else if minWeight > curEdge[1].weight {
          // found a edge with a smaller weight.
          minSrcVertexId = curEdge[0].id
          minWeight = curEdge[0].weight
        }
      }
    }

    return minSrcVertexId, minVertex.id
  } else {
    return -1, -1
  }
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

  // intial setup
  //curVertex := s
  // basically adding the source vertex to visited and removing it from unvisited to kick things off.
  visited[s] = true
  delete(unvisited, s)

  srcVertex:=0
  destVertex:=0
  for ; destVertex != -1 ; {
    fmt.Println("> loop start ===============")

    srcVertex, destVertex = g.minDistVertex2(&visited, &unvisited, &visitedAllAdj)

    // if no next vertex is found,
    // you are done. break out.
    if destVertex == -1 {
      break
    }

    // add smallest cost edge to mst.
    edges = append(edges, []int{srcVertex, destVertex})
  }

  return edges
}

func main() {
  var g Graph

  // redo this so start at vertex 0

  // = EXAMPLE 1
  
  g.init(3)

  g.insertEdge(0, 1, 10)
  g.insertEdge(1, 0, 10)

  g.insertEdge(1 ,2, 20)
  g.insertEdge(2 ,1, 20)
  
  g.insertEdge(0, 2, 30)
  g.insertEdge(2, 0, 30)

  edges := g.prim(0)

  // = EXAMPLE 2

  // g.init(5)

  // g.insertEdge(0, 1, 10)
  // g.insertEdge(1, 0, 10)
  
  // g.insertEdge(1 ,2, 20)
  // g.insertEdge(2 ,1, 20)
  
  // g.insertEdge(2, 3, 30)
  // g.insertEdge(3, 2, 30)
  
  // g.insertEdge(2, 4, 40)
  // g.insertEdge(4, 2, 40)
  
  // g.insertEdge(3, 4, 50)
  // g.insertEdge(4, 3, 50)

  // // sanity check
  // fmt.Println(g.adj(1))
  
  // edges := g.prim(1)

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