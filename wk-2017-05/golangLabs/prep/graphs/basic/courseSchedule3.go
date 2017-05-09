package main

import "fmt"

type Node struct {
  val int
  next *Node
}

type Graph struct {
  vertexes []*Node
}

// ==== HELPER ====

func sliceIntInit(num, val int) []int {
  vals := make([]int, num)
  for i := 0 ; i < num ; i++ {
    vals[i] = val
  }
  return vals
}

// using this as a cheat for the unvisited vector.
func sliceIntInitInc(num int) []int {
  vals := make([]int, num)
  for i := 0 ; i < num ; i++ {
    vals[i] = i
  }
  return vals 
}

// inefficinent method.
func sliceRemoveIndex(vals []int, pos int) []int {
  res := make([]int, 0)
  if len(vals) > 0 && pos >= 0 && pos < len(vals) {
    res = append(res, vals[:pos]...)
    res = append(res, vals[pos+1:]...)
  }
  return res
}

func sliceReverse(vals []int) []int {
  for i :=0 ; i < len(vals)/2 ; i++ {
    vals[i], vals[(len(vals)-1)-i] = vals[(len(vals)-1)-i], vals[i]
  }
  return vals
}

func intInSlice(needle int, haystack []int) bool {
  for _, v := range haystack {
    if v == needle {
      return true
    }
  }
  return false
}

// ==== GRAPH ====

func (g *Graph) init(numVertexes int) {
  g.vertexes = make([]*Node, numVertexes)
}

func (g *Graph) addEdge(src, dest int) {
  curNodeP := g.vertexes[src]
  newNode := Node{dest, nil}
  if curNodeP == nil {
    g.vertexes[src] = &newNode
  } else {
    for ; (*curNodeP).next != nil ; curNodeP = (*curNodeP).next {}
    (*curNodeP).next = &newNode
  }
}

func (g *Graph) adj(src int) []int {
  adjEdges := make([]int, 0)

  for curNodeP := g.vertexes[src] ; curNodeP != nil ; curNodeP = (*curNodeP).next {
    adjEdges = append(adjEdges, (*curNodeP).val)
  }

  return adjEdges
}

// add multiple edges. reverse the ordering.
// remember that....
func (g *Graph) addPrereqs(classes [][]int) {
  for i := 0 ; i < len(classes) ; i ++ {
    g.addEdge(classes[i][1], classes[i][0])
  }
}

func (g *Graph) calcAllEdges() []int {
  numEdges := sliceIntInit(len(g.vertexes), -1)
  
  for i := 0 ; i < len(g.vertexes) ; i ++ {
    adjEdges := g.adj(i)
    numEdges[i] = len(adjEdges)
  }

  return numEdges
}

func (g *Graph) topSort() []int {
  unvisited := sliceIntInitInc(len(g.vertexes))
  visited := make([]int, 0)
  edgesRemaining := g.calcAllEdges()
  
  //removeEdge := false
  for ; len(unvisited) > 0 ; {
    //fmt.Println("\n> topSort > unvisited", unvisited, "edgesRemaining", edgesRemaining, "visited", visited)
    // 1. pick out a vertex that has 0 remaing neighbors. This is wrong.
    ui := 0
    for ; ui < len(unvisited) && edgesRemaining[unvisited[ui]] != 0 ; ui++ {}

    // mark if you are going to remove the edge.
    // this is how i'm checking for cycles...
    if ui < len(unvisited) && edgesRemaining[unvisited[ui]] == 0 {
      //removeEdge = true
    } else {
      break
    }

    // 2. stash the vertex
    curV := unvisited[ui]
    //fmt.Println("> unvisited index >", ui, "> curV - unvisited[ui] >", curV)

    // 3. remove zero vertex from unvisited.
    unvisited = sliceRemoveIndex(unvisited, ui)

    // 4. add zero vertex to visited
    visited = append(visited, curV)

    // 5. update the edges.
    for i := 0 ; i < len(edgesRemaining) ; i++ {
      curAdj := g.adj(i)
      if intInSlice(curV, curAdj) {
        //fmt.Println("> updating edge > ", i)
        edgesRemaining[i]--
      }
    }

    // check if you remove an edge.
    // if removeEdge {
    //   removeEdge = false
    // } else {
    //   break
    // }
  }
  
  //fmt.Println(">", len(visited), ">", len(g.vertexes))
  if len(visited) == len(g.vertexes) {
    return sliceReverse(visited)
  } else {
    return []int{}
  }
}

func findOrder(numCourses int, prerequisites [][]int) []int {
  var g Graph
  g.init(numCourses)
  g.addPrereqs(prerequisites)
  return g.topSort()
}

func main() {
  var numCourses int
  prerequisites := make([][]int, 0)

  // [1,0]
  // [0,1]
  // numCourses = 3
  // prerequisites = append(prerequisites, []int{1,0})
  // prerequisites = append(prerequisites, []int{2,1})
  
  // [[1,0],[2,0],[3,1],[3,2]]
  // [0,2,1,3]
  // numCourses = 4
  // prerequisites = append(prerequisites, []int{1,0})
  // prerequisites = append(prerequisites, []int{2,0})
  // prerequisites = append(prerequisites, []int{3,1})
  // prerequisites = append(prerequisites, []int{3,2})
  
  // numCourses = 2
  // prerequisites = append(prerequisites, []int{1,0})
  // prerequisites = append(prerequisites, []int{0,1})
  
  // 8
  //[[1,0],[2,6],[1,7],[6,4],[7,0],[0,5]]
  //[5,4,6,3,2,0,7,1]
  numCourses = 8
  prerequisites = append(prerequisites, []int{1,0})
  prerequisites = append(prerequisites, []int{2,6})
  prerequisites = append(prerequisites, []int{1,7})
  prerequisites = append(prerequisites, []int{6,4})
  prerequisites = append(prerequisites, []int{7,0})
  prerequisites = append(prerequisites, []int{0,5})
  
  // [[1,0],[2,0],[3,1],[3,2]]
  // [0,2,1,3]
  // numCourses = 5
  // prerequisites = append(prerequisites, []int{1,0})
  // prerequisites = append(prerequisites, []int{2,1})
  // prerequisites = append(prerequisites, []int{3,2})
  // prerequisites = append(prerequisites, []int{4,2})
  // prerequisites = append(prerequisites, []int{3,4})
  
  // numCourses = 2
  // blank test case.

  fmt.Println("> num courses", numCourses) 
  fmt.Println("> prereq", prerequisites, "\n")
  fmt.Println("\n> result", findOrder(numCourses, prerequisites))

  // var g Graph
  // g.init(2)
  // g.addEdge(1, 0)
  // g.addEdge(1, 3)
  // g.addEdge(0, 1)
  // g.addEdge(1, 20)
  // g.addEdge(1, 4)
  // fmt.Println(g.adj(1))

  // 2, [[1,0]]
  // [0,1]
  // 4, [[1,0],[2,0],[3,1],[3,2]]
  // [0,1,2,3]
  // [0,2,1,3]
}

/*
  10:45am

  You probably need.
  - Adj graph.
  - A way to reverse the pairs on insert. Some helper func.... You shouldn't touch the graph algo.
  - 2 passes. 1 to compute the adj edges. 2nd pass to compute the sort....

  Top sort:
  1. You can go backwards. or forwards.
  2. Remove the inbound edge that cost zero. store this in a stack.
  3. Recompute the other inbound edge. (basically go to that vertex. get the adj. -1 all adj.)
  4. Repeat until you have no more edges.

  10:52pm
  start on the graph part.

  11:08pm staring at uber eats....

  02:54 didnt' read about the cycle..

  04:02 edge case about not having a prerequisite.

  05:13 pm... this passed.
  but the algo is 1.5 seconds vs. 100ms...
  i'm guessing its the cycle detection algoright...

  removing the dfs cycle detection drop it to 600ms
  so thats a bit better...
  not sure whats left...

  maybe going forward vs. backward in the toplogical sort...


  === Course schedule II ====
  Basically [1,0] means that class 0 is a requirement for class 1.
  Can you construct an list of course where the person can graduate.
  This is basically topological sort.
*/